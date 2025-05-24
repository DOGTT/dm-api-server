package service

import (
	"context"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) ChannelGet(ctx context.Context, req *api.ChannelGetReq) (res *api.ChannelGetRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = new(api.ChannelGetRes)
	channel, err := s.data.GetChannelFullInfo(ctx, utils.StrToUint64(req.GetChannelId()))
	if err != nil {
		err = putDescByDBErr(err)
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

func (s *Service) ChannelQueryByLocationBound(ctx context.Context, req *api.ChannelQueryByLocationBoundReq) (res *api.ChannelQueryByLocationBoundRes, err error) {
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
	res = &api.ChannelQueryByLocationBoundRes{
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

func (s *Service) ChannelQueryByUser(ctx context.Context, req *api.ChannelQueryByUserReq) (res *api.ChannelQueryByUserRes, err error) {
	log.D(ctx, "request in", "req", req)
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
		Id:       utils.GenSnowflakeId(),
		TypeId:   ch.GetTypeId(),
		UId:      tc.UId,
		Title:    ch.GetTitle(),
		LngLat:   utils.PointCoordToGeometry(ch.GetLocation().GetLngLat()),
		AvatarId: ch.GetAvatar().GetUuid(),
		Intro:    ch.GetIntro(),
		PoiDetail: rds.PoiDetail{
			Address: ch.GetLocation().GetAddress(),
		},
	}
	var postInit *rds.PostInfo
	if req.GetPostInit() != nil {
		postInit = &rds.PostInfo{
			Id:      utils.GenSnowflakeId(),
			UId:     tc.UId,
			Content: req.GetPostInit().GetContent(),
		}
	}
	if err = s.dbCreateChannel(ctx, channel, postInit); err != nil {
		log.E(ctx, "create channel error", err)
		err = putDescByDBErr(err)
		return
	}
	res.Channel, err = s.convertToChannelInfo(ctx, channel)
	if err != nil {
		log.E(ctx, "convert channel info error", err)
		return
	}
	return
}

func (s *Service) dbCreateChannel(ctx context.Context, channel *rds.ChannelInfo, initPost *rds.PostInfo) (err error) {
	tx, err := s.data.NewTransaction(ctx)
	if err != nil {
		err = putDescByDBErr(err)
		return
	}
	defer func() {
		if err != nil {
			err = putDescByDBErr(err)
			if rbErr := tx.Rollback(); rbErr != nil {
				log.E(ctx, "rollback error", rbErr)
			}
		}
	}()
	channel.Set = rds.ChannelSet{
		Id: channel.Id,
	}
	channel.Stats = rds.ChannelStats{
		Id: channel.Id,
	}
	err = s.data.CreateChannelInfo(ctx, channel)
	if err != nil {
		return
	}
	if initPost == nil {
		return
	}
	initPost.RootId = channel.Id
	err = s.data.CreatePostInfo(ctx, initPost)
	if err != nil {
		return
	}
	return err
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
		err = putDescByDBErr(err)
		return
	}
	return
}

func (s *Service) validChannelPermission(ctx context.Context, tc *utils.TokenClaims, channalId uint64) error {
	uid, err := s.data.GetChannelCreatorId(ctx, channalId)
	if err != nil {
		log.E(ctx, "get channel creater id error", err)
		err = putDescByDBErr(err)
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
	chanData := &rds.ChannelInfo{
		Id:       chanId,
		Title:    ch.GetTitle(),
		AvatarId: ch.GetAvatar().GetUuid(),
		Intro:    ch.GetIntro(),
	}
	if err = s.data.UpdateChannelInfo(ctx, chanData); err != nil {
		log.E(ctx, "update channel info error", err)
		err = putDescByDBErr(err)
		return
	}
	chanData, err = s.data.GetChannelInfo(ctx, chanId)
	if err != nil {
		log.E(ctx, "get channel info error", err)
		err = putDescByDBErr(err)
		return
	}
	res.Channel, err = s.convertToChannelInfo(ctx, chanData)
	if err != nil {
		log.E(ctx, "convert channel info error", err)
		return
	}
	return
}

func (s *Service) convertToChannelStats(ctx context.Context, stat *rds.ChannelStats) *api.ChannelStats {
	statsApi := &api.ChannelStats{
		ViewsCnt: int32(stat.ViewsCnt),
		StarsCnt: int32(stat.StarsCnt),
		PeeCnt:   int32(stat.PeeCnt),
		PostsCnt: int32(stat.PostsCnt),
	}
	if !stat.LastStarAt.IsZero() {
		statsApi.LastStarAt = stat.LastStarAt.UnixMilli()
	}
	if !stat.LastPostAt.IsZero() {
		statsApi.LastPostAt = stat.LastPostAt.UnixMilli()
	}
	if !stat.LastPeeAt.IsZero() {
		statsApi.LastPeeAt = stat.LastPeeAt.UnixMilli()
	}
	if !stat.UpdatedAt.IsZero() {
		statsApi.UpdatedAt = stat.UpdatedAt.UnixMilli()
	}
	return statsApi
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
		res.Avatar.GetUrl, err = s.data.GenerateGetPresignedURLByMediaInfo(ctx, res.Avatar, 0)
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
			status      = inxReq.GetInxState()
			petEvent    = inxReq.GetPetInxEvent()
			rdsStatItem rds.ChannelStatsType
		)
		if petEvent != api.ChannelPetIxnEvent_CPIE_UNSPECIFIED {
			switch petEvent {
			case api.ChannelPetIxnEvent_PEE:
				rdsStatItem = rds.ChannelStatsPee
			}
			err = s.data.ChannelStatsIncrease(ctx, channelId, rdsStatItem)
			if err != nil {
				log.E(ctx, "event stats increase error", err, "inx_req", inxReq)
			}
		}
		if status != api.ChannelUserIxnState_CUIS_UNSPECIFIED {
			switch status {
			case api.ChannelUserIxnState_STAR:
				rdsStatItem = rds.ChannelStatsStar
			}
			if inxReq.GetInxStateUndo() == api.InxUndoType_UNDO {
				err = s.data.ChannelStatsDecrease(ctx, channelId, rdsStatItem)
			} else {
				err = s.data.ChannelStatsIncrease(ctx, channelId, rdsStatItem)
			}
			if err != nil {
				log.E(ctx, "status stats change error", err,
					"inx_req", inxReq, "is_undo", inxReq.GetInxStateUndo().String())
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
