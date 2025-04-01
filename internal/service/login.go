package service

import (
	"context"
	"fmt"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/fds"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) WeChatLogin(ctx context.Context, req *api.LoginWeChatReq) (res *api.LoginWeChatRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = &api.LoginWeChatRes{}
	wxId, err := s.wxCodeToWxId(ctx, req.GetWxCode())
	if err != nil {
		return
	}
	// query user info
	userInfo, err := s.data.GetUserInfoByWeChatID(ctx, wxId)
	if rds.IsNotFound(err) {
		err = EM_AuthFail_NotFound.PutDesc(err.Error())
		return
	}
	if err != nil {
		err = putDescByDBErr(err)
		return
	}
	// query UserPets
	userPets, err := s.data.ListUserPetByUId(ctx, userInfo.Id)
	if err != nil {
		err = putDescByDBErr(err)
		return
	}
	res.User.UserPets = make([]*api.UserPetInfo, len(userPets))
	for i, up := range userPets {
		res.User.UserPets[i], err = s.convertToUserPetInfo(ctx, up)
		if err != nil {
			return
		}
	}
	res.User, err = s.convertToUserInfo(ctx, userPets[0].User)
	if err != nil {
		return
	}
	token, err := s.kp.GenerateToken(utils.TokenClaims{
		UId: userInfo.Id,
	})
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res.Token = token
	return
}

func (s *Service) dbCreateUserAndPet(ctx context.Context, up *rds.UserPet) (err error) {
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
	err = s.data.CreateUserInfo(ctx, up.User)
	if rds.IsDuplicateErr(err) {
		err = EM_UserFail_AlreadyExist.PutDesc(err.Error())
		return
	}
	if err != nil {
		return
	}
	up.UId = up.User.Id
	up.Pet.UId = up.User.Id
	err = s.data.CreatePetInfo(ctx, up.Pet)
	if err != nil {
		return
	}
	up.PId = up.Pet.Id
	err = s.data.CreateUserPet(ctx, up)
	if err != nil {
		return
	}
	return err
}
func (s *Service) WeChatRegisterFast(ctx context.Context, req *api.FastRegisterWeChatReq) (res *api.FastRegisterWeChatRes, err error) {
	log.D(ctx, "request in", "req", map[string]any{
		"avatar_len": len(req.GetRegData().GetPetAvatarData()),
		"wx_code":    req.GetWxCode(),
	})
	if err = validRegisterRequest(req); err != nil {
		return
	}
	res = &api.FastRegisterWeChatRes{}
	wxId, err := s.wxCodeToWxId(ctx, req.GetWxCode())
	if err != nil {
		return
	}
	userPet := &rds.UserPet{
		PetTitle: req.GetRegData().GetPetTitle(),
		User: &rds.UserInfo{
			WeChatId: wxId,
		},
		Pet: &rds.PetInfo{
			Name:     req.GetRegData().GetPetName(),
			AvatarId: utils.GenUUID(),
		},
	}
	// 2.save avatar
	if req.GetRegData().GetPetAvatarData() != "" {
		avatarData := utils.Base64ToBytes(req.GetRegData().GetPetAvatarData())
		err = s.data.PutObject(ctx,
			fds.GetBucketName(api.MediaType_USER_AVA),
			userPet.Pet.AvatarId, avatarData)
		if err != nil {
			err = putDescByDBErr(err)
			return
		}
	}
	// 3. create user
	err = s.dbCreateUserAndPet(ctx, userPet)
	if err != nil {
		log.E(ctx, "dbCreateUserAndPet error", err)
		return
	}
	// 3. Query Info
	res.User = &api.UserInfo{
		Id: utils.Uint64ToStr(userPet.User.Id),
	}
	var userPetRes *api.UserPetInfo
	userPetRes, err = s.convertToUserPetInfo(ctx, userPet)
	if err != nil {
		return
	}
	res.User.UserPets = []*api.UserPetInfo{userPetRes}
	token, err := s.kp.GenerateToken(utils.TokenClaims{
		UId: userPet.UId,
	})
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res.Token = token
	return
}

func (s *Service) wxCodeToWxId(ctx context.Context, wxCode string) (wxId string, err error) {
	if s.conf.TestWxCode == "*" {
		return wxCode, nil
	}
	if wxCode != "" &&
		wxCode == s.conf.TestWxCode {
		return wxCode, nil
	}
	resAuth, err := s.miniAppHandle.GetAuth().Code2SessionContext(ctx, wxCode)
	if err != nil {
		err = EM_AuthFail_WX.PutDesc(err.Error())
		return
	}
	if resAuth.ErrCode != 0 {
		err = EM_AuthFail_WX.PutDesc(fmt.Sprintf("auth_code:%d", resAuth.ErrCode))
	}
	wxId = resAuth.UnionID
	return
}

func (s *Service) convertToUserInfo(ctx context.Context, user *rds.UserInfo) (res *api.UserInfo, err error) {
	res = &api.UserInfo{
		Id: utils.Uint64ToStr(user.Id),
	}
	return
}

func (s *Service) convertToPetInfo(ctx context.Context, pet *rds.PetInfo) (res *api.PetInfo, err error) {
	res = &api.PetInfo{
		Id:        utils.Uint64ToStr(pet.Id),
		Name:      pet.Name,
		Gender:    uint32(pet.Gender),
		BirthDate: pet.BirthDate,

		CreatedAt: pet.CreatedAt.UnixMilli(),
		UpdatedAt: pet.UpdatedAt.UnixMilli(),
	}
	if pet.AvatarId != "" {
		media := &api.MediaInfo{
			Uuid: pet.AvatarId,
			Type: api.MediaType_USER_AVA,
		}
		media.GetUrl, err = s.data.GenerateGetPresignedURLByMediaInfo(ctx, media, utils.TokenExpireDuration)
		if err != nil {
			err = EM_CommonFail_Internal.PutDesc(err.Error())
			return
		}
		res.Avatar = media
	}
	return
}

func (s *Service) convertToUserPetInfo(ctx context.Context, up *rds.UserPet) (res *api.UserPetInfo, err error) {
	res = &api.UserPetInfo{
		PetTitle: up.PetTitle,
	}
	res.Pet, err = s.convertToPetInfo(ctx, up.Pet)
	return
}

func validRegisterRequest(req *api.FastRegisterWeChatReq) (err error) {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetWxCode() == "" {
		return EM_CommonFail_BadRequest.PutDesc("wx_code is required")
	}
	regData := req.GetRegData()
	if regData == nil {
		return EM_CommonFail_BadRequest.PutDesc("data is required")
	}
	regData.PetName, err = utils.ValidateUsername(regData.GetPetName())
	if err != nil {
		return EM_CommonFail_BadRequest.PutDesc(err.Error())
	}
	// if regData.GetPetAvatarData() == "" {
	// 	return EM_CommonFail_BadRequest.PutDesc("pet_avatar_data is required")
	// }
	return nil
}
