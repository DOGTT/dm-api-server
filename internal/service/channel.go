package service

import (
	"context"

	base_api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/fds"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/davecgh/go-spew/spew"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func validChannelCreateRequest(req *base_api.ChannelCreateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetChannel() == nil {
		return EM_CommonFail_BadRequest.PutDesc("Channel is required")
	}
	c := req.GetChannel()
	if c.GetTitle() == "" {
		return EM_CommonFail_BadRequest.PutDesc("title is required")
	}
	// ..
	return nil
}

func (s *Service) ChannelTypeList(ctx context.Context, req *base_api.ChannelTypeListReq) (res *base_api.ChannelTypeListRes, err error) {
	res = &base_api.ChannelTypeListRes{}
	data, err := s.data.ListChannelTypeInfo(ctx)
	if err != nil {
		return
	}
	res.ChannelTypes = make([]*base_api.ChannelTypeInfo, len(data))
	for i, v := range data {
		res.ChannelTypes[i] = &base_api.ChannelTypeInfo{
			Id:             uint32(v.Id),
			Name:           v.Name,
			CoverageRadius: int32(v.CoverageRadius),
			ThemeColor:     v.ThemeColor,
			CreatedAt:      v.CreatedAt.UnixMilli(),
			UpdatedAt:      v.UpdatedAt.UnixMilli(),
		}
	}
	return
}

func (s *Service) ChannelCreate(ctx context.Context, req *base_api.ChannelCreateReq) (res *base_api.ChannelCreateRes, err error) {
	// valid
	if err = validChannelCreateRequest(req); err != nil {
		return
	}
	res = &base_api.ChannelCreateRes{}
	tc := getClaimFromContext(ctx)
	// TODO: 检查类型，检查坐标是否太近
	ch := req.GetChannel()
	ChannelInfo := &rds.ChannelInfo{
		UUID:     utils.GenShortenUUID(),
		TypeId:   uint16(ch.GetTypeId()),
		UId:      tc.UID,
		Title:    ch.GetTitle(),
		LngLat:   rds.PointCoordToGeometry(ch.GetLocation().GetLngLat()),
		AvatarId: ch.GetAvatar().GetUuid(),
		Intro:    ch.GetIntro(),
		PoiDetail: rds.PoiDetail{
			Address: ch.GetLocation().GetAddress(),
		},
	}
	if err = s.data.CreateChannelInfo(ctx, ChannelInfo); err != nil {
		return
	}
	res.Channel, err = s.convertToChannelInfo(ctx, ChannelInfo)
	if err != nil {
		return
	}
	return
}

func (s *Service) ChannelDelete(ctx context.Context, req *base_api.ChannelDeleteReq) (res *base_api.ChannelDeleteRes, err error) {
	res = &base_api.ChannelDeleteRes{}
	// tc := getClaimFromContext(ctx)
	// TODO，权限检查
	if err = s.data.DeleteChannelInfo(ctx, req.GetUuid()); err != nil {
		return
	}
	return
}

func validChannelUpdateRequest(req *base_api.ChannelUpdateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetChannel() == nil {
		return EM_CommonFail_BadRequest.PutDesc("Channel is required")
	}
	po := req.GetChannel()
	if po.GetTitle() == "" {
		return EM_CommonFail_BadRequest.PutDesc("title is required")
	}
	// ..
	return nil
}

func (s *Service) ChannelUpdate(ctx context.Context, req *base_api.ChannelUpdateReq) (res *base_api.ChannelUpdateRes, err error) {
	// valid
	if err = validChannelUpdateRequest(req); err != nil {
		return
	}
	log.Ctx(ctx).Debug("update request", zap.String("req", spew.Sdump(req)))
	res = &base_api.ChannelUpdateRes{}
	// TODO，权限检查
	ch := req.GetChannel()
	ChannelInfo := &rds.ChannelInfo{
		UUID:     ch.GetUuid(),
		Title:    ch.GetTitle(),
		AvatarId: ch.GetAvatar().GetUuid(),
		Intro:    ch.GetIntro(),
	}
	if err = s.data.UpdateChannelInfo(ctx, ChannelInfo); err != nil {
		return
	}
	return
}

func (s *Service) ChannelDetailQueryById(ctx context.Context, req *base_api.ChannelDetailQueryByIdReq) (res *base_api.ChannelDetailQueryByIdRes, err error) {
	var cInfo *rds.ChannelInfo
	cInfo, err = s.data.GetChannelInfo(ctx, req.GetUuid())
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res = &base_api.ChannelDetailQueryByIdRes{}
	res.Channel, err = s.convertToChannelInfo(ctx, cInfo)
	if err != nil {
		return
	}
	// 异步添加访问信息
	go func() {
		if _, err := s.data.IncreaseChannelViewCount(ctx, cInfo.UUID); err != nil {
			log.Ctx(ctx).Error("increase Channel view count fail", zap.Error(err))
		}
	}()
	return
}

func (s *Service) convertToChannelInfo(ctx context.Context, pInfo *rds.ChannelInfo) (res *base_api.ChannelInfo, err error) {
	res = &base_api.ChannelInfo{
		Uuid:   pInfo.UUID,
		Uid:    pInfo.UId,
		TypeId: uint32(pInfo.TypeId),
		Title:  pInfo.Title,
		Avatar: &base_api.MediaInfo{
			Uuid: pInfo.AvatarId,
		},
		Intro: pInfo.Intro,
		Location: &base_api.LocationInfo{
			LngLat:  rds.PointCoordFromGeometry(pInfo.LngLat),
			Address: pInfo.PoiDetail.Address,
		},
	}
	if !pInfo.CreatedAt.IsZero() {
		res.CreatedAt = pInfo.CreatedAt.UnixMilli()
	}
	if !pInfo.UpdatedAt.IsZero() {
		res.UpdatedAt = pInfo.UpdatedAt.UnixMilli()
	}
	// set stats
	res.Stats = &base_api.ChannelStats{
		ViewsCnt:    int32(pInfo.Stats.ViewsCnt),
		LikesCnt:    int32(pInfo.Stats.LikesCnt),
		MarksCnt:    int32(pInfo.Stats.MarksCnt),
		CommentsCnt: int32(pInfo.Stats.CommentsCnt),
	}
	if !pInfo.Stats.LastView.IsZero() {
		res.Stats.LastView = pInfo.Stats.LastView.UnixMilli()
	}
	if !pInfo.Stats.LastMark.IsZero() {
		res.Stats.LastMark = pInfo.Stats.LastMark.UnixMilli()
	}

	if res.Avatar != nil {
		res.Avatar.GetUrl, err = s.data.GenerateGetPresignedURL(ctx,
			fds.BucketNameChannelImage, res.Avatar.GetUuid(), utils.TokenExpireDuration)
		if err != nil {
			err = EM_CommonFail_Internal.PutDesc(err.Error())
			return
		}
	}
	return
}

func (s *Service) ChannelFullQueryById(ctx context.Context, req *base_api.ChannelFullQueryByIdReq) (res *base_api.ChannelFullQueryByIdRes, err error) {
	// TODO
	return
}

func (s *Service) ChannelBaseQueryByBound(ctx context.Context, req *base_api.ChannelBaseQueryByBoundReq) (res *base_api.ChannelBaseQueryByBoundRes, err error) {
	var pofoList []*rds.ChannelInfo
	log.Ctx(ctx).Debug("query request", zap.String("req", spew.Sdump(req)))
	pofoList, err = s.data.BatchQueryChannelInfoListByBound(ctx,
		utils.ConvertToUintSlice(req.GetTypeIds()), req.GetBound())
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	log.Ctx(ctx).Debug("query result", zap.String("pofoList", spew.Sdump(pofoList)))
	res = &base_api.ChannelBaseQueryByBoundRes{
		Channels: make([]*base_api.ChannelInfo, len(pofoList)),
	}
	for i, pofo := range pofoList {
		res.Channels[i], err = s.convertToChannelInfo(ctx, pofo)
		if err != nil {
			return
		}
	}
	return
}

func (s *Service) ChannelInteraction(ctx context.Context, req *base_api.ChannelInteractionReq) (res *base_api.ChannelInteractionRes, err error) {
	tc := getClaimFromContext(ctx)
	res = new(base_api.ChannelInteractionRes)
	err = s.data.CreateChannelIxnRecordWithCount(ctx, &rds.UserChannelIxnRecord{
		ChannelUUID: req.GetUuid(),
		IntType:     rds.InxType(req.GetIxnType()),
		PId:         tc.PID,
		UId:         tc.UID,
	})
	return
}

func (s *Service) ChannelComment(ctx context.Context, req *base_api.ChannelCommentReq) (res *base_api.ChannelCommentRes, err error) {
	// TODO
	tc := getClaimFromContext(ctx)
	// 查询Channel信息，检查是否存在
	ChannelUUID := req.GetComment().GetRootUuid()
	if err = s.data.ExistChannelInfo(ctx, ChannelUUID); err != nil {
		return
	}
	res = new(base_api.ChannelCommentRes)

	err = s.data.CreateChannelIxnRecordWithCount(ctx, &rds.UserChannelIxnRecord{
		ChannelUUID: ChannelUUID,
		IntType:     rds.InxTypeComment,
		PId:         tc.PID,
		UId:         tc.UID,
	})

	return
}
