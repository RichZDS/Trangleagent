package forum

import (
	"context"

	v1 "TrangleAgent/api/forum/v1"
	"TrangleAgent/internal/service"
)



// ForumBoardsCreate 创建版块
func (c *ControllerV1) ForumBoardsCreate(ctx context.Context, req *v1.ForumBoardsCreateReq) (res *v1.ForumBoardsCreateRes, err error) {
	return service.ForumBoards().Create(ctx, req)
}

// ForumBoardsUpdate 更新版块
func (c *ControllerV1) ForumBoardsUpdate(ctx context.Context, req *v1.ForumBoardsUpdateReq) (res *v1.ForumBoardsUpdateRes, err error) {
	return service.ForumBoards().Update(ctx, req)
}

// ForumBoardsDelete 删除版块
func (c *ControllerV1) ForumBoardsDelete(ctx context.Context, req *v1.ForumBoardsDeleteReq) (res *v1.ForumBoardsDeleteRes, err error) {
	return service.ForumBoards().Delete(ctx, req)
}

// ForumBoardsView 查看版块
func (c *ControllerV1) ForumBoardsView(ctx context.Context, req *v1.ForumBoardsViewReq) (res *v1.ForumBoardsViewRes, err error) {
	return service.ForumBoards().View(ctx, req)
}

// ForumBoardsList 获取版块列表
func (c *ControllerV1) ForumBoardsList(ctx context.Context, req *v1.ForumBoardsListReq) (res *v1.ForumBoardsListRes, err error) {
	return service.ForumBoards().List(ctx, req)
}
