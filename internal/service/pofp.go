package service

import (
	"context"

	base_api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/davecgh/go-spew/spew"
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

func (s *Service) PofpTypeList(ctx context.Context, req *base_api.PofpTypeListReq) (res *base_api.PofpTypeListResp, err error) {
	res = &base_api.PofpTypeListResp{}
	data, err := s.data.ListPofpTypeInfo(ctx)
	if err != nil {
		return
	}
	res.PofpTypes = make([]*base_api.PofpTypeInfo, len(data))
	for i, v := range data {
		res.PofpTypes[i] = &base_api.PofpTypeInfo{
			Id:             uint32(v.Id),
			Name:           v.Name,
			CoverageRadius: int32(v.CoverageRadius),
			ThemeColor:     v.ThemeColor,
			CreatedAt:      timestamppb.New(v.CreatedAt),
			UpdatedAt:      timestamppb.New(v.UpdatedAt),
		}
	}
	return
}

func (s *Service) PofpCreate(ctx context.Context, req *base_api.PofpCreateReq) (res *base_api.PofpCreateResp, err error) {
	// valid
	if err = validPofpCreateRequest(req); err != nil {
		return
	}
	res = &base_api.PofpCreateResp{}
	tc := getClaimFromContext(ctx)
	// TODO: 检查类型，检查坐标是否太近
	po := req.GetPofp()
	pofpInfo := &rds.PofpInfo{
		UUID:    utils.GenShortenUUID(),
		TypeId:  uint(po.GetTypeId()),
		PId:     tc.PID,
		Title:   po.GetTitle(),
		LngLat:  rds.PointCoordToGeometry(po.GetLngLat()),
		Photos:  nil,
		Content: po.GetContent(),
		Address: po.GetAddress(),
		PoiId:   po.GetPoiId(),
	}
	if err = s.data.CreatePofpInfo(ctx, pofpInfo); err != nil {
		return
	}
	res.Pofp, err = s.convertToPofpInfo(ctx, pofpInfo)
	if err != nil {
		return
	}
	return
}

func (s *Service) PofpDelete(ctx context.Context, req *base_api.PofpDeleteReq) (res *base_api.PofpDeleteResp, err error) {
	res = &base_api.PofpDeleteResp{}
	tc := getClaimFromContext(ctx)
	// TODO，权限检查
	if err = s.data.DeletePofpInfo(ctx, req.GetUuid(), tc.PID); err != nil {
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
	log.Ctx(ctx).Debug("update request", zap.String("req", spew.Sdump(req)))
	res = &base_api.PofpUpdateResp{}
	// TODO，权限检查
	po := req.GetPofp()
	pofpInfo := &rds.PofpInfo{
		UUID:    po.GetUuid(),
		Title:   po.GetTitle(),
		Photos:  nil,
		Content: po.GetContent(),
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
		Pid:         pInfo.PId,
		TypeId:      uint32(pInfo.TypeId),
		Title:       pInfo.Title,
		LngLat:      rds.PointCoordFromGeometry(pInfo.LngLat),
		Photos:      nil,
		Content:     pInfo.Content,
		Address:     pInfo.Address,
		PoiId:       pInfo.PoiId,
		ViewsCnt:    int32(pInfo.ViewsCnt),
		LikesCnt:    int32(pInfo.LikesCnt),
		CommentsCnt: int32(pInfo.CommentsCnt),
	}
	if !pInfo.LastView.IsZero() {
		res.LastView = timestamppb.New(pInfo.LastView)
	}
	if !pInfo.LastMark.IsZero() {
		res.LastMark = timestamppb.New(pInfo.LastMark)
	}
	if !pInfo.CreatedAt.IsZero() {
		res.CreatedAt = timestamppb.New(pInfo.CreatedAt)
	}
	if !pInfo.UpdatedAt.IsZero() {
		res.UpdatedAt = timestamppb.New(pInfo.UpdatedAt)
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
	log.Ctx(ctx).Debug("query request", zap.String("req", spew.Sdump(req)))
	pofoList, err = s.data.BatchQueryPofpInfoListByBound(ctx,
		utils.ConvertToUintSlice(req.GetTypeIds()), req.GetBound())
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	log.Ctx(ctx).Debug("query result", zap.String("pofoList", spew.Sdump(pofoList)))
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
