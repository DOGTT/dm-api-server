package service

import (
	"context"
	"fmt"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/fds"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
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

func (s *Service) WeChatLogin(ctx context.Context, req *api.LoginWeChatReq) (res *api.LoginWeChatRes, err error) {
	log := otelzap.Ctx(ctx)
	log.Debug("wx login get req", zap.Any("req", req))
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
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res.UserInfo, err = s.convertToUserInfo(ctx, userInfo)
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

func validRegisterRequest(req *api.FastRegisterWeChatReq) error {
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

func (s *Service) WeChatRegisterFast(ctx context.Context, req *api.FastRegisterWeChatReq) (res *api.FastRegisterWeChatRes, err error) {
	log := otelzap.Ctx(ctx)
	// Implement me
	log.Debug("grpc impl get req", zap.Any("req", req))
	if validRegisterRequest(req) != nil {
		err = EM_CommonFail_BadRequest.PutDesc("pet is required")
		return
	}
	res = &api.FastRegisterWeChatRes{}
	wxId, err := s.wxCodeToWxId(ctx, req.GetWxCode())
	if err != nil {
		return
	}
	pet := rds.PetInfo{
		Name:     req.GetRegData().GetName(),
		AvatarId: utils.GenUUID(),
	}
	user := &rds.UserInfo{
		WeChatId: wxId,
		Pets:     []rds.PetInfo{pet},
	}
	// 2. create user
	err = s.data.CreateUserInfo(ctx, user)
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
		err = s.data.PutObject(ctx,
			fds.GetBucketName(api.MediaType_USER_AVA),
			pet.AvatarId, req.GetRegData().GetAvatarData())
		if err != nil {
			err = EM_CommonFail_DBError.PutDesc(err.Error())
			return
		}
	}
	// 3. Query Info
	res.UserInfo, err = s.convertToUserInfo(ctx, user)
	if err != nil {
		return
	}
	token, err := s.kp.GenerateToken(utils.TokenClaims{
		UId: user.Id,
	})
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res.Token = token
	return
}

func (s *Service) convertToUserInfo(ctx context.Context, userInfo *rds.UserInfo) (res *api.UserInfo, err error) {
	PIDs := userInfo.GetPIDs()
	res = &api.UserInfo{
		Id:   utils.Uint64ToStr(userInfo.Id),
		Pets: make([]*api.PetInfo, len(PIDs)),
	}
	for i, pet := range userInfo.Pets {
		res.Pets[i] = &api.PetInfo{
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
			media.GetUrl, err = s.data.GenerateGetPresignedURLByMediaInfo(ctx, media)
			if err != nil {
				err = EM_CommonFail_Internal.PutDesc(err.Error())
				return
			}
			res.Pets[i].Avatar = media
		}
	}
	return
}
