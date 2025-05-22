package service

import (
	"context"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) SystemNotifyQuery(ctx context.Context, req *api.SystemNotifyQueryReq) (res *api.SystemNotifyQueryRes, err error) {
	log.D(ctx, "request in", "req", req)
	return
}
