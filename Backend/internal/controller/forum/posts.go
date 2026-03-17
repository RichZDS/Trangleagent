package forum

import (
	"context"

	v1 "TrangleAgent/api/forum/v1"
	"TrangleAgent/internal/service"
)

func (c *ControllerV1) ForumPostsLike(ctx context.Context, req *v1.ForumPostsLikeReq) (res *v1.ForumPostsLikeRes, err error) {
	return service.ForumPosts().Like(ctx, req)
}

func (c *ControllerV1) ForumPostsCreate(ctx context.Context, req *v1.ForumPostsCreateReq) (res *v1.ForumPostsCreateRes, err error) {
	return service.ForumPosts().Create(ctx, req)
}

func (c *ControllerV1) ForumPostsUpdate(ctx context.Context, req *v1.ForumPostsUpdateReq) (res *v1.ForumPostsUpdateRes, err error) {
	return service.ForumPosts().Update(ctx, req)
}

func (c *ControllerV1) ForumPostsDelete(ctx context.Context, req *v1.ForumPostsDeleteReq) (res *v1.ForumPostsDeleteRes, err error) {
	return service.ForumPosts().Delete(ctx, req)
}

func (c *ControllerV1) ForumPostsView(ctx context.Context, req *v1.ForumPostsViewReq) (res *v1.ForumPostsViewRes, err error) {
	return service.ForumPosts().View(ctx, req)
}

func (c *ControllerV1) ForumPostsList(ctx context.Context, req *v1.ForumPostsListReq) (res *v1.ForumPostsListRes, err error) {
	return service.ForumPosts().List(ctx, req)
}
