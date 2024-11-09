package post

import (
	"errors"

	"github.com/dog-g/dog-api-server/internal/conf"
	"github.com/dog-g/dog-api-server/internal/runner/model"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type PostProcess interface {
	Do(in *model.ModelInferenceOut) (out *PostProcessOut, err error)
}

type PostProcessOut struct {
	Choices []*Choice
}

type Choice struct {
	Text string
}

func New(c *conf.PostProcess) (p PostProcess, err error) {
	switch c.Kind {
	case PostProcessKind:
		return PostProcessDefaultNew(c)
	default:
		err = errors.New("invalid kind")
		log.L().Error("new error", zap.Error(err))
		return
	}
}
