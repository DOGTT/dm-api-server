package service

import (
	"context"

	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func (s *Service) Auth(token string) error {
	return nil
}

func (s *Service) WeChatLogin(ctx context.Context, req *grpc_api.WeChatLoginReq) (res *grpc_api.WeChatLoginResp, err error) {
	// Implement me

	s.miniAppHandle.GetAuth().Code2SessionContext(ctx, req.GetWxCode())

	log.Ctx(c).Debug("grpc impl get req", zap.Any("req", req))
	return
}

func (s *Service) WeChatRegisterWithLogin(c context.Context, req *grpc_api.RegisterWithLoginReq) (res *grpc_api.RegisterWithLoginResp, err error) {
	// Implement me
	log.Ctx(c).Debug("grpc impl get req", zap.Any("req", req))
	return
}
