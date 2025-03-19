package conf

import (
	"strings"
	"time"

	"go.uber.org/zap"
)

type Config struct {
	Server *Server
	Log    *zap.Config
	Metric *MetricConfig

	Service *ServiceConfig
	Data    *DataConfig
}

type Server struct {
	HTTP HTTPServer
	GRPC GRPCServer
}

type HTTPServer struct {
	Enable            bool          `default:"true"`
	Addr              string        `default:":8080"`
	Timeout           time.Duration `default:"1s"`
	AuthWhitePathlist []string      `yaml:"auth_white_pathlist"`
	EnableMetric      bool          `yaml:"enable_metric"`
	EnableTrace       bool          `yaml:"enable_trace"`
	EnableSwagger     bool          `yaml:"enable_swagger"`
	EnableCORS        bool          `yaml:"enable_cors"`
}

func GetAddrSplit(addr string) (ip string, portStr string) {

	s := strings.Split(addr, ":")
	if len(s) < 2 {
		return
	}
	return s[0], s[1]
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

type GRPCSpec struct {
	GRPCEndpoint string        `yaml:"grpc_endpoint"`
	Timeout      time.Duration `default:"1s"`
}
