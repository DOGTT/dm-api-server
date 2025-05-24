package service

import (
	"context"

	base_api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/fds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) MediaPutURLBatchGet(ctx context.Context, req *base_api.MediaPutURLBatchGetReq) (res *base_api.MediaPutURLBatchGetRes, err error) {
	log.D(ctx, "request in", "req", req)
	var (
		n = int(req.GetCount())
	)
	res = &base_api.MediaPutURLBatchGetRes{
		Media: make([]*base_api.MediaInfo, n),
	}
	for i := 0; i < n; i++ {
		res.Media[i] = &base_api.MediaInfo{
			Bucket: req.GetBucket(),
			Uuid:   utils.GenUUID(),
		}
		res.Media[i].PutUrl, err = s.data.GeneratePutPresignedURL(ctx,
			fds.GetBucketName(req.GetBucket()),
			res.Media[i].Uuid,
			fds.PreSignDurationDefault)
		if err != nil {
			err = EM_CommonFail_Internal.PutDesc(err.Error())
			return nil, err
		}
	}
	return
}
