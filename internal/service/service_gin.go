package service

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"

	gin_api "github.com/DOGTT/dm-api-server/api/gin"
	api "github.com/DOGTT/dm-api-server/api/grpc"
	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
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
	req := &gin_api.WeChatRegisterFastReq{}
	if err := c.ShouldBind(&req); err != nil {
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
	res, err := s.WeChatRegisterFast(withGinContext(c), gReq)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceWeChatLogin(c *gin.Context) {
	req := &grpc_api.WeChatLoginReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
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
	req := &grpc_api.LocationCommonSearchReq{
		Input: *params.Input,
	}
	res, err := s.LocationCommonSearch(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceObjectPutPresignURLBatchGet(c *gin.Context, params gin_api.BaseServiceObjectPutPresignURLBatchGetParams) {
	req := &grpc_api.ObjectPutPresignURLBatchGetReq{
		ObjectType:  grpc_api.ObjectType(*params.ObjectType),
		ObjectCount: *params.ObjectCount,
	}
	res, err := s.ObjectPutPresignURLBatchGet(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpCreate(c *gin.Context) {
	req := &grpc_api.PofpCreateReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.PofpCreate(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpDelete(c *gin.Context) {
	req := &grpc_api.PofpDeleteReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
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
	req := &grpc_api.PofpUpdateReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
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
	req := &grpc_api.PofpBaseQueryByBoundReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
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
	req := &grpc_api.PofpDetailQueryByIdReq{
		Uuid: *params.Uuid,
	}
	res, err := s.PofpDetailQueryById(withGinContext(c), req)
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
	res, err := s.PofpFullQueryById(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePofpInteraction(c *gin.Context) {
	req := &grpc_api.PofpInteractionReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
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
	req := &grpc_api.PofpCommentReq{}
	if err := c.ShouldBind(&req); err != nil {
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
