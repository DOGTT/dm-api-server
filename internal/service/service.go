package service

import (
	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/data"
	"github.com/DOGTT/dm-api-server/internal/utils"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	wechat_miniapp_conf "github.com/silenceper/wechat/v2/miniprogram/config"
)

type Service struct {
	grpc_api.UnimplementedBaseServiceServer
	conf *conf.ServiceConfig
	data *data.DataEntry

	kp *utils.KeyPair

	wcClient *wechat.Wechat

	miniAppHandle *miniprogram.MiniProgram
}

func New(conf *conf.ServiceConfig, data *data.DataEntry) (*Service, error) {
	var err error
	s := &Service{
		conf: conf,
		data: data,

		wcClient: wechat.NewWechat(),
		// miniAppHandle:
	}
	s.kp, err = utils.LoadKeyPair(conf.KeyPair.PublicKey, conf.KeyPair.PrivateKey)
	if err != nil {
		return nil, err
	}
	memory := cache.NewMemory()
	cfg := &wechat_miniapp_conf.Config{
		AppID:     "xxx",
		AppSecret: "xxx",
		Token:     "xxx",
		//EncodingAESKey: "xxxx",
		Cache: memory,
		// Cache: redisCache,
	}
	s.miniAppHandle = s.wcClient.GetMiniProgram(cfg)

	return s, nil
}
