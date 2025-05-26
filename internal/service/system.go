package service

import (
	"context"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) SystemNotifyGet(ctx context.Context, req *api.SystemNotifyGetReq) (res *api.SystemNotifyGetRes, err error) {
	log.D(ctx, "request in", "req", req)
	return
}

func (s *Service) PetMarkInfoList(ctx context.Context, req *api.PetMarkInfoListReq) (res *api.PetMarkInfoListRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = &api.PetMarkInfoListRes{}
	data, err := s.data.ListChannelTypeInfo(ctx)
	if err != nil {
		return
	}
	res.PetMarks = make([]*api.PetMarkInfo, len(data))
	for i, v := range data {
		res.PetMarks[i] = &api.PetMarkInfo{
			Id:             int32(v.Id),
			Name:           v.Name,
			CoverageRadius: int32(v.CoverageRadius),
			ThemeColor:     v.ThemeColor,
			CreatedAt:      v.CreatedAt.UnixMilli(),
			UpdatedAt:      v.UpdatedAt.UnixMilli(),
		}
	}
	return
}
