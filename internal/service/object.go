package service

import (
	"context"

	base_api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/fds"
	"github.com/DOGTT/dm-api-server/internal/utils"
)

func (s *Service) MediaPutPresignURLBatchGet(ctx context.Context, req *base_api.MediaPutPresignURLBatchGetReq) (res *base_api.MediaPutPresignURLBatchGetRes, err error) {

	var (
		n = int(req.GetCount())
	)
	res = &base_api.MediaPutPresignURLBatchGetRes{
		Media: make([]*base_api.MediaInfo, 0, n),
	}
	for i := 0; i < n; i++ {
		res.Media[i] = &base_api.MediaInfo{
			Type: req.GetMediaType(),
			Uuid: utils.GenUUID(),
		}
		res.Media[i].PutUrl, err = s.data.GeneratePutPresignedURL(ctx,
			fds.GetBucketName(req.GetMediaType()), res.Media[i].Uuid, fds.PreSignDurationDefault)
		if err != nil {
			err = EM_CommonFail_Internal.PutDesc(err.Error())
			return nil, err
		}
	}
	return
}
