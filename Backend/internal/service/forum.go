// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "TrangleAgent/api/forum/v1"
	"context"
)

type (
	IForumBoards interface {
		Create(ctx context.Context, req *v1.ForumBoardsCreateReq) (res *v1.ForumBoardsCreateRes, err error)
		Update(ctx context.Context, req *v1.ForumBoardsUpdateReq) (res *v1.ForumBoardsUpdateRes, err error)
		Delete(ctx context.Context, req *v1.ForumBoardsDeleteReq) (res *v1.ForumBoardsDeleteRes, err error)
		View(ctx context.Context, req *v1.ForumBoardsViewReq) (res *v1.ForumBoardsViewRes, err error)
		List(ctx context.Context, req *v1.ForumBoardsListReq) (res *v1.ForumBoardsListRes, err error)
	}
	IForumComments interface {
		Like(ctx context.Context, req *v1.ForumCommentsLikeReq) (res *v1.ForumCommentsLikeRes, err error)
		Create(ctx context.Context, req *v1.ForumCommentsCreateReq) (res *v1.ForumCommentsCreateRes, err error)
		Update(ctx context.Context, req *v1.ForumCommentsUpdateReq) (res *v1.ForumCommentsUpdateRes, err error)
		Delete(ctx context.Context, req *v1.ForumCommentsDeleteReq) (res *v1.ForumCommentsDeleteRes, err error)
		View(ctx context.Context, req *v1.ForumCommentsViewReq) (res *v1.ForumCommentsViewRes, err error)
		List(ctx context.Context, req *v1.ForumCommentsListReq) (res *v1.ForumCommentsListRes, err error)
	}
	IForumPosts interface {
		Like(ctx context.Context, req *v1.ForumPostsLikeReq) (res *v1.ForumPostsLikeRes, err error)
		Create(ctx context.Context, req *v1.ForumPostsCreateReq) (res *v1.ForumPostsCreateRes, err error)
		Update(ctx context.Context, req *v1.ForumPostsUpdateReq) (res *v1.ForumPostsUpdateRes, err error)
		Delete(ctx context.Context, req *v1.ForumPostsDeleteReq) (res *v1.ForumPostsDeleteRes, err error)
		View(ctx context.Context, req *v1.ForumPostsViewReq) (res *v1.ForumPostsViewRes, err error)
		List(ctx context.Context, req *v1.ForumPostsListReq) (res *v1.ForumPostsListRes, err error)
	}
)

var (
	localForumBoards   IForumBoards
	localForumComments IForumComments
	localForumPosts    IForumPosts
)

func ForumBoards() IForumBoards {
	if localForumBoards == nil {
		panic("implement not found for interface IForumBoards, forgot register?")
	}
	return localForumBoards
}

func RegisterForumBoards(i IForumBoards) {
	localForumBoards = i
}

func ForumComments() IForumComments {
	if localForumComments == nil {
		panic("implement not found for interface IForumComments, forgot register?")
	}
	return localForumComments
}

func RegisterForumComments(i IForumComments) {
	localForumComments = i
}

func ForumPosts() IForumPosts {
	if localForumPosts == nil {
		panic("implement not found for interface IForumPosts, forgot register?")
	}
	return localForumPosts
}

func RegisterForumPosts(i IForumPosts) {
	localForumPosts = i
}
