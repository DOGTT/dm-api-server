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
	req := &api.MediaPutURLBatchGetReq{}
	if params.Bucket != nil {
		req.Bucket = api.MediaBucket(*params.Bucket)
	}
	if params.Count != nil {
		req.Count = *params.Count
	}
	res, err := s.MediaPutURLBatchGet(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceSystemNotifyGet(c *gin.Context, params gin_api.BaseServiceSystemNotifyGetParams) {
	req := &api.SystemNotifyGetReq{}
	if params.LastNotifyId != nil {
		req.LastNotifyId = *params.LastNotifyId
	}
	res, err := s.SystemNotifyGet(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceLoginWeChat(c *gin.Context) {
	req := &api.LoginWeChatReq{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.LoginWeChat(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceFastRegisterWeChat(c *gin.Context) {
	req := &api.FastRegisterWeChatReq{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.FastRegisterWeChat(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceUserPetUpdate(c *gin.Context) {
	req := &api.UserPetUpdateReq{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
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

func (s *Service) BaseServiceUserInx(c *gin.Context) {
	req := &api.UserInxReq{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.UserInx(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePetMarkerInfoList(c *gin.Context) {
	req := &api.PetMarkerInfoReq{}
	res, err := s.PetMarkerInfoList(withGinContext(c), req)
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

func (s *Service) BaseServiceChannelUpdate(c *gin.Context) {
	req := &api.ChannelUpdateReq{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
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

func (s *Service) BaseServiceChannelDelete(c *gin.Context, params gin_api.BaseServiceChannelDeleteParams) {
	req := &api.ChannelDeleteReq{}
	if params.ChannelId != nil {
		req.ChannelId = *params.ChannelId
	}
	res, err := s.ChannelDelete(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelGet(c *gin.Context, params gin_api.BaseServiceChannelGetParams) {
	req := &api.ChannelGetReq{}
	if params.ChannelId != nil {
		req.ChannelId = *params.ChannelId
	}
	res, err := s.ChannelGet(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelQueryByUser(c *gin.Context, params gin_api.BaseServiceChannelQueryByUserParams) {
	req := &api.ChannelQueryByUserReq{}
	if params.UserId != nil {
		req.UserId = *params.UserId
	}
	if params.IxnState != nil {
		req.IxnState = api.UserIxnState(*params.IxnState)
	}
	if params.IxnEvent != nil {
		req.IxnEvent = api.UserIxnEvent(*params.IxnEvent)
	}
	if params.ExtTypes != nil {
		req.ExtTypes = make([]api.ChannelExtType, 0)
		for _, t := range *params.ExtTypes {
			req.ExtTypes = append(req.ExtTypes, api.ChannelExtType(t))
		}
	}
	res, err := s.ChannelQueryByUser(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServiceChannelQueryByLocationBound(c *gin.Context, params gin_api.BaseServiceChannelQueryByLocationBoundParams) {
	req := &api.ChannelQueryByLocationBoundReq{}
	req.Bound = &api.BoundCoord{
		Ne: &api.PointCoord{},
		Sw: &api.PointCoord{},
	}
	if params.BoundSwLat != nil {
		req.Bound.Sw.Lat = *params.BoundSwLat
	}
	if params.BoundSwLng != nil {
		req.Bound.Sw.Lng = *params.BoundSwLng
	}
	if params.BoundNeLat != nil {
		req.Bound.Ne.Lat = *params.BoundNeLat
	}
	if params.BoundNeLng != nil {
		req.Bound.Ne.Lng = *params.BoundNeLng
	}
	if params.MarkerIds != nil {
		req.MarkerIds = *params.MarkerIds
	}
	res, err := s.ChannelQueryByLocationBound(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePostLoad(c *gin.Context, params gin_api.BaseServicePostLoadParams) {
	req := &api.PostLoadReq{}
	if params.ChannelId != nil {
		req.ChannelId = *params.ChannelId
	}
	if params.Limit != nil {
		req.Limit = *params.Limit
	}
	if params.LastPostId != nil {
		req.LastPostId = *params.LastPostId
	}
	res, err := s.PostLoad(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePostQuery(c *gin.Context, params gin_api.BaseServicePostQueryParams) {
	req := &api.PostQueryReq{}
	if params.ChannelId != nil {
		req.ChannelId = *params.ChannelId
	}
	if params.Limit != nil {
		req.Limit = *params.Limit
	}
	if params.MarkerId != nil {
		req.MarkerId = *params.MarkerId
	}
	res, err := s.PostQuery(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePostQueryByUser(c *gin.Context, params gin_api.BaseServicePostQueryByUserParams) {
	req := &api.PostQueryByUserReq{}
	if params.UserId != nil {
		req.UserId = *params.UserId
	}
	if params.IxnState != nil {
		req.IxnState = *params.IxnState
	}
	if params.IxnEvent != nil {
		req.IxnEvent = *params.IxnEvent
	}
	if params.ExtTypes != nil {
		req.ExtTypes = params.ExtTypes
	}
	res, err := s.PostQueryByUser(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePostCreate(c *gin.Context) {
	req := &api.PostCreateReq{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.PostCreate(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePostUpdate(c *gin.Context) {
	req := &api.PostUpdateReq{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.PostUpdate(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePostDelete(c *gin.Context, params gin_api.BaseServicePostDeleteParams) {
	req := &api.PostDeleteReq{}
	if params.ChannelId != nil {
		req.ChannelId = *params.ChannelId
	}
	if params.PostId != nil {
		req.PostId = *params.PostId
	}
	res, err := s.PostDelete(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) BaseServicePostReact(c *gin.Context) {
	req := &api.PostReactReq{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		s.putGinError(c, EM_CommonFail_BadRequest)
		return
	}
	res, err := s.PostReact(withGinContext(c), req)
	if err != nil {
		s.putGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
