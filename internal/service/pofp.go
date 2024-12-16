package service

import (
	"context"

	base_api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func validPofpCreateRequest(req *base_api.PofpCreateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetPofp() == nil {
		return EM_CommonFail_BadRequest.PutDesc("pofp is required")
	}
	po := req.GetPofp()
	if po.GetTitle() == "" {
		return EM_CommonFail_BadRequest.PutDesc("title is required")
	}
	// ..
	return nil
}

func (s *Service) PofpCreate(ctx context.Context, req *base_api.PofpCreateReq) (res *base_api.PofpCreateResp, err error) {
	// valid
	if err = validPofpCreateRequest(req); err != nil {
		return
	}
	res = &base_api.PofpCreateResp{}
	tc := getClaimFromContext(ctx)
	po := req.GetPofp()
	pofpInfo := &rds.PofpInfo{
		UUID:    utils.GenShortenUUID(),
		TypeId:  uint(po.GetTypeId()),
		PId:     tc.PID,
		Title:   po.GetTitle(),
		LatLng:  rds.PointCoordToGeometry(po.GetLatLng()),
		Photos:  nil,
		Content: po.GetContent(),
		Address: po.GetAddress(),
		PoiId:   po.GetPoiId(),
	}
	if err = s.data.CreatePofpInfo(ctx, pofpInfo); err != nil {
		return
	}
	return
}

func (s *Service) PofpDelete(ctx context.Context, req *base_api.PofpDeleteReq) (res *base_api.PofpDeleteResp, err error) {
	res = &base_api.PofpDeleteResp{}
	// tc := getClaimFromContext(ctx)
	if err = s.data.DeletePofpInfo(ctx, req.GetUuid()); err != nil {
		return
	}
	return
}

func validPofpUpdateRequest(req *base_api.PofpUpdateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetPofp() == nil {
		return EM_CommonFail_BadRequest.PutDesc("pofp is required")
	}
	po := req.GetPofp()
	if po.GetTitle() == "" {
		return EM_CommonFail_BadRequest.PutDesc("title is required")
	}
	// ..
	return nil
}

func (s *Service) PofpUpdate(ctx context.Context, req *base_api.PofpUpdateReq) (res *base_api.PofpUpdateResp, err error) {
	// valid
	if err = validPofpUpdateRequest(req); err != nil {
		return
	}
	po := req.GetPofp()
	pofpInfo := &rds.PofpInfo{
		UUID:    utils.GenShortenUUID(),
		TypeId:  uint(po.GetTypeId()),
		Title:   po.GetTitle(),
		LatLng:  rds.PointCoordToGeometry(po.GetLatLng()),
		Photos:  nil,
		Content: po.GetContent(),
		Address: po.GetAddress(),
		PoiId:   po.GetPoiId(),
	}
	if err = s.data.UpdatePofpInfo(ctx, pofpInfo); err != nil {
		return
	}
	return
}

func (s *Service) PofpDetailQueryById(ctx context.Context, req *base_api.PofpDetailQueryByIdReq) (res *base_api.PofpDetailQueryByIdResp, err error) {
	var pofoInfo *rds.PofpInfo
	pofoInfo, err = s.data.GetPofpInfo(ctx, req.GetUuid())
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res = &base_api.PofpDetailQueryByIdResp{}
	res.Pofp, err = s.convertToPofpInfo(ctx, pofoInfo)
	if err != nil {
		return
	}
	// 异步添加访问信息
	go func() {
		if _, err := s.data.IncreasePofpViewCount(ctx, pofoInfo.UUID); err != nil {
			log.Ctx(ctx).Error("increase pofp view count fail", zap.Error(err))
		}
	}()
	return
}

func (s *Service) convertToPofpInfo(ctx context.Context, pInfo *rds.PofpInfo) (res *base_api.PofpInfo, err error) {
	res = &base_api.PofpInfo{
		Uuid:        pInfo.UUID,
		TypeId:      uint32(pInfo.TypeId),
		Title:       pInfo.Title,
		LatLng:      rds.PointCoordFromGeometry(pInfo.LatLng),
		Photos:      nil,
		Content:     pInfo.Content,
		Address:     pInfo.Address,
		PoiId:       pInfo.PoiId,
		ViewsCnt:    int32(pInfo.ViewsCnt),
		LikesCnt:    int32(pInfo.LikesCnt),
		CommentsCnt: int32(pInfo.CommentsCnt),
		LastView:    timestamppb.New(pInfo.LastView),
		LastMark:    timestamppb.New(pInfo.LastMark),

		CreatedAt: timestamppb.New(pInfo.CreatedAt),
		UpdatedAt: timestamppb.New(pInfo.UpdatedAt),
	}
	// res.Pets[i].Avatar, err = s.data.GeneratePresignedURL(ctx,
	// 	fds.BucketNameAvatar, pet.AvatarId, utils.TokenExpireDuration)
	// if err != nil {
	// 	err = EM_CommonFail_Internal.PutDesc(err.Error())
	// 	return
	// }
	return
}

func (s *Service) PofpFullQueryById(ctx context.Context, req *base_api.PofpFullQueryByIdReq) (res *base_api.PofpFullQueryByIdResp, err error) {
	// TODO
	return
}

func (s *Service) PofpBaseQueryByBound(ctx context.Context, req *base_api.PofpBaseQueryByBoundReq) (res *base_api.PofpBaseQueryByBoundResp, err error) {
	var pofoList []*rds.PofpInfo
	pofoList, err = s.data.BatchQueryPofpInfoListByBound(ctx,
		utils.ConvertToUintSlice(req.GetTypeIds()), req.GetBound())
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res = &base_api.PofpBaseQueryByBoundResp{
		Pofps: make([]*base_api.PofpInfo, len(pofoList)),
	}
	for i, pofo := range pofoList {
		res.Pofps[i], err = s.convertToPofpInfo(ctx, pofo)
		if err != nil {
			return
		}
	}
	return
}

func (s *Service) PofpInteraction(ctx context.Context, req *base_api.PofpInteractionReq) (res *base_api.PofpInteractionResp, err error) {
	tc := getClaimFromContext(ctx)
	err = s.data.CreatePofpIxnRecordWithCount(ctx, &rds.UserPofpIxnRecord{
		PofpUUID: req.GetUuid(),
		IntType:  rds.InxType(req.GetIxnType()),
		PId:      tc.PID,
		UId:      tc.UID,
	})
	return
}

func (s *Service) PofpComment(ctx context.Context, req *base_api.PofpCommentReq) (res *base_api.PofpCommentResp, err error) {
	// TODO
	return
}
