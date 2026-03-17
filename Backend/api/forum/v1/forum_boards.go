package v1

import (
	"TrangleAgent/internal/model"
	"TrangleAgent/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

// ForumBoardsCreateReq 创建版块请求
type ForumBoardsCreateReq struct {
	g.Meta `path:"/forum/boards/create" method:"post" tags:"论坛版块" summary:"创建版块"`
	model.ForumBoard
}

// ForumBoardsCreateRes 创建版块响应
type ForumBoardsCreateRes struct {
	Id uint64 `json:"id" description:"版块ID"`
}

// ForumBoardsUpdateReq 更新版块请求
type ForumBoardsUpdateReq struct {
	g.Meta `path:"/forum/boards/update" method:"put" tags:"论坛版块" summary:"更新版块"`
	Id     uint64 `json:"id" description:"版块ID"`
	model.ForumBoard
}

// ForumBoardsUpdateRes 更新版块响应
type ForumBoardsUpdateRes struct {
	Id uint64 `json:"id" description:"版块ID"`
}

// ForumBoardsDeleteReq 删除版块请求
type ForumBoardsDeleteReq struct {
	g.Meta `path:"/forum/boards/delete" method:"delete" tags:"论坛版块" summary:"删除版块"`
	Id     uint64 `json:"id" description:"版块ID"`
}

// ForumBoardsDeleteRes 删除版块响应
type ForumBoardsDeleteRes struct {
}

// ForumBoardsViewReq 查看版块请求
type ForumBoardsViewReq struct {
	g.Meta `path:"/forum/boards/view" method:"get" tags:"论坛版块" summary:"查看版块"`
	Id     uint64 `json:"id" description:"版块ID"`
}

// ForumBoardsViewRes 查看版块响应
type ForumBoardsViewRes struct {
	model.ForumBoardViewParams
}

// ForumBoardsListReq 版块列表请求
type ForumBoardsListReq struct {
	response.PageResult
	g.Meta    `path:"/forum/boards/list" method:"get" tags:"论坛版块" summary:"版块列表"`
	SectionId uint64 `json:"sectionId,omitempty" description:"所属频道ID"`
	Name      string `json:"name,omitempty"      description:"版块名称"`
	Status    string `json:"status,omitempty"    description:"版块状态"`
}

// ForumBoardsListRes 版块列表响应
type ForumBoardsListRes struct {
	response.PageResult
	List []*model.ForumBoardViewParams `json:"list" description:"版块列表"`
}
