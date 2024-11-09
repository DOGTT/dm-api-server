package post

import (
	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/runner/model"
)

const (
	PostProcessKind = "default"
)

func PostProcessDefaultNew(c *conf.PostProcess) (p PostProcess, err error) {
	// todo
	// grpc client
	return &PostProcessDefault{}, nil
}

type PostProcessDefault struct {
}

func (p *PostProcessDefault) Do(in *model.ModelInferenceOut) (out *PostProcessOut, err error) {
	return
}
