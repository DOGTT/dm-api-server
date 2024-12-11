package service

import (
	"encoding/base64"
	"errors"
	"net/http"

	gin_api "github.com/DOGTT/dm-api-server/api/gin"
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

func (s *Service) BaseServiceWeChatRegisterFast(c *gin.Context) {
	req := &gin_api.WeChatRegisterFastReq{}
	if err := c.Bind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	gReq := &api.WeChatRegisterFastReq{
		WxCode: *req.WxCode,
	}
	if req.Pet != nil {
		avatarData, err := base64.StdEncoding.DecodeString(*req.Pet.AvatarData)
		if err != nil {
			s.putGinError(c, EM_CommonFail_BadRequest)
			return
		}
		gReq.Pet = &api.PetInfoReg{
			Name:       *req.Pet.Name,
			AvatarData: avatarData,
		}
	}
	res, err := s.WeChatRegisterFast(c, gReq)
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

func (s *Service) BaseServiceLocationCommonSearch(c *gin.Context, params gin_api.BaseServiceLocationCommonSearchParams) {
	req := &grpc_api.LocationCommonSearchReq{
		Input: *params.Input,
	}
	res, err := s.LocationCommonSearch(c, req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpCreate(c *gin.Context) {
	req := &grpc_api.PofpCreateReq{}
	if err := c.Bind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.PofpCreate(c, req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpDelete(c *gin.Context) {
	req := &grpc_api.PofpDeleteReq{}
	if err := c.Bind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.PofpDelete(c, req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpUpdate(c *gin.Context) {
	req := &grpc_api.PofpUpdateReq{}
	if err := c.Bind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.PofpUpdate(c, req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpDetailQueryById(c *gin.Context, params gin_api.BaseServicePofpDetailQueryByIdParams) {
	req := &grpc_api.PofpDetailQueryByIdReq{
		Uuid: *params.Uuid,
	}
	res, err := s.PofpDetailQueryById(c, req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpFullQueryById(c *gin.Context, params gin_api.BaseServicePofpFullQueryByIdParams) {
	req := &grpc_api.PofpFullQueryByIdReq{
		Uuid: *params.Uuid,
	}
	res, err := s.PofpFullQueryById(c, req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpInteraction(c *gin.Context) {
	req := &grpc_api.PofpInteractionReq{}
	if err := c.Bind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.PofpInteraction(c, req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpComment(c *gin.Context) {
	req := &grpc_api.PofpCommentReq{}
	if err := c.Bind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.PofpComment(c, req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
