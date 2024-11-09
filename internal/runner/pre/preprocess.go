package pre

import (
	"errors"

	"github.com/dog-g/dog-api-server/internal/conf"
	"github.com/dog-g/dog-api-server/internal/runner/model"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type PreProcess interface {
	Do(in *PreProcessIn) (out *model.ModelInferenceIn, err error)
}

type PreProcessIn struct {
	Prompt string
}

func New(c *conf.PreProcess) (p PreProcess, err error) {
	switch c.Kind {
	case PreProcessKindDefault:
		return PreProcessDefaultNew(c)
	default:
		err = errors.New("invalid kind")
		log.L().Error("new error", zap.Error(err))
		return
	}
}
