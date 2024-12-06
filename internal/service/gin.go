package service

import (
	"net/http"

	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/gin-gonic/gin"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func (s *Service) BaseServiceWeChatLogin(c *gin.Context) {
	// Implement me
	req := &grpc_api.WeChatLoginReq{}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Ctx(c).Debug("get req", zap.Any("req", req))
	res, err := s.WeChatLogin(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceWeChatRegisterWithLogin(c *gin.Context) {
	// Implement me
	req := &grpc_api.RegisterWithLoginReq{}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Ctx(c).Debug("get req", zap.Any("req", req))
	res, err := s.WeChatRegisterWithLogin(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
