NAME := $(shell basename ${PWD})
GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
BUILDTS:=$(shell date -u '+%Y-%m-%d %I:%M:%S')
GIT_HASH:=$(shell git rev-parse HEAD)
GIT_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
VERSION:=$(shell git describe --tags --always)
IMAGE_VERSION:=$(GIT_BRANCH)-$(VERSION)

LDFLAGS += -X 'main.Name=$(NAME)'
LDFLAGS += -X 'main.Version=$(VERSION)'
LDFLAGS += -X 'main.BuildTS=$(BUILDTS)'
LDFLAGS += -X 'main.GitHash=$(GIT_HASH)'
LDFLAGS += -X 'main.GitBranch=$(GIT_BRANCH)'

# basic go commands
export GOPROXY=https://goproxy.cn,direct

# proto files
INTERNAL_PROTO_FILES=$(shell find pkg -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

# protoc install
PROTOC_ZIP:=protoc-3.14.0-linux-x86_64.zip
ifeq ($(GOHOSTOS), darwin)
	PROTOC_ZIP=protoc-3.14.0-osx-x86_64.zip
endif

.PHONY: setup
# setup common utils: protoc
setup:
	@echo "install protoc"
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$(PROTOC_ZIP)
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_ZIP)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.62.2

.PHONY: init
# setup go utils
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: internal
# generate internal proto
internal:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	@mkdir -p ./api/base
	@mkdir -p ./api/openapi
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api/base \
		   --go-grpc_out=paths=source_relative:./api/base \
	       --openapi_out=fq_schema_naming=false,naming=proto,default_response=false,paths=source_relative:./api/openapi/ \
	       $(API_PROTO_FILES)
	oapi-codegen -package apigin -generate types,spec,client,gin ./api/openapi/openapi.yaml > ./api/gin/gin.gen.go

.PHONY: generate
# generate
generate:
	go mod tidy
	go generate ./...

.PHONY: vendor
# vendor code
vendor: go.mod go.sum
	go mod download
	go mod vendor
	go mod tidy

.PHONY: lint
# lint code
lint:
	golangci-lint run -v --allow-parallel-runners --fix --timeout  10m

GOSEC_IMAGE=registry.xxxxx.com/gosec:2.13.1
.PHONY: gosec
# go securty check
gosec:
	docker run -it --rm -v ${PWD}:/code -w /code ${GOSEC_IMAGE} -exclude=G101,G104,G105 -fmt=text -exclude-generated -exclude-dir=tools ./...

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "$(LDFLAGS)" -o ./bin/ ./...

UNAME=$(shell uname)
# package docker image
package:
	docker build --build-arg APP_NAME=$(NAME) -f Dockerfile -t ghcr.io/dogtt/$(NAME):$(IMAGE_VERSION) .

.PHONY: docker
UNAME=$(shell uname)
# build docker image
docker:
	docker build --build-arg APP_NAME=$(NAME) -f Dockerfile -t ghcr.io/dogtt/$(NAME):$(IMAGE_VERSION) ./

.PHONY: push-image
# push docker image to repo
push-image: docker
	docker push ghcr.io/dogtt/$(NAME):$(IMAGE_VERSION)

.PHONY: release-chart
# release helm chart
release-chart:
	sed -i "s|version:\ .*|version:\ \"${CHART_VERSION}\"|" ./charts/ams-inference-gateway/Chart.yaml
	sed -i "s|Version:\ [0-9]\\+\.[0-9]\\+\.[0-9]\\+|Version:\ ${CHART_VERSION}|" ./charts/ams-inference-gateway/README.md
	sed -i "s|badge/Version-[0-9]\\+\.[0-9]\\+\.[0-9]\\+|badge/Version-${VERSION}--${GIT_HASH}|" ./charts/ams-inference-gateway/README.md

	sed -i "s|appVersion:\ .*|appVersion:\ \"${APP_VERSION}\"|" ./charts/ams-inference-gateway/Chart.yaml
	sed -i "s|AppVersion:\ [0-9]\\+\.[0-9]\\+\.[0-9]\\+|AppVersion:\ ${APP_VERSION}|" ./charts/ams-inference-gateway/README.md
	sed -i "s|badge/AppVersion-[0-9]\\+\.[0-9]\\+\.[0-9]\\+|badge/AppVersion-${VERSION}--${GIT_HASH}|" ./charts/ams-inference-gateway/README.md

	sed -i "s|tag:\ .*|tag:\ ${IMAGE_VERSION}|" ./charts/ams-inference-gateway/values.yaml

	helm repo update
	helm cm-push ./charts/ams-inference-gateway xxxxx-studio

packages = $(shell go list ./...|grep -v /vendor/)
.PHONY: test
# ut test
test:
	go test -v -coverprofile=coverage.out  ${packages}
	gocov convert coverage.out | gocov-html > coverage.html
	# go tool cover -func=coverage.out
	gocov convert coverage.out | gocov report

.PHONY: test_only_%
# ut test for specific Function
test_only_%:
	go test -v -run=^$*$$ ${packages}


# show help
help:
	@echo ''
	@echo VERSION=$(VERSION)
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
