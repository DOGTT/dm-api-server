package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	api "github.com/DOGTT/dm-api-server/api/base"
	gin_api "github.com/DOGTT/dm-api-server/api/gin"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/gin-gonic/gin"
)

func withGinContext(c *gin.Context) context.Context {
	ctx := context.WithValue(c, utils.ContextKey("Server-Origin"), "gin-server")
	for key, value := range c.Keys {
		ctx = context.WithValue(ctx, utils.ContextKey(key), value)
	}
	return ctx
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
	req := &api.FastRegisterWeChatReq{}
	if err := c.ShouldBind(req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
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

func (s *Service) BaseServiceUserPeGet(c *gin.Context, params gin_api.BaseServiceUserPeGetParams) {
	req := &api.UserPeGetReq{}
	if params.UserId != nil {
		req.UserId = *params.UserId
	}
	res, err := s.UserPeGet(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceUserPetUpdate(c *gin.Context) {
	req := &api.UserPetUpdateReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.UserPetUpdate(withGinContext(c), req)
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

func (s *Service) BaseServiceMediaPutURLBatchGet(c *gin.Context, params gin_api.BaseServiceMediaPutURLBatchGetParams) {
	req := &api.MediaPutURLBatchGetReq{
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
		fmt.Println(err)
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
		ChanId: *params.ChanId,
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

func (s *Service) BaseServiceChannelFullQueryById(c *gin.Context, params gin_api.BaseServiceChannelFullQueryByIdParams) {
	req := &api.ChannelFullQueryByIdReq{
		ChanId: *params.ChanId,
	}
	res, err := s.ChannelFullQueryById(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelInx(c *gin.Context) {
	req := &api.ChannelInxReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.ChannelInx(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelPostDelete(c *gin.Context, params gin_api.BaseServiceChannelPostDeleteParams) {
	req := &api.ChannelPostDeleteReq{
		ChanId: utils.CopyFromPtr(params.ChanId),
		PostId: utils.CopyFromPtr(params.PostId),
	}
	res, err := s.ChannelPostDelete(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelPostCreate(c *gin.Context) {
	req := &api.ChannelPostCreateReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.ChannelPostCreate(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelPostUpdate(c *gin.Context) {
	req := &api.ChannelPostUpdateReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.ChannelPostUpdate(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelPostInx(c *gin.Context) {
	req := &api.ChannelPostInxReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.ChannelPostInx(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelPostLoad(c *gin.Context) {
	req := &api.ChannelPostLoadReq{}
	if err := c.ShouldBind(&req); err != nil {
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.ChannelPostLoad(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
