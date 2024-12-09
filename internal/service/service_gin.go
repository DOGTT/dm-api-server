package service

import (
	"errors"
	"net/http"

	api "github.com/DOGTT/dm-api-server/api/grpc"
	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/gin-gonic/gin"
)

func (s *Service) putGinError(c *gin.Context, err error) {

	var (
		httpCode int
		em       *ErrMsg
		emAPI    *api.ErrorMessage
	)
	if errors.As(err, &em) {
		emAPI = &api.ErrorMessage{
			Code: em.Code,
			Desc: em.Desc,
		}
		httpCode = em.HttpStatus
	} else {
		emAPI = &api.ErrorMessage{
			Code: "CommonFail.Unknown",
			Desc: err.Error(),
		}
		httpCode = http.StatusInternalServerError
	}
	c.JSON(httpCode, emAPI)
}

func (s *Service) BaseServiceWeChatFastRegister(c *gin.Context) {
	req := &grpc_api.WeChatFastRegisterReq{}
	if err := c.Bind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.WeChatFastRegister(c, req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceWeChatLogin(c *gin.Context) {
	req := &grpc_api.WeChatLoginReq{}
	if err := c.Bind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.WeChatLogin(c, req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
