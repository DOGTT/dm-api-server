package service

import (
	"context"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) LocationCommonSearch(ctx context.Context, req *api.LocationCommonSearchReq) (res *api.LocationCommonSearchRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = &api.LocationCommonSearchRes{}
	// 调用地图查询
	searchRes, err := s.data.MapSearch(ctx, req.GetInput())
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	log.D(ctx, "search result", "searchRes", searchRes)

	// 转换搜索结果
	// res.Locations = make([]*api.LocationInfo, len(searchRes))
	// for i, item := range searchRes {
	// 	res.Locations[i] = &api.LocationInfo{
	// 		Id:      item.ID,
	// 		Title:   item.Title,
	// 		Address: item.Address,
	// 		LngLat: &api.PointCoord{
	// 			Lng: item.Location.Lng,
	// 			Lat: item.Location.Lat,
	// 		},
	// 	}
	// }
	return
}
