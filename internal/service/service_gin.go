package service

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"

	api "github.com/DOGTT/dm-api-server/api/base"
	base_api "github.com/DOGTT/dm-api-server/api/base"
	gin_api "github.com/DOGTT/dm-api-server/api/gin"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/gin-gonic/gin"
)

type ContextKey string

const (
	TOKEN_CLAIM_KEY ContextKey = "Token-Claim"
)

func withGinContext(c *gin.Context) context.Context {
	ctx := context.WithValue(c, ContextKey("Server-Origin"), "gin-server")
	for key, value := range c.Keys {
		ctx = context.WithValue(ctx, ContextKey(key), value)
	}
	return ctx
}

func getClaimFromContext(ctx context.Context) utils.TokenClaims {
	return ctx.Value(TOKEN_CLAIM_KEY).(utils.TokenClaims)
}

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
	reqG := &gin_api.WeChatRegisterFastReq{}
	if err := c.ShouldBind(&reqG); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	req := &api.WeChatRegisterFastReq{
		WxCode: *reqG.WxCode,
	}
	if req.Pet != nil {
		avatarData, err := base64.StdEncoding.DecodeString(*reqG.Pet.AvatarData)
		if err != nil {
			s.putGinError(c, EM_CommonFail_BadRequest)
			return
		}
		req.Pet = &api.PetInfoReg{
			Name:       *reqG.Pet.Name,
			AvatarData: avatarData,
		}
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.WeChatRegisterFast(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceWeChatLogin(c *gin.Context) {
	req := &base_api.WeChatLoginReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.WeChatLogin(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceLocationCommonSearch(c *gin.Context, params gin_api.BaseServiceLocationCommonSearchParams) {
	req := &base_api.LocationCommonSearchReq{
		Input: *params.Input,
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.LocationCommonSearch(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceObjectPutPresignURLBatchGet(c *gin.Context, params gin_api.BaseServiceObjectPutPresignURLBatchGetParams) {
	req := &base_api.ObjectPutPresignURLBatchGetReq{
		ObjectType:  base_api.ObjectType(*params.ObjectType),
		ObjectCount: *params.ObjectCount,
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.ObjectPutPresignURLBatchGet(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpCreate(c *gin.Context) {
	req := &base_api.PofpCreateReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.PofpCreate(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpDelete(c *gin.Context, params gin_api.BaseServicePofpDeleteParams) {

	req := &base_api.PofpDeleteReq{
		Uuid: *params.Uuid,
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.PofpDelete(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpUpdate(c *gin.Context) {
	req := &base_api.PofpUpdateReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.PofpUpdate(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpBaseQueryByBound(c *gin.Context) {
	req := &base_api.PofpBaseQueryByBoundReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.PofpBaseQueryByBound(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpDetailQueryById(c *gin.Context, params gin_api.BaseServicePofpDetailQueryByIdParams) {
	req := &base_api.PofpDetailQueryByIdReq{
		Uuid: *params.Uuid,
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.PofpDetailQueryById(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpFullQueryById(c *gin.Context, params gin_api.BaseServicePofpFullQueryByIdParams) {
	req := &base_api.PofpFullQueryByIdReq{
		Uuid: *params.Uuid,
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.PofpFullQueryById(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpInteraction(c *gin.Context) {
	req := &base_api.PofpInteractionReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.PofpInteraction(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpComment(c *gin.Context) {
	req := &base_api.PofpCommentReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	if err := req.Validate(); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.PofpComment(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
