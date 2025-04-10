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
		up    = req.GetUserPet()
		petId = utils.StrToUint64(up.GetPet().GetId())
		tc    = utils.GetClaimFromContext(ctx)
	)
	if err = s.validPetPermission(ctx, tc, petId); err != nil {
		return
	}
	upData, err := s.convertFromUserPetInfo(ctx, up)
	if err != nil {
		log.E(ctx, "convert from user pet info error", err)
		return
	}
	if err = s.dbUpdateUserPet(ctx, upData); err != nil {
		log.E(ctx, "update user pet info error", err)
		return
	}
	return
}

func (s *Service) dbUpdateUserPet(ctx context.Context, up *rds.UserPet) (err error) {
	tx, err := s.data.NewTransaction(ctx)
	if err != nil {
		err = putDescByDBErr(err)
		return
	}
	defer func() {
		if err != nil {
			err = putDescByDBErr(err)
			if rbErr := tx.Rollback(); rbErr != nil {
				log.E(ctx, "rollback error", rbErr)
			}
		}
	}()
	err = s.data.UpdateUserPet(ctx, up)
	if err != nil {
		return
	}
	if up.Pet != nil {
		err = s.data.UpdatePetInfo(ctx, up.Pet)
		if err != nil {
			return
		}
	}
	return
}

func (s *Service) convertFromPetInfo(ctx context.Context, up *api.PetInfo) (res *rds.PetInfo, err error) {
	res = &rds.PetInfo{
		Id:        utils.StrToUint64(up.Id),
		UId:       utils.StrToUint64(up.Uid),
		Status:    uint8(up.Status),
		Specie:    up.Specie,
		Name:      up.Name,
		Intro:     up.Intro,
		Gender:    uint8(up.Gender),
		BirthDate: up.BirthDate,
		AvatarId:  up.Avatar.GetUuid(),
		Size:      up.Size,
		Breed:     up.Breed,
		Weight:    int(up.Weight),
	}
	return
}

func (s *Service) convertFromUserPetInfo(ctx context.Context, up *api.UserPetInfo) (res *rds.UserPet, err error) {
	res = &rds.UserPet{
		PetTitle:  up.PetTitle,
		PetStatus: uint8(up.PetStatus),
	}
	res.Pet, err = s.convertFromPetInfo(ctx, up.Pet)
	return
}

func (s *Service) validPetPermission(ctx context.Context, tc *utils.TokenClaims, petId uint64) error {
	uid, err := s.data.GetPetCreatorId(ctx, petId)
	if err != nil {
		log.E(ctx, "get channel creater id error", err)
		err = putDescByDBErr(err)
		return err
	}
	if uid != tc.UId {
		return EM_CommonFail_Forbidden.PutDesc("user has no permission")
	}
	return nil
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
