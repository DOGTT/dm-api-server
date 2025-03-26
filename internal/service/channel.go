package service

import (
	"context"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) ChannelFullQueryById(ctx context.Context, req *api.ChannelFullQueryByIdReq) (res *api.ChannelFullQueryByIdRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = new(api.ChannelFullQueryByIdRes)
	channel, err := s.data.GetChannelFullInfo(ctx, utils.StrToUint64(req.GetChanId()))
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		log.E(ctx, "data get channel error", err)
		return
	}
	// 添加访问记录
	s.asyncIncreaseChannelViews(ctx, channel.Id)
	res.Channel, err = s.convertToChannelInfo(ctx, channel)
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		log.E(ctx, "convert error", err)
		return
	}
	return
}

func (s *Service) ChannelBaseQueryByBound(ctx context.Context, req *api.ChannelBaseQueryByBoundReq) (res *api.ChannelBaseQueryByBoundRes, err error) {
	log.D(ctx, "request in", "req", req)
	channels, err := s.data.ListChannelInfo(ctx,
		&rds.ChannelFilter{
			TypeIDs:    req.GetTypeIds(),
			BoundCoord: req.GetBound(),
		})
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		log.E(ctx, "list query error", err)
		return
	}
	log.D(ctx, "query result", "channels", channels)
	res = &api.ChannelBaseQueryByBoundRes{
		Channels: make([]*api.ChannelInfo, len(channels)),
	}
	for i, channel := range channels {
		// base query should not return avatar id
		channel.AvatarId = ""
		res.Channels[i], err = s.convertToChannelInfo(ctx, channel)
		if err != nil {
			return
		}
	}
	return
}

func (s *Service) ChannelInx(ctx context.Context, req *api.ChannelInxReq) (res *api.ChannelInxRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = new(api.ChannelInxRes)
	tc := utils.GetClaimFromContext(ctx)
	chanId := utils.StrToUint64(req.GetChanId())
	_, err = s.data.GetChannelCreaterId(ctx, chanId)
	// 查询Channel信息，检查是否存在
	if err != nil {
		log.E(ctx, "channel not exist", err)
		err = EM_CommonFail_DBError.PutDesc(err.Error())
		return
	}
	// 状态互动基于用户
	if req.GetIxnState() != 0 {
		dbIn := &rds.UserChannelIxnState{
			UId:       tc.UId,
			ChannelId: chanId,
			IxnState:  req.GetIxnState(),
		}
		if req.GetIsStateUndo() {
			err = s.data.DeleteUserChannelIxnState(ctx, dbIn)
		} else {
			err = s.data.CreateUserChannelIxnState(ctx, dbIn)
		}
		if err != nil {
			log.E(ctx, "create user channel ixn state error", err)
			err = EM_CommonFail_DBError.PutDesc(err.Error())
			return
		}
	}
	// 事件互动基于用户和爱宠
	if req.GetIxnState() != 0 {
		// 加载爱宠id
		var userInfo *rds.UserInfo
		userInfo, err = s.data.GetUserInfoByID(ctx, tc.UId)
		if err != nil {
			log.E(ctx, "user load error", err)
			err = EM_CommonFail_DBError.PutDesc(err.Error())
		}
		pids := userInfo.GetPIDs()
		events := make([]*rds.UserPetChannelIxnEvent, len(pids))
		for i, pid := range userInfo.GetPIDs() {
			events[i] = &rds.UserPetChannelIxnEvent{
				UId:       tc.UId,
				PId:       pid,
				ChannelId: chanId,
				IxnEvent:  req.GetIxnEvent(),
			}
		}
		err = s.data.BatchCreateUserPetChannelIxnEvent(ctx, events)
		if err != nil {
			log.E(ctx, "create user channel ixn event error", err)
			err = EM_CommonFail_DBError.PutDesc(err.Error())
			return
		}
	}
	// 更新统计信息
	s.asyncUpdateChannelStatsByInx(ctx, chanId, req)
	return
}

