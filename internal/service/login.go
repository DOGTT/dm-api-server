package service

import (
	"context"
	"errors"
	"fmt"

	base_api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/fds"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *Service) wxCodeToWxId(ctx context.Context, wxCode string) (wxId string, err error) {
	// testID
	if wxCode != "" && wxCode == s.conf.TestWxCode {
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

func (s *Service) WeChatLogin(ctx context.Context, req *base_api.LoginWeChatReq) (res *base_api.LoginWeChatRes, err error) {
	log.Ctx(ctx).Debug("wx login get req", zap.Any("req", req))
	res = &base_api.LoginWeChatRes{}
	wxId, err := s.wxCodeToWxId(ctx, req.GetWxCode())
	if err != nil {
		return
	}
	// query user info
	userInfo, err := s.data.GetUserInfoByWeChatID(ctx, wxId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = EM_AuthFail_NotFound.PutDesc(err.Error())
		return
	}
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	ackPId := uint64(0)
	if len(userInfo.PIds) > 0 {
		ackPId = userInfo.PIds[0]
	}
	res.UserInfo, err = s.convertToUserInfo(ctx, userInfo)
	if err != nil {
		return
	}
	token, err := s.kp.GenerateToken(utils.TokenClaims{
		UID: userInfo.Id,
		PID: ackPId,
	})
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res.Token = token
	return
}

func validRegisterRequest(req *base_api.FastRegisterWeChatReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetWxCode() == "" {
		return EM_CommonFail_BadRequest.PutDesc("wx_code is required")
	}
	if req.GetRegData() == nil {
		return EM_CommonFail_BadRequest.PutDesc("pet is required")
	}
	return nil
}

func (s *Service) WeChatRegisterFast(ctx context.Context, req *base_api.FastRegisterWeChatReq) (res *base_api.FastRegisterWeChatRes, err error) {
	// Implement me
	log.Ctx(ctx).Debug("grpc impl get req", zap.Any("req", req))
	if validRegisterRequest(req) != nil {
		err = EM_CommonFail_BadRequest.PutDesc("pet is required")
		return
	}
	res = &base_api.FastRegisterWeChatRes{}
	wxId, err := s.wxCodeToWxId(ctx, req.GetWxCode())
	if err != nil {
		return
	}
	user := &rds.UserInfo{
		WeChatId: wxId,
	}
	pet := &rds.PetInfo{
		Name:     req.GetRegData().GetName(),
		AvatarId: utils.GenShortenUUID(),
	}
	// 2. create user
	err = s.data.CreateUserInfoWithPet(ctx, user, pet)
	if rds.IsDuplicateErr(err) {
		err = EM_UserFail_AlreadyExist.PutDesc(err.Error())
		return
	}
	if err != nil {
		err = EM_CommonFail_DBError.PutDesc(err.Error())
		return
	}
	// 3.save avatar
	if req.GetRegData().GetAvatarData() != nil {
		err = s.data.PutObject(ctx, fds.BucketNameAvatar,
			getAvatarIdFromFDS(pet.AvatarId, pet.Id), req.GetRegData().GetAvatarData())
		if err != nil {
			err = EM_CommonFail_DBError.PutDesc(err.Error())
			return
		}
	}
	// 3. Query Info
	// user.PIds = []*rds.PetInfo{pet}
	res.UserInfo, err = s.convertToUserInfo(ctx, user)
	if err != nil {
		return
	}
	token, err := s.kp.GenerateToken(utils.TokenClaims{
		UID: user.Id,
		PID: pet.Id,
	})
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res.Token = token
	return
}

func getAvatarIdFromFDS(avatarId string, pid uint64) string {
	return fmt.Sprintf("%d_%s", pid, avatarId)
}

func (s *Service) convertToUserInfo(ctx context.Context, userInfo *rds.UserInfo) (res *base_api.UserInfo, err error) {
	res = &base_api.UserInfo{
		Id:   userInfo.Id,
		Pets: make([]*base_api.PetInfo, len(userInfo.PIds)),
	}
	for i, pet := range userInfo.Pets {
		res.Pets[i] = &base_api.PetInfo{
			Id:        pet.Id,
			Name:      pet.Name,
			Gender:    uint32(pet.Gender),
			BirthDate: pet.BirthDate,

			CreatedAt: pet.CreatedAt.UnixMilli(),
			UpdatedAt: pet.UpdatedAt.UnixMilli(),
		}
		if pet.AvatarId != "" {
			res.Pets[i].Avatar, err = s.data.GenerateGetPresignedURL(ctx,
				fds.BucketNameAvatar, getAvatarIdFromFDS(pet.AvatarId, pet.Id), fds.PreSignDurationDefault)
			if err != nil {
				err = EM_CommonFail_Internal.PutDesc(err.Error())
				return
			}
		}
	}
	return
}
