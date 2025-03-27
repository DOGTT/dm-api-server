package service

import (
	"context"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/data/rds"
	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/DOGTT/dm-api-server/internal/utils/log"
)

func (s *Service) ChannelPostLoad(ctx context.Context, req *api.ChannelPostLoadReq) (res *api.ChannelPostLoadRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = new(api.ChannelPostLoadRes)
	var (
		// tc        = utils.GetClaimFromContext(ctx)
		chanId    = utils.StrToUint64(req.GetChanId())
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
		err = EM_CommonFail_DBError.PutDesc(err.Error())
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

func (s *Service) ChannelPostInx(ctx context.Context, req *api.ChannelPostInxReq) (res *api.ChannelPostInxRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = new(api.ChannelPostInxRes)
	// var (
	// 	tc     =  utils.GetClaimFromContext(ctx)
	// 	chanId = utils.StrToUint64(req.GetChanId())
	// )
	// log.D(ctx, "request in", "req", req)
	// TODO
	return
}

func (s *Service) ChannelPostCreate(ctx context.Context, req *api.ChannelPostCreateReq) (res *api.ChannelPostCreateRes, err error) {
	log.D(ctx, "request in", "req", req)
	// valid
	if err = validPostCreateRequest(req); err != nil {
		return
	}
	res = new(api.ChannelPostCreateRes)
	var (
		tc      = utils.GetClaimFromContext(ctx)
		postReq = req.GetPost()
		post    = &rds.PostInfo{
			Id:       utils.GenSnowflakeID(),
			UId:      tc.UId,
			RootId:   utils.StrToUint64(postReq.GetRootId()),
			ParentId: utils.StrToUint64(postReq.GetParentId()),
			Content:  postReq.GetContent(),
		}
	)
	if err = s.data.CreatePostInfo(ctx, post); err != nil {
		log.E(ctx, "create post info error", err)
		err = EM_CommonFail_DBError.PutDesc(err.Error())
		return
	}
	res.Post, err = s.convertToPostInfo(ctx, post)
	if err != nil {
		log.E(ctx, "convert post info error", err)
		return
	}
	s.asyncUpdateChannelStatsByPost(ctx, post.RootId, false)
	return
}

func (s *Service) convertToPostInfo(ctx context.Context, in *rds.PostInfo) (res *api.PostInfo, err error) {
	res = &api.PostInfo{
		Id:       utils.Uint64ToStr(in.Id),
		Uid:      utils.Uint64ToStr(in.UId),
		RootId:   utils.Uint64ToStr(in.RootId),
		ParentId: utils.Uint64ToStr(in.ParentId),
		Content:  in.Content,
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
		err = EM_CommonFail_DBError.PutDesc(err.Error())
		return err
	}
	if uid != tc.UId {
		return EM_CommonFail_Forbidden.PutDesc("user has no permission")
	}
	return nil
}

func (s *Service) ChannelPostDelete(ctx context.Context, req *api.ChannelPostDeleteReq) (res *api.ChannelPostDeleteRes, err error) {
	log.D(ctx, "request in", "req", req)
	res = &api.ChannelPostDeleteRes{}
	var (
		chanId = utils.StrToUint64(req.GetChanId())
		postId = utils.StrToUint64(req.GetPostId())
		tc     = utils.GetClaimFromContext(ctx)
	)
	if err = s.validPostPermission(ctx, tc, postId); err != nil {
		return
	}
	if err = s.data.DeletePostInfo(ctx, postId); err != nil {
		log.E(ctx, "delete post error", err)
		err = EM_CommonFail_DBError.PutDesc(err.Error())
		return
	}
	s.asyncUpdateChannelStatsByPost(ctx, chanId, true)
	return
}

func (s *Service) validChannelPostUpdateRequest(req *api.ChannelPostUpdateReq) error {
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

func (s *Service) ChannelPostUpdate(ctx context.Context, req *api.ChannelPostUpdateReq) (res *api.ChannelPostUpdateRes, err error) {
	log.D(ctx, "request in", "req", req)
	if err = s.validChannelPostUpdateRequest(req); err != nil {
		return
	}
	res = new(api.ChannelPostUpdateRes)
	var (
		post   = req.GetPost()
		postId = utils.StrToUint64(post.GetId())
		tc     = utils.GetClaimFromContext(ctx)
	)
	if err = s.validChannelPermission(ctx, tc, postId); err != nil {
		return
	}
	if err = s.data.UpdatePostInfo(ctx, &rds.PostInfo{
		Id:      postId,
		Content: post.GetContent(),
	}); err != nil {
		log.E(ctx, "update post info error", err)
		err = EM_CommonFail_DBError.PutDesc(err.Error())
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

func validPostCreateRequest(req *api.ChannelPostCreateReq) error {
	if req == nil {
		return EM_CommonFail_BadRequest.PutDesc("req is required")
	}
	if req.GetPost() == nil {
		return EM_CommonFail_BadRequest.PutDesc("post is required")
	}
	p := req.GetPost()
	if p.GetRootId() == "" {
		return EM_CommonFail_BadRequest.PutDesc("root id is required")
	}
	// ..
	return nil
}
