package model

import "github.com/DOGTT/dm-api-server/internal/conf"

const (
	ModelInferenceKindTriton = "triton"
)

func ModelInferenceTritonNew(c *conf.ModelInference) (m ModelInference, err error) {
	// todo
	// grpc client
	return &ModelInferenceTriton{}, nil
}

type ModelInferenceTriton struct {
}

func (p *ModelInferenceTriton) Do(in *ModelInferenceIn) (out *ModelInferenceOut, err error) {
	return
}
