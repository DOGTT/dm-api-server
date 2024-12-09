package service

import (
	"context"
	"errors"
	"fmt"

	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
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
	res.Data.UserInfo = convertToUserInfo(userInfo)
	token, err := s.kp.GenerateToken(userInfo.ID)
	if err != nil {
		err = EM_CommonFail_Internal.PutDesc(err.Error())
		return
	}
	res.Data.Token = token
	return
}

func (s *Service) WeChatFastRegister(ctx context.Context, req *grpc_api.WeChatFastRegisterReq) (res *grpc_api.WeChatFastRegisterResp, err error) {
	// Implement me
	log.Ctx(ctx).Debug("grpc impl get req", zap.Any("req", req))
	res = &grpc_api.WeChatFastRegisterResp{}
	resAuth, err := s.miniAppHandle.GetAuth().Code2SessionContext(ctx, req.GetWxCode())
	if err != nil {
		err = EM_AuthFail_WX.PutDesc(err.Error())
		return
	}
	if resAuth.ErrCode != 0 {
		err = EM_AuthFail_WX.PutDesc(fmt.Sprintf("auth_code:%d", resAuth.ErrCode))
		return
	}
	// save avatar

	// 2. create user

	user := &rds.UserInfo{
		WeChatID: resAuth.UnionID,
	}
	pet := &rds.PetInfo{
		Name:   req.GetPet().GetName(),
		Avatar: "",
	}
	err = s.data.CreateUserInfoWithPet(ctx, user, pet)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		err = EM_AuthFail_NotFound.PutDesc(err.Error())
		return
	}

	return
}

func convertToUserInfo(userInfo *rds.UserInfo) *grpc_api.UserInfo {
	r := &grpc_api.UserInfo{
		Id:   uint32(userInfo.ID),
		Pets: make([]*grpc_api.PetInfo, len(userInfo.Pets)),
	}
	for i, pet := range userInfo.Pets {
		r.Pets[i] = &grpc_api.PetInfo{
			Id:        uint32(pet.ID),
			Name:      pet.Name,
			Avatar:    pet.Avatar,
			Gender:    uint32(pet.Gender),
			BirthDate: pet.BirthDate,

			CreatedAt: timestamppb.New(pet.CreatedAt),
			UpdatedAt: timestamppb.New(pet.UpdatedAt),
		}
	}
	return r
}
