package service

import (
	"context"
	"log"

	base_api "github.com/DOGTT/dm-api-server/api/base"
)

func (s *Service) LocationCommonSearch(ctx context.Context, req *base_api.LocationCommonSearchReq) (res *base_api.LocationCommonSearchRes, err error) {
	res = &base_api.LocationCommonSearchRes{}
	// 调用地图查询
	searchRes, err := s.data.MapSearch(ctx, req.GetInput())
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	log.Println(searchRes)
	// 转换搜索结果
	// res.Locations = make([]*base_api.LocationInfo, len(searchRes))
	// for i, item := range searchRes {
	// 	res.Locations[i] = &base_api.LocationInfo{
	// 		Id:      item.ID,
	// 		Title:   item.Title,
	// 		Address: item.Address,
	// 		LngLat: &base_api.PointCoord{
	// 			Lng: item.Location.Lng,
	// 			Lat: item.Location.Lat,
	// 		},
	// 	}
	// }
	return
}
