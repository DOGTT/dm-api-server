package service

import (
	"context"
	"net/http"

	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/data"
	"github.com/gin-gonic/gin"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type Service struct {
	grpc_api.UnimplementedDemoRunnerServiceServer
	conf *conf.ServiceConfig
	data *data.DataEntry
}

func New(conf *conf.ServiceConfig, data *data.DataEntry) *Service {
	return &Service{
		conf: conf,
		data: data,
	}
}

func (s *Service) DemoRunnerServiceTextCompletions(c *gin.Context) {
	// Implement me
	req := &grpc_api.TextCompletionsReq{}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Ctx(c).Debug("get req", zap.Any("req", req))
	res, err := s.TextCompletions(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) TextCompletions(c context.Context, req *grpc_api.TextCompletionsReq) (res *grpc_api.TextCompletionsResp, err error) {
	// Implement me
	log.Ctx(c).Debug("grpc impl get req", zap.Any("req", req))
	return
}
