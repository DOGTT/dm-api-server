package service

import (
	"context"

	base_api "github.com/DOGTT/dm-api-server/api/base"
)

func (s *Service) LocationCommonSearch(ctx context.Context, req *base_api.LocationCommonSearchReq) (res *base_api.LocationCommonSearchResp, err error) {
	res = &base_api.LocationCommonSearchResp{}
	// 调用地图api
	return
}