func (s *Service) ChannelTypeList(ctx context.Context, req *api.ChannelTypeListReq) (res *api.ChannelTypeListRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = &api.ChannelTypeListRes{}
	data, err := s.data.ListChannelTypeInfo(ctx)
	if err != nil {
		return
	}
	res.ChannelTypes = make([]*api.ChannelTypeInfo, len(data))
	for i, v := range data {
		res.ChannelTypes[i] = &api.ChannelTypeInfo{
			Id:             utils.Uint64ToStr(v.Id),
			Name:           v.Name,
			CoverageRadius: int32(v.CoverageRadius),
			ThemeColor:     v.ThemeColor,
			CreatedAt:      v.CreatedAt.UnixMilli(),
			UpdatedAt:      v.UpdatedAt.UnixMilli(),
		}
	}
	return
}

func (s *Service) ChannelCreate(ctx context.Context, req *api.ChannelCreateReq) (res *api.ChannelCreateRes, err error) {
	log.D(ctx, "request in", "req", req)
	// valid
	if err = validChannelCreateRequest(req); err != nil {
		return
	}
	res = &api.ChannelCreateRes{}
	tc := utils.GetClaimFromContext(ctx)
	// TODO: 检查类型，检查坐标是否太近
	ch := req.GetChannel()
	channel := &rds.ChannelInfo{
		Id:       utils.GenSnowflakeID(),
		TypeId:   ch.GetTypeId(),
		UId:      tc.UId,
		Title:    ch.GetTitle(),
		LngLat:   utils.PointCoordToGeometry(ch.GetLocation().GetLngLat()),
		AvatarId: ch.GetAvatar().GetUuid(),
		Intro:    ch.GetIntro(),
		PoiDetail: rds.PoiDetail{
			Address: ch.GetLocation().GetAddress(),
		},
		Stats: rds.ChannelStats{},
		Set:   rds.ChannelSet{},
	}
	if err = s.data.CreateChannelInfo(ctx, channel); err != nil {
		log.E(ctx, "create channel error", err)
		err = EM_CommonFail_DBError.PutDesc(err.Error())
		return
	}
	res.Channel, err = s.convertToChannelInfo(ctx, channel)
	if err != nil {
		log.E(ctx, "convert channel info error", err)
		return
	}
	return
}

func (s *Service) ChannelDelete(ctx context.Context, req *api.ChannelDeleteReq) (res *api.ChannelDeleteRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = &api.ChannelDeleteRes{}
	var (
		chanId = utils.StrToUint64(req.GetChanId())
		tc     = utils.GetClaimFromContext(ctx)
	)
	if err = s.validChannelPermission(ctx, tc, chanId); err != nil {
		return
	}
	if err = s.data.DeleteChannelInfo(ctx, chanId); err != nil {
		log.E(ctx, "delete channel error", err)
		err = EM_CommonFail_DBError.PutDesc(err.Error())
		return
	}
	return
}

func (s *Service) validChannelPermission(ctx context.Context, tc *utils.TokenClaims, channalId uint64) error {
	uid, err := s.data.GetChannelCreaterId(ctx, channalId)
	if err != nil {
		log.E(ctx, "get channel creater id error", err)
		err = EM_CommonFail_DBError.PutDesc(err.Error())
		return err
	}
	if uid != tc.UId {
		return EM_CommonFail_Forbidden.PutDesc("user has no permission")
	}
	return nil
}

func (s *Service) ChannelUpdate(ctx context.Context, req *api.ChannelUpdateReq) (res *api.ChannelUpdateRes, err error) {
	log.D(ctx, "request in", "req", req)
	if err = validChannelUpdateRequest(req); err != nil {
		return
	}
	res = new(api.ChannelUpdateRes)
	var (
		ch     = req.GetChannel()
		chanId = utils.StrToUint64(ch.GetId())
		tc     = utils.GetClaimFromContext(ctx)
	)
	if err = s.validChannelPermission(ctx, tc, chanId); err != nil {
		return
	}
	if err = s.data.UpdateChannelInfo(ctx, &rds.ChannelInfo{
		Id:       chanId,
		Title:    ch.GetTitle(),
		AvatarId: ch.GetAvatar().GetUuid(),
		Intro:    ch.GetIntro(),
	}); err != nil {
		log.E(ctx, "update channel info error", err)
		err = EM_CommonFail_DBError.PutDesc(err.Error())
		return
	}
	return
}

