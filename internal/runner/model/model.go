package model

import (
	"errors"

	"github.com/dog-g/dog-api-server/internal/conf"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type ModelInference interface {
	Do(in *ModelInferenceIn) (out *ModelInferenceOut, err error)
}

type ModelInferenceIn struct {
	TokenIDs []int
}

type ModelInferenceOut struct {
	TokenIDs []int
}

func New(c *conf.ModelInference) (m ModelInference, err error) {
	switch c.Kind {
	case ModelInferenceKindTriton:
		return ModelInferenceTritonNew(c)
	default:
		err = errors.New("invalid kind")
		log.L().Error("new error", zap.Error(err))
		return
	}
}
