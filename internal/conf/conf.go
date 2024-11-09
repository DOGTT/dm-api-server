package conf

import (
	"time"

	"go.uber.org/zap"
)

type Config struct {
	Server *Server
	Log    *zap.Config
	Metric *MetricConfig
	Runner *RunnerConfig
}

type Server struct {
	HTTP HTTPServer
	GRPC GRPCServer
}

type HTTPServer struct {
	Enable       bool          `default:"true"`
	Addr         string        `default:":8080"`
	Timeout      time.Duration `default:"1s"`
	EnableMetric bool          `yaml:"enable_metric"`
	EnableTrace  bool          `yaml:"enable_trace"`
}

type GRPCServer struct {
	Enable       bool          `default:"true"`
	Addr         string        `default:":8081"`
	Timeout      time.Duration `default:"1s"`
	EnableMetric bool          `yaml:"enable_metric"`
	EnableTrace  bool          `yaml:"enable_trace"`
}

type MetricConfig struct {
	Enable bool   `default:"true"`
	Addr   string `default:":8002"`
}

type RunnerConfig struct {
	ModelPipelines map[ModelName]*ModelPipeline `yaml:"model_pipelines"`
}

type ModelName string

type ModelPipeline struct {
	PreProcess     *PreProcess     `yaml:"pre_process"`
	ModelInference *ModelInference `yaml:"model_inference"`
	PostProcess    *PostProcess    `yaml:"post_process"`
}

type ModelInference struct {
	Kind     string
	GRPCSpec *GRPCSpec `yaml:"spec"`
}

type PreProcess struct {
	Kind     string
	GRPCSpec *GRPCSpec `yaml:"spec"`
}

type PostProcess struct {
	Kind     string
	GRPCSpec *GRPCSpec `yaml:"spec"`
}

type GRPCSpec struct {
	GRPCEndpoint string        `yaml:"grpc_endpoint"`
	Timeout      time.Duration `default:"1s"`
}
