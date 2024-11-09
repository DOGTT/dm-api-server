package pre

import (
	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/runner/model"
)

const (
	PreProcessKindDefault = "default"
)

func PreProcessDefaultNew(c *conf.PreProcess) (p PreProcess, err error) {
	// todo
	// grpc client
	return &PreProcessDefault{}, nil
}

type PreProcessDefault struct {
}

func (p *PreProcessDefault) Do(in *PreProcessIn) (out *model.ModelInferenceIn, err error) {
	return
}
