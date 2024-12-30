package service

import (
	"context"

	base_api "github.com/DOGTT/dm-api-server/api/base"
)

func (s *Service) LocationCommonSearch(ctx context.Context, req *base_api.LocationCommonSearchReq) (res *base_api.LocationCommonSearchResp, err error) {
	res = &base_api.LocationCommonSearchResp{}
	// 调用地图api
	// 调用腾讯地图API搜索地点
	searchRes, err := s.data.TencentMapSearch(ctx, req.GetInput(), req.GetBound())
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}

	// 转换搜索结果
	res.Locations = make([]*base_api.LocationInfo, len(searchRes))
	for i, item := range searchRes {
		res.Locations[i] = &base_api.LocationInfo{
			Id:      item.ID,
			Title:   item.Title,
			Address: item.Address,
			LngLat: &base_api.PointCoord{
				Lng: item.Location.Lng,
				Lat: item.Location.Lat,
			},
		}
	}
	return
}
