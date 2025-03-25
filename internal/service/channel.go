package service

import (
	"context"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/davecgh/go-spew/spew"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func (s *Service) ChannelFullQueryById(ctx context.Context, req *api.ChannelFullQueryByIdReq) (res *api.ChannelFullQueryByIdRes, err error) {
	log := log.Ctx(ctx)
	log.Debug("query request", zap.String("req", spew.Sdump(req)))
	res = new(api.ChannelFullQueryByIdRes)
	channel, err := s.data.GetChannelFullInfo(ctx, utils.StrToUint64(req.GetChId()))
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		log.Error("get query error", zap.Error(err))
		return
	}
	res.Channel, err = s.convertToChannelInfo(ctx, channel)
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		log.Error("convert error", zap.Error(err))
		return
	}
	return
}

func (s *Service) ChannelBaseQueryByBound(ctx context.Context, req *api.ChannelBaseQueryByBoundReq) (res *api.ChannelBaseQueryByBoundRes, err error) {
	log := log.Ctx(ctx)
	log.Debug("base query request", zap.String("req", spew.Sdump(req)))
	channels, err := s.data.ListChannelInfo(ctx,
		&rds.ChannelFilter{
			TypeIDs:    req.GetTypeIds(),
			BoundCoord: req.GetBound(),
		})
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		log.Error("list query error", zap.Error(err))
		return
	}
	log.Debug("query result", zap.String("channels", spew.Sdump(channels)))
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
	log := log.Ctx(ctx)
	log.Debug("query request", zap.String("req", spew.Sdump(req)))
	res = new(api.ChannelInxRes)
	// err = s.data.CreateChannelIxnRecordWithCount(ctx, &rds.UserChannelIxnRecord{
	// 	ChannelUUID: req.GetUuid(),
	// 	IntType:     rds.InxType(req.GetIxnType()),
	// 	PId:         tc.PID,
	// 	UId:         tc.UID,
	// })
	return
}

func (s *Service) ChannelComment(ctx context.Context, req *api.ChannelCommentReq) (res *api.ChannelCommentRes, err error) {
	// TODO
	// tc := getClaimFromContext(ctx)
	// 查询Channel信息，检查是否存在
	// ChannelUUID := req.GetComment().GetRootUuid()
	// if err = s.data.ExistChannelInfo(ctx, ChannelUUID); err != nil {
	// 	return
	// }
	res = new(api.ChannelCommentRes)

	// err = s.data.CreateChannelIxnRecordWithCount(ctx, &rds.UserChannelIxnRecord{
	// 	ChannelUUID: ChannelUUID,
	// 	IntType:     rds.InxTypeComment,
	// 	PId:         tc.PID,
	// 	UId:         tc.UID,
	// })

	return
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

func (s *Service) ChannelTypeList(ctx context.Context, req *api.ChannelTypeListReq) (res *api.ChannelTypeListRes, err error) {
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
	log := log.Ctx(ctx)
	// valid
	if err = validChannelCreateRequest(req); err != nil {
		return
	}
	res = &api.ChannelCreateRes{}
	tc := getClaimFromContext(ctx)
	// TODO: 检查类型，检查坐标是否太近
	ch := req.GetChannel()
	ChannelInfo := &rds.ChannelInfo{
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
	}
	if err = s.data.CreateChannelInfo(ctx, ChannelInfo); err != nil {
		return
	}
	res.Channel, err = s.convertToChannelInfo(ctx, ChannelInfo)
	if err != nil {
		log.Error("convert channel info error", zap.Error(err))
		return
	}
	return
}

func (s *Service) ChannelDelete(ctx context.Context, req *api.ChannelDeleteReq) (res *api.ChannelDeleteRes, err error) {
	log := log.Ctx(ctx)
	res = &api.ChannelDeleteRes{}
	var (
		chanId = utils.StrToUint64(req.GetChId())
		tc     = getClaimFromContext(ctx)
	)
	if err = s.validChannelPermission(ctx, tc, chanId); err != nil {
		log.Error("valid channel permission error", zap.Error(err))
		return
	}
	if err = s.data.DeleteChannelInfo(ctx, utils.StrToUint64(req.GetChId())); err != nil {
		log.Error("delete channel error", zap.Error(err))
		return
	}
	return
}

func (s *Service) validChannelUpdateRequest(req *api.ChannelUpdateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetChannel() == nil {
		return EM_CommonFail_BadRequest.PutDesc("channel is required")
	}
	po := req.GetChannel()
	if po.GetTitle() == "" {
		return EM_CommonFail_BadRequest.PutDesc("title is required")
	}
	// ..
	return nil
}

func (s *Service) validChannelPermission(ctx context.Context, tc *utils.TokenClaims, channalId uint64) error {
	uid, err := s.data.GetChannelCreaterId(ctx, channalId)
	if err != nil {
		return err
	}
	if uid != tc.UId {
		return EM_CommonFail_Forbidden.PutDesc("user has no permission")
	}
	return nil
}

func (s *Service) ChannelUpdate(ctx context.Context, req *api.ChannelUpdateReq) (res *api.ChannelUpdateRes, err error) {
	log := log.Ctx(ctx)
	log.Debug("update request", zap.String("req", spew.Sdump(req)))
	// valid
	if err = s.validChannelUpdateRequest(req); err != nil {
		log.Debug("valid error", zap.Error(err))
		return
	}
	res = new(api.ChannelUpdateRes)
	var (
		ch     = req.GetChannel()
		chanId = utils.StrToUint64(ch.GetId())
		tc     = getClaimFromContext(ctx)
	)
	if err = s.validChannelPermission(ctx, tc, chanId); err != nil {
		log.Error("valid channel permission error", zap.Error(err))
		return
	}
	if err = s.data.UpdateChannelInfo(ctx, &rds.ChannelInfo{
		Id:       utils.StrToUint64(ch.GetId()),
		Title:    ch.GetTitle(),
		AvatarId: ch.GetAvatar().GetUuid(),
		Intro:    ch.GetIntro(),
	}); err != nil {
		log.Error("update channel info error", zap.Error(err))
		return
	}
	return
}

func (s *Service) convertToChannelStats(ctx context.Context, stat *rds.ChannelStats) *api.ChannelStats {
	return &api.ChannelStats{
		ViewsCnt:    int32(stat.ViewsCnt),
		StarsCnt:    int32(stat.StarsCnt),
		PetMarksCnt: int32(stat.PetMarksCnt),
		PostsCnt:    int32(stat.PostsCnt),
		LastView:    stat.LastView.UnixMilli(),
		LastPetMark: stat.LastPetMark.UnixMilli(),
		LastInx:     stat.LastInx.UnixMilli(),
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
