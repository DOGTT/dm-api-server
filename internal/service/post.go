package service

import (
	"context"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) PostLoad(ctx context.Context, req *api.PostLoadReq) (res *api.PostLoadRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = new(api.PostLoadRes)
	var (
		// tc        = utils.GetClaimFromContext(ctx)
		chanId    = utils.StrToUint64(req.GetChannelId())
		loadLimit = uint32(100)
		filter    = &rds.PostFilter{
			RootId: chanId,
			Limit:  loadLimit,
		}
	)
	if req.GetLastPostId() != "" {
		filter.IdFrom = utils.StrToUint64(req.GetLastPostId())
	} else {
		// 倒序查询最新的数量
		filter.OrderByIdDesc = true
	}
	posts, err := s.data.ListPostInfo(ctx, filter)
	if err != nil {
		log.E(ctx, "list post info error", err)
		err = putDescByDBErr(err)
		return
	}
	res.Posts = make([]*api.PostInfo, len(posts))
	for i := range posts {
		res.Posts[i], err = s.convertToPostInfo(ctx, posts[i])
		if err != nil {
			log.E(ctx, "convert post info error", err)
			return
		}
	}
	return
}

func (s *Service) PostQuery(ctx context.Context, req *api.PostQueryReq) (res *api.PostQueryRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = new(api.PostQueryRes)
	// var (
	// 	tc     =  utils.GetClaimFromContext(ctx)
	// 	chanId = utils.StrToUint64(req.GetChanId())
	// )
	// log.D(ctx, "request in", "req", req)
	// TODO
	return
}

func (s *Service) PostQueryByUser(ctx context.Context, req *api.PostQueryByUserReq) (res *api.PostQueryByUserRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = new(api.PostQueryByUserRes)
	// var (
	// 	tc     =  utils.GetClaimFromContext(ctx)
	// 	chanId = utils.StrToUint64(req.GetChanId())
	// )
	// log.D(ctx, "request in", "req", req)
	// TODO
	return
}
func (s *Service) PostReact(ctx context.Context, req *api.PostReactReq) (res *api.PostReactRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = new(api.PostReactRes)
	// var (
	// 	tc     =  utils.GetClaimFromContext(ctx)
	// 	chanId = utils.StrToUint64(req.GetChanId())
	// )
	// log.D(ctx, "request in", "req", req)
	// TODO
	return
}

func (s *Service) PostCreate(ctx context.Context, req *api.PostCreateReq) (res *api.PostCreateRes, err error) {
	log.D(ctx, "request in", "req", req)
	// valid
	if err = validPostCreateRequest(req); err != nil {
		return
	}
	res = new(api.PostCreateRes)
	var (
		tc      = utils.GetClaimFromContext(ctx)
		postReq = req.GetPost()
		post    = &rds.PostInfo{
			Id:        utils.GenSnowflakeId(),
			UId:       tc.UId,
			ChannelId: utils.StrToUint64(postReq.GetChannelId()),
			ParentId:  utils.StrToUint64(postReq.GetParentId()),
			Content:   postReq.GetContent(),
		}
	)
	if err = s.data.CreatePostInfo(ctx, post); err != nil {
		log.E(ctx, "create post info error", err)
		err = putDescByDBErr(err)
		return
	}
	res.Post, err = s.convertToPostInfo(ctx, post)
	if err != nil {
		log.E(ctx, "convert post info error", err)
		return
	}
	s.asyncUpdateChannelStatsByPost(ctx, post.ChannelId, false)
	return
}

func (s *Service) PostDelete(ctx context.Context, req *api.PostDeleteReq) (res *api.PostDeleteRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = &api.PostDeleteRes{}
	var (
		chanId = utils.StrToUint64(req.GetChannelId())
		postId = utils.StrToUint64(req.GetPostId())
		tc     = utils.GetClaimFromContext(ctx)
	)
	if err = s.validPostPermission(ctx, tc, postId); err != nil {
		return
	}
	if err = s.data.DeletePostInfo(ctx, postId); err != nil {
		log.E(ctx, "delete post error", err)
		err = putDescByDBErr(err)
		return
	}
	s.asyncUpdateChannelStatsByPost(ctx, chanId, true)
	return
}

