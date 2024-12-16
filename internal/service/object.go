package service

import (
	"context"

	base_api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/fds"
	"github.com/DOGTT/dm-api-server/internal/utils"
)

func (s *Service) ObjectPutPresignURLBatchGet(ctx context.Context, req *base_api.ObjectPutPresignURLBatchGetReq) (res *base_api.ObjectPutPresignURLBatchGetResp, err error) {

	var (
		urls      = make([]string, 0, req.GetObjectCount())
		objectIDs = make([]string, 0, req.GetObjectCount())
	)
	for i := 0; i < int(req.GetObjectCount()); i++ {
		uuid := utils.GenUUID()
		url, err := s.data.GeneratePutPresignedURL(ctx,
			fds.GetBucketName(req.GetObjectType()), uuid, fds.PreSignDurationDefault)
		if err != nil {
			err = EM_CommonFail_Internal.PutDesc(err.Error())
			return nil, err
		}
		urls = append(urls, url)
		objectIDs = append(objectIDs, uuid)
	}
	res = &base_api.ObjectPutPresignURLBatchGetResp{
		Urls:      urls,
		ObjectIds: objectIDs,
	}
	return
}
