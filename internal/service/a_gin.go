package service

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"

	api "github.com/DOGTT/dm-api-server/api/base"
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

func getClaimFromContext(ctx context.Context) *utils.TokenClaims {
	return ctx.Value(TOKEN_CLAIM_KEY).(*utils.TokenClaims)
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

func (s *Service) BaseServiceFastRegisterWeChat(c *gin.Context) {
	reqG := &gin_api.FastRegisterWeChatReq{}
	if err := c.ShouldBind(&reqG); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	req := &api.FastRegisterWeChatReq{
		WxCode: *reqG.WxCode,
	}
	if reqG.RegData != nil {
		req.RegData = &api.FastRegisterData{
			Name: *reqG.RegData.Name,
		}
		avatarData, err := base64.StdEncoding.DecodeString(*reqG.RegData.AvatarData)
		if err != nil {
			s.putGinError(c, EM_CommonFail_BadRequest)
			return
		}
		req.RegData.AvatarData = avatarData
	}
	res, err := s.WeChatRegisterFast(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceLoginWeChat(c *gin.Context) {
	req := &api.LoginWeChatReq{}
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
	req := &api.LocationCommonSearchReq{}
	if params.Input != nil {
		req.Input = *params.Input
	}
	res, err := s.LocationCommonSearch(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceMediaPutPresignURLBatchGet(c *gin.Context, params gin_api.BaseServiceMediaPutPresignURLBatchGetParams) {
	req := &api.MediaPutPresignURLBatchGetReq{
		MediaType: api.MediaType(*params.MediaType),
		Count:     *params.Count,
	}
	res, err := s.MediaPutPresignURLBatchGet(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelTypeList(c *gin.Context) {
	req := &api.ChannelTypeListReq{}
	res, err := s.ChannelTypeList(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelCreate(c *gin.Context) {
	req := &api.ChannelCreateReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.ChannelCreate(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelDelete(c *gin.Context, params gin_api.BaseServiceChannelDeleteParams) {
	req := &api.ChannelDeleteReq{
		ChId: utils.ConvertToUint64(*params.ChId),
	}
	res, err := s.ChannelDelete(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelUpdate(c *gin.Context) {
	req := &api.ChannelUpdateReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.ChannelUpdate(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelBaseQueryByBound(c *gin.Context) {
	req := &api.ChannelBaseQueryByBoundReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.ChannelBaseQueryByBound(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelDetailQueryById(c *gin.Context, params gin_api.BaseServiceChannelDetailQueryByIdParams) {
	req := &api.ChannelDetailQueryByIdReq{
		ChId: utils.ConvertToUint64(*params.ChId),
	}
	res, err := s.ChannelDetailQueryById(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelFullQueryById(c *gin.Context, params gin_api.BaseServiceChannelFullQueryByIdParams) {
	req := &api.ChannelFullQueryByIdReq{
		ChId: utils.ConvertToUint64(*params.ChId),
	}
	res, err := s.ChannelFullQueryById(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelInteraction(c *gin.Context) {
	req := &api.ChannelInteractionReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.ChannelInteraction(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelComment(c *gin.Context) {
	req := &api.ChannelCommentReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest.PutDesc(err.Error()))
		return
	}
	res, err := s.ChannelComment(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
