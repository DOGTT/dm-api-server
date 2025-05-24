package service

import (
	"context"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) UserInx(ctx context.Context, req *api.UserInxReq) (res *api.UserInxRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = new(api.UserInxRes)
	tc := utils.GetClaimFromContext(ctx)
	chanId := utils.StrToUint64(req.GetChannelId())
	_, err = s.data.GetChannelCreatorId(ctx, chanId)
	// 查询Channel信息，检查是否存在
	if err != nil {
		log.E(ctx, "channel not exist", err)
		err = putDescByDBErr(err)
		return
	}
	// 状态互动基于用户
	if req.GetInxState() != 0 {
		dbIn := &rds.ChannelIxnState{
			UId:       tc.UId,
			ChannelId: chanId,
			IxnState:  req.GetInxState(),
		}
		if req.GetInxStateAction() == api.ActionType_ACTION_TYPE_UNDO {
			err = s.data.DeleteChannelIxnState(ctx, dbIn)
		} else {
			err = s.data.CreateChannelIxnState(ctx, dbIn)
		}
		if err != nil {
			log.E(ctx, "mod user channel ixn state error", err)
			err = putDescByDBErr(err)
			return
		}
	}
	// 事件互动基于用户和爱宠
	if req.GetInxEvent() != 0 {
		// 加载爱宠id
		var petIds []uint64
		petIds, err = s.data.GetPetIdsFromUserId(ctx, tc.UId)
		if err != nil {
			log.E(ctx, "user load error", err)
			err = putDescByDBErr(err)
			return
		}
		events := make([]*rds.ChannelPetIxnEvent, len(petIds))
		for i, pid := range petIds {
			events[i] = &rds.ChannelPetIxnEvent{
				UId:       tc.UId,
				PId:       pid,
				ChannelId: chanId,
				IxnEvent:  req.GetInxEvent(),
			}
		}
		err = s.data.BatchCreateChannelPetIxnEvent(ctx, events)
		if err != nil {
			log.E(ctx, "create user channel ixn event error", err)
			err = putDescByDBErr(err)
			return
		}
	}
	// 更新统计信息
	s.asyncUpdateChannelStatsByInx(ctx, chanId, req)
	return
}
