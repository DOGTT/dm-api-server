package runner

import (
	"context"
	"errors"

	grpc_api "github.com/dog-g/dog-api-server/api/grpc"
	"github.com/dog-g/dog-api-server/internal/conf"
	"github.com/dog-g/dog-api-server/internal/runner/model"
	"github.com/dog-g/dog-api-server/internal/runner/post"
	"github.com/dog-g/dog-api-server/internal/runner/pre"
)

func New(c *conf.RunnerConfig) (r *Runner, err error) {
	pipeMap := map[string]*Pipe{}
	for name, pipe := range c.ModelPipelines {
		pre, err := pre.New(pipe.PreProcess)
		if err != nil {
			return nil, err
		}
		post, err := post.New(pipe.PostProcess)
		if err != nil {
			return nil, err
		}
		model, err := model.New(pipe.ModelInference)
		if err != nil {
			return nil, err
		}
		pipeMap[string(name)] = &Pipe{
			Pre:   pre,
			Post:  post,
			Model: model,
		}
	}
	return &Runner{
		c: c,
	}, nil
}

type Runner struct {
	c       *conf.RunnerConfig
	pipeMap map[string]*Pipe
}

type Pipe struct {
	Pre   pre.PreProcess
	Model model.ModelInference
	Post  post.PostProcess
}

func (u *Runner) GetPipe(modelName string) (p *Pipe, err error) {
	p, ok := u.pipeMap[modelName]
	if !ok {
		err = errors.New("model pipeline not exist")
	}
	return
}

func (u *Runner) Do(c context.Context, req *grpc_api.TextCompletionsReq) (res *grpc_api.TextCompletionsResp, err error) {
	p, err := u.GetPipe(req.GetModel())
	if err != nil {
		return nil, err
	}
	modelIn, err := p.Pre.Do(&pre.PreProcessIn{
		Prompt: req.GetPrompt(),
	})
	if err != nil {
		return nil, err
	}
	modelOut, err := p.Model.Do(modelIn)
	if err != nil {
		return nil, err
	}

	pipeRes, err := p.Post.Do(modelOut)
	if err != nil {
		return nil, err
	}

	choices := make([]*grpc_api.TextCompletionChoice, len(pipeRes.Choices))
	for i := range pipeRes.Choices {
		choices[i] = &grpc_api.TextCompletionChoice{
			Text: pipeRes.Choices[i].Text,
		}
	}
	return &grpc_api.TextCompletionsResp{
		Id:      "uuid-todo",
		Model:   req.GetModel(),
		Choices: choices,
	}, nil

}
