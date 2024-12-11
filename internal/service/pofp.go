package service

import (
	"context"

	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
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
	return
}

func (s *Service) PofpFullQueryById(ctx context.Context, req *grpc_api.PofpFullQueryByIdReq) (res *grpc_api.PofpFullQueryByIdResp, err error) {
	return
}

func (s *Service) PofpInteraction(ctx context.Context, req *grpc_api.PofpInteractionReq) (res *grpc_api.PofpInteractionResp, err error) {
	return
}

func (s *Service) PofpComment(ctx context.Context, req *grpc_api.PofpCommentReq) (res *grpc_api.PofpCommentResp, err error) {
	return
}
