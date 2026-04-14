package v1

import (
	"TrangleAgent/internal/model"
	"TrangleAgent/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

// ForumPostsLikeReq 点赞帖子请求
type ForumPostsLikeReq struct {
	g.Meta `path:"/forum/posts/like" method:"post" tags:"论坛帖子" summary:"点赞帖子"`
	Id     uint64 `json:"id" v:"required" description:"帖子ID"`
}

// ForumPostsLikeRes 点赞帖子响应
type ForumPostsLikeRes struct {
	IsLiked bool `json:"isLiked" description:"当前用户是否已点赞"`
}

// ForumPostsCreateReq 创建帖子请求

type ForumPostsCreateReq struct {
	g.Meta     `path:"/forum/posts/create" method:"post" tags:"论坛帖子" summary:"创建帖子"`
	Id         uint64 `json:"id"            orm:"id"              description:"帖子ID"`
	BoardId    uint64 `json:"boardId"       orm:"board_id"        description:"所属版块ID"`
	UserId     uint64 `json:"userId"        orm:"user_id"         description:"发帖用户ID"`
	Title      string `json:"title"         orm:"title"           description:"帖子标题"`
	Content    string `json:"content"       orm:"content"         description:"帖子正文"`
	CoverImage string `json:"coverImage"    orm:"cover_image"     description:"帖子封面图URL"`
	Status     string `json:"status"        orm:"status"          description:"帖子状态"`
}

// ForumPostsCreateRes 创建帖子响应
type ForumPostsCreateRes struct {
	Id uint64 `json:"id" description:"帖子ID"`
}

// ForumPostsUpdateReq 更新帖子请求
type ForumPostsUpdateReq struct {
	g.Meta `path:"/forum/posts/update" method:"put" tags:"论坛帖子" summary:"更新帖子"`
	Id     uint64 `json:"id" description:"帖子ID"`
	model.ForumPost
}

// ForumPostsUpdateRes 更新帖子响应
type ForumPostsUpdateRes struct {
	Id uint64 `json:"id" description:"帖子ID"`
}

// ForumPostsDeleteReq 删除帖子请求
type ForumPostsDeleteReq struct {
	g.Meta `path:"/forum/posts/delete" method:"delete" tags:"论坛帖子" summary:"删除帖子"`
	Id     uint64 `json:"id" description:"帖子ID"`
}

// ForumPostsDeleteRes 删除帖子响应
type ForumPostsDeleteRes struct {
}

// ForumPostsViewReq 查看帖子请求
type ForumPostsViewReq struct {
	g.Meta `path:"/forum/posts/view" method:"get" tags:"论坛帖子" summary:"查看帖子"`
	Id     uint64 `json:"id" description:"帖子ID"`
}

// ForumPostsViewRes 查看帖子响应
type ForumPostsViewRes struct {
	UserName string `json:"name" description:"发帖用户昵称"`
	model.ForumPostViewParams
}

// ForumPostsListReq 帖子列表请求
type ForumPostsListReq struct {
	response.PageResult
	g.Meta    `path:"/forum/posts/list" method:"get" tags:"论坛帖子" summary:"帖子列表"`
	BoardId   uint64 `json:"boardId,omitempty" description:"所属版块ID"`
	UserId    uint64 `json:"userId,omitempty"  description:"发帖用户ID"`
	Title     string `json:"title,omitempty"   description:"帖子标题"`
	Status    string `json:"status,omitempty"  description:"帖子状态"`
	IsPinned  int    `json:"isPinned,omitempty" description:"是否置顶"`
	IsEssence int    `json:"isEssence,omitempty" description:"是否精华"`
}

// ForumPostsListRes 帖子列表响应
type ForumPostsListRes struct {
	response.PageResult
	List []*model.ForumPostViewParams `json:"list" description:"帖子列表"`
}
