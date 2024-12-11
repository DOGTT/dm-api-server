package service

import (
	"context"

	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
)

func (s *Service) LocationCommonSearch(ctx context.Context, req *grpc_api.LocationCommonSearchReq) (res *grpc_api.LocationCommonSearchResp, err error) {
	// 调用地图api
	return
}
