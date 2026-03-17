// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package forum

import (
	"context"

	"TrangleAgent/api/forum/v1"
)

type IForumV1 interface {
	ForumBoardsCreate(ctx context.Context, req *v1.ForumBoardsCreateReq) (res *v1.ForumBoardsCreateRes, err error)
	ForumBoardsUpdate(ctx context.Context, req *v1.ForumBoardsUpdateReq) (res *v1.ForumBoardsUpdateRes, err error)
	ForumBoardsDelete(ctx context.Context, req *v1.ForumBoardsDeleteReq) (res *v1.ForumBoardsDeleteRes, err error)
	ForumBoardsView(ctx context.Context, req *v1.ForumBoardsViewReq) (res *v1.ForumBoardsViewRes, err error)
	ForumBoardsList(ctx context.Context, req *v1.ForumBoardsListReq) (res *v1.ForumBoardsListRes, err error)
	ForumCommentsLike(ctx context.Context, req *v1.ForumCommentsLikeReq) (res *v1.ForumCommentsLikeRes, err error)
	ForumCommentsCreate(ctx context.Context, req *v1.ForumCommentsCreateReq) (res *v1.ForumCommentsCreateRes, err error)
	ForumCommentsUpdate(ctx context.Context, req *v1.ForumCommentsUpdateReq) (res *v1.ForumCommentsUpdateRes, err error)
	ForumCommentsDelete(ctx context.Context, req *v1.ForumCommentsDeleteReq) (res *v1.ForumCommentsDeleteRes, err error)
	ForumCommentsView(ctx context.Context, req *v1.ForumCommentsViewReq) (res *v1.ForumCommentsViewRes, err error)
	ForumCommentsList(ctx context.Context, req *v1.ForumCommentsListReq) (res *v1.ForumCommentsListRes, err error)
	ForumPostsLike(ctx context.Context, req *v1.ForumPostsLikeReq) (res *v1.ForumPostsLikeRes, err error)
	ForumPostsCreate(ctx context.Context, req *v1.ForumPostsCreateReq) (res *v1.ForumPostsCreateRes, err error)
	ForumPostsUpdate(ctx context.Context, req *v1.ForumPostsUpdateReq) (res *v1.ForumPostsUpdateRes, err error)
	ForumPostsDelete(ctx context.Context, req *v1.ForumPostsDeleteReq) (res *v1.ForumPostsDeleteRes, err error)
	ForumPostsView(ctx context.Context, req *v1.ForumPostsViewReq) (res *v1.ForumPostsViewRes, err error)
	ForumPostsList(ctx context.Context, req *v1.ForumPostsListReq) (res *v1.ForumPostsListRes, err error)
}
