package service

import (
	"context"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) UserPeGet(ctx context.Context, req *api.UserPeGetReq) (res *api.UserPeGetRes, err error) {
	log.D(ctx, "request in", "req", req)
	tc := utils.GetClaimFromContext(ctx)
	userId := utils.StrToUint64(req.UserId)
	if userId == 0 || tc.UId != userId {
		err = EM_CommonFail_Forbidden.PutDesc("user id invalid")
		return
	}
	res = new(api.UserPeGetRes)
	userPets, err := s.data.ListUserPetByUId(ctx, userId)
	if err != nil {
		log.E(ctx, "create channel error", err)
		err = putDescByDBErr(err)
		return
	}
	res.User = &api.UserInfo{
		Id:       req.UserId,
		UserPets: make([]*api.UserPetInfo, len(userPets)),
	}
	for i := range userPets {
		res.User.UserPets[i], err = s.convertToUserPetInfo(ctx, userPets[i])
		if err != nil {
			return
		}
	}
	return
}

func (s *Service) UserPetUpdate(ctx context.Context, req *api.UserPetUpdateReq) (res *api.UserPetUpdateRes, err error) {
	log.D(ctx, "request in", "req", req)
	if err = validUserPetUpdateReq(req); err != nil {
		return
	}
	res = new(api.UserPetUpdateRes)

	var (
		up     = req.GetUserPet()
		chanId = utils.StrToUint64(ch.GetId())
		tc     = utils.GetClaimFromContext(ctx)
	)
	if err = s.validChannelPermission(ctx, tc, chanId); err != nil {
		return
	}
	if err = s.data.UpdateChannelInfo(ctx, &rds.ChannelInfo{
		Id:       chanId,
		Title:    ch.GetTitle(),
		AvatarId: ch.GetAvatar().GetUuid(),
		Intro:    ch.GetIntro(),
	}); err != nil {
		log.E(ctx, "update channel info error", err)
		err = putDescByDBErr(err)
		return
	}
	return
	return
}

func validUserPetUpdateReq(req *api.UserPetUpdateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetUserPet() == nil {
		return EM_CommonFail_BadRequest.PutDesc("channel is required")
	}
	c := req.GetUserPet()
	if c.GetPetTitle() == "" {
		return EM_CommonFail_BadRequest.PutDesc("title is required")
	}
	// ..
	return nil
}