func (s *Service) convertToChannelStats(ctx context.Context, stat *rds.ChannelStats) *api.ChannelStats {
	return &api.ChannelStats{
		ViewsCnt:   int32(stat.ViewsCnt),
		StarsCnt:   int32(stat.StarsCnt),
		PeeCnt:     int32(stat.PeeCnt),
		PostsCnt:   int32(stat.PostsCnt),
		LastStarAt: stat.LastStarAt.UnixMilli(),
		LastPostAt: stat.LastPostAt.UnixMilli(),
		LastPeeAt:  stat.LastPeeAt.UnixMilli(),
		UpdatedAt:  stat.UpdatedAt.UnixMilli(),
	}
}

func (s *Service) convertToChannelInfo(ctx context.Context, in *rds.ChannelInfo) (res *api.ChannelInfo, err error) {
	res = &api.ChannelInfo{
		Id:     utils.Uint64ToStr(in.Id),
		Uid:    utils.Uint64ToStr(in.UId),
		TypeId: in.TypeId,
		Title:  in.Title,
		Avatar: &api.MediaInfo{
			Uuid: in.AvatarId,
			Type: api.MediaType_CHANNEL_AVA,
		},
		Intro: in.Intro,
		Location: &api.LocationInfo{
			LngLat:  utils.PointCoordFromGeometry(in.LngLat),
			Address: in.PoiDetail.Address,
		},
	}
	if !in.CreatedAt.IsZero() {
		res.CreatedAt = in.CreatedAt.UnixMilli()
	}
	if !in.UpdatedAt.IsZero() {
		res.UpdatedAt = in.UpdatedAt.UnixMilli()
	}
	// set stats
	res.Stats = s.convertToChannelStats(ctx, &in.Stats)
	if res.Avatar != nil && res.Avatar.GetUuid() != "" {
		res.Avatar.GetUrl, err = s.data.GenerateGetPresignedURLByMediaInfo(ctx, res.Avatar)
		if err != nil {
			err = EM_CommonFail_Internal.PutDesc(err.Error())
			return
		}
	}
	return
}

func (s *Service) asyncUpdateChannelStatsByInx(ctx context.Context, channelId uint64, inxReq *api.ChannelInxReq) {
	go func() {
		var (
			err         error
			event       = inxReq.GetIxnEvent()
			status      = inxReq.GetIxnState()
			rdsStatItem rds.ChannelStatsType
		)
		if event != api.ChannelIxnEvent_EVENT_DEFAULT {
			switch event {
			case api.ChannelIxnEvent_PEE:
				rdsStatItem = rds.ChannelStatsPee
			}
			err = s.data.ChannelStatsIncrease(ctx, channelId, rdsStatItem)
			if err != nil {
				log.E(ctx, "event stats increase error", err, "inx_req", inxReq)
			}
		}
		if status != api.ChannelIxnState_STATE_DEFAULT {
			switch status {
			case api.ChannelIxnState_STAR:
				rdsStatItem = rds.ChannelStatsStar
			}
			if inxReq.GetIsStateUndo() {
				err = s.data.ChannelStatsDecrease(ctx, channelId, rdsStatItem)
			} else {
				err = s.data.ChannelStatsIncrease(ctx, channelId, rdsStatItem)
			}
			if err != nil {
				log.E(ctx, "status stats change error", err, "inx_req", inxReq, "is_undo", inxReq.GetIsStateUndo())
			}
		}
	}()
}

func (s *Service) asyncIncreaseChannelViews(ctx context.Context, channelId uint64) {
	go func() {
		err := s.data.ChannelStatsIncrease(ctx, channelId, rds.ChannelStatsView)
		if err != nil {
			log.E(ctx, "view stats increase error", err)
		}
	}()
}

func validChannelCreateRequest(req *api.ChannelCreateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetChannel() == nil {
		return EM_CommonFail_BadRequest.PutDesc("channel is required")
	}
	c := req.GetChannel()
	if c.GetTitle() == "" {
		return EM_CommonFail_BadRequest.PutDesc("title is required")
	}
	// ..
	return nil
}

func validChannelUpdateRequest(req *api.ChannelUpdateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetChannel() == nil {
		return EM_CommonFail_BadRequest.PutDesc("channel is required")
	}
	ch := req.GetChannel()
	if ch.GetTitle() == "" {
		return EM_CommonFail_BadRequest.PutDesc("title is required")
	}
	// ..
	return nil
}
