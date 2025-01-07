FROM golang:1.23 AS builder
COPY . /app

WORKDIR /app

RUN make build

FROM debian:stable-slim
RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

ARG APP_NAME
ENV APP_NAME ${APP_NAME}

WORKDIR /app
COPY --from=builder /app/bin/${APP_NAME} bin/
COPY --from=builder /app/configs configs/

CMD "./bin/${APP_NAME}"