func (s *Service) PostUpdate(ctx context.Context, req *api.PostUpdateReq) (res *api.PostUpdateRes, err error) {
	log.D(ctx, "request in", "req", req)
	if err = s.validPostUpdateRequest(req); err != nil {
		return
	}
	res = new(api.PostUpdateRes)
	var (
		post   = req.GetPost()
		postId = utils.StrToUint64(post.GetId())
		tc     = utils.GetClaimFromContext(ctx)
	)
	if err = s.validPostPermission(ctx, tc, postId); err != nil {
		return
	}
	postData := &rds.PostInfo{
		Id:      postId,
		Content: post.GetContent(),
	}
	if err = s.data.UpdatePostInfo(ctx, postData); err != nil {
		log.E(ctx, "update post info error", err)
		err = putDescByDBErr(err)
		return
	}
	postData, err = s.data.GetPostInfo(ctx, postId)
	if err != nil {
		log.E(ctx, "get post info error", err)
		err = putDescByDBErr(err)
		return
	}
	res.Post, err = s.convertToPostInfo(ctx, postData)
	if err != nil {
		log.E(ctx, "convert post info error", err)
		return
	}
	return
}

func (s *Service) asyncUpdateChannelStatsByPost(ctx context.Context, channelId uint64, isDecrease bool) {
	go func() {
		var err error
		if isDecrease {
			err = s.data.ChannelStatsDecrease(ctx, channelId, rds.ChannelStatsPost)
		} else {
			err = s.data.ChannelStatsIncrease(ctx, channelId, rds.ChannelStatsPost)
		}
		if err != nil {
			log.E(ctx, "post stats update error", err)
		}
	}()
}

func (s *Service) convertToPostInfo(ctx context.Context, in *rds.PostInfo) (res *api.PostInfo, err error) {
	res = &api.PostInfo{
		Id:        utils.Uint64ToStr(in.Id),
		Uid:       utils.Uint64ToStr(in.UId),
		ChannelId: utils.Uint64ToStr(in.ChannelId),
		ParentId:  utils.Uint64ToStr(in.ParentId),
		Content:   in.Content,
	}
	if !in.CreatedAt.IsZero() {
		res.CreatedAt = in.CreatedAt.UnixMilli()
	}
	if !in.UpdatedAt.IsZero() {
		res.UpdatedAt = in.UpdatedAt.UnixMilli()
	}
	return
}

func (s *Service) validPostPermission(ctx context.Context, tc *utils.TokenClaims, postId uint64) error {
	uid, err := s.data.GetPostCreatorId(ctx, postId)
	if err != nil {
		log.E(ctx, "get post creater id error", err)
		err = putDescByDBErr(err)
		return err
	}
	if uid != tc.UId {
		return EM_CommonFail_Forbidden.PutDesc("user has no permission")
	}
	return nil
}

func (s *Service) validPostUpdateRequest(req *api.PostUpdateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetPost() == nil {
		return EM_CommonFail_BadRequest.PutDesc("post is required")
	}
	post := req.GetPost()
	if post.GetId() == "" {
		return EM_CommonFail_BadRequest.PutDesc("id is required")
	}
	if post.GetContent() == "" {
		return EM_CommonFail_BadRequest.PutDesc("content is required")
	}
	// ..
	return nil
}

func validPostCreateRequest(req *api.PostCreateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetPost() == nil {
		return EM_CommonFail_BadRequest.PutDesc("post is required")
	}
	p := req.GetPost()
	if p.GetChannelId() == "" {
		return EM_CommonFail_BadRequest.PutDesc("channel id is required")
	}
	// ..
	return nil
}
