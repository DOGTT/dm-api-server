package service

import (
	"context"

	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func validPofpCreateRequest(req *grpc_api.PofpCreateReq) error {
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

func (s *Service) PofpCreate(ctx context.Context, req *grpc_api.PofpCreateReq) (res *grpc_api.PofpCreateResp, err error) {
	// valid
	if err = validPofpCreateRequest(req); err != nil {
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
	if err = s.data.CreatePofpInfo(ctx, pofpInfo); err != nil {
		return
	}
	return
}

func (s *Service) PofpDelete(ctx context.Context, req *grpc_api.PofpDeleteReq) (res *grpc_api.PofpDeleteResp, err error) {

	if err = s.data.DeletePofpInfo(ctx, req.GetUuid()); err != nil {
		return
	}
	return
}

func validPofpUpdateRequest(req *grpc_api.PofpUpdateReq) error {
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

func (s *Service) PofpUpdate(ctx context.Context, req *grpc_api.PofpUpdateReq) (res *grpc_api.PofpUpdateResp, err error) {
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

func (s *Service) PofpDetailQueryById(ctx context.Context, req *grpc_api.PofpDetailQueryByIdReq) (res *grpc_api.PofpDetailQueryByIdResp, err error) {
	var pofoInfo *rds.PofpInfo
	pofoInfo, err = s.data.GetPofpInfo(ctx, req.GetUuid())
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res = &grpc_api.PofpDetailQueryByIdResp{}
	res.Pofp, err = s.convertToPofpInfo(ctx, pofoInfo)
	if err != nil {
		return
	}
	// 添加访问信息
	// tokenInfo := getCliamFromContext(ctx)

	return
}

func (s *Service) convertToPofpInfo(ctx context.Context, pInfo *rds.PofpInfo) (res *grpc_api.PofpInfo, err error) {
	res = &grpc_api.PofpInfo{
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

func (s *Service) PofpFullQueryById(ctx context.Context, req *grpc_api.PofpFullQueryByIdReq) (res *grpc_api.PofpFullQueryByIdResp, err error) {
	// TODO
	return
}

func (s *Service) PofpBaseQueryByBound(ctx context.Context, req *grpc_api.PofpBaseQueryByBoundReq) (res *grpc_api.PofpBaseQueryByBoundResp, err error) {
	// TODO
	return
}

func (s *Service) PofpInteraction(ctx context.Context, req *grpc_api.PofpInteractionReq) (res *grpc_api.PofpInteractionResp, err error) {
	return
}

func (s *Service) PofpComment(ctx context.Context, req *grpc_api.PofpCommentReq) (res *grpc_api.PofpCommentResp, err error) {
	// TODO
	return
}
