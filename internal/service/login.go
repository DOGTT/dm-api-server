package service

import (
	"context"
	"errors"
	"fmt"

	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/DOGTT/dm-api-server/internal/data/fds"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (s *Service) WeChatLogin(ctx context.Context, req *grpc_api.WeChatLoginReq) (res *grpc_api.WeChatLoginResp, err error) {
	// Implement me
	log.Ctx(ctx).Debug("wx login get req", zap.Any("req", req))
	res = &grpc_api.WeChatLoginResp{}
	resAuth, err := s.miniAppHandle.GetAuth().Code2SessionContext(ctx, req.GetWxCode())
	if err != nil {
		err = EM_AuthFail_WX.PutDesc(err.Error())
		return
	}
	if resAuth.ErrCode != 0 {
		err = EM_AuthFail_WX.PutDesc(fmt.Sprintf("auth_code:%d", resAuth.ErrCode))
		return
	}
	// query user info
	userInfo, err := s.data.GetUserInfoByWeChatID(ctx, resAuth.UnionID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
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
	token, err := s.kp.GenerateToken(userInfo.Id)
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res.Token = token
	return
}

func validRegisterRequest(req *grpc_api.WeChatRegisterFastReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetPet() == nil {
		return EM_CommonFail_BadRequest.PutDesc("pet is required")
	}
	if req.GetWxCode() == "" {
		return EM_CommonFail_BadRequest.PutDesc("wx_code is required")
	}
	return nil
}

func (s *Service) WeChatRegisterFast(ctx context.Context, req *grpc_api.WeChatRegisterFastReq) (res *grpc_api.WeChatRegisterFastResp, err error) {
	// Implement me
	log.Ctx(ctx).Debug("grpc impl get req", zap.Any("req", req))
	if validRegisterRequest(req) != nil {
		err = EM_CommonFail_BadRequest.PutDesc("pet is required")
		return
	}
	res = &grpc_api.WeChatRegisterFastResp{}
	resAuth, err := s.miniAppHandle.GetAuth().Code2SessionContext(ctx, req.GetWxCode())
	if err != nil {
		err = EM_AuthFail_WX.PutDesc(err.Error())
		return
	}
	if resAuth.ErrCode != 0 {
		err = EM_AuthFail_WX.PutDesc(fmt.Sprintf("auth_code:%d", resAuth.ErrCode))
		return
	}
	user := &rds.UserInfo{
		WeChatId: resAuth.UnionID,
	}
	pet := &rds.PetInfo{
		Name:     req.GetPet().GetName(),
		AvatarId: utils.GenShortenUUID(),
	}
	// save avatar
	if req.GetPet().GetAvatarData() != nil {
		err = s.data.PutObject(ctx, fds.BucketNameAvatar,
			pet.AvatarId, req.GetPet().GetAvatarData())
		if err != nil {
			err = EM_CommonFail_Internal.PutDesc(err.Error())
			return
		}
	}
	// 2. create user
	err = s.data.CreateUserInfoWithPet(ctx, user, pet)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err = EM_AuthFail_NotFound.PutDesc(err.Error())
		return
	}

	return
}

func (s *Service) convertToUserInfo(ctx context.Context, userInfo *rds.UserInfo) (res *grpc_api.UserInfo, err error) {
	res = &grpc_api.UserInfo{
		Id:   uint32(userInfo.Id),
		Pets: make([]*grpc_api.PetInfo, len(userInfo.Pets)),
	}
	for i, pet := range userInfo.Pets {
		res.Pets[i] = &grpc_api.PetInfo{
			Id:        uint32(pet.Id),
			Name:      pet.Name,
			Gender:    uint32(pet.Gender),
			BirthDate: pet.BirthDate,

			CreatedAt: timestamppb.New(pet.CreatedAt),
			UpdatedAt: timestamppb.New(pet.UpdatedAt),
		}
		if pet.AvatarId != "" {
			res.Pets[i].Avatar, err = s.data.GeneratePresignedURL(ctx,
				fds.BucketNameAvatar, pet.AvatarId, utils.TokenExpireDuration)
			if err != nil {
				err = EM_CommonFail_Internal.PutDesc(err.Error())
				return
			}
		}
	}
	return
}
