package v1

import (
	"TrangleAgent/internal/model"
	"TrangleAgent/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

// ForumCommentsLikeReq 点赞评论请求
type ForumCommentsLikeReq struct {
	g.Meta `path:"/forum/comments/like" method:"post" tags:"论坛评论" summary:"点赞评论"`
	Id     uint64 `json:"id" v:"required" description:"评论ID"`
}

// ForumCommentsLikeRes 点赞评论响应
type ForumCommentsLikeRes struct {
	IsLiked bool `json:"isLiked" description:"当前用户是否已点赞"`
}

// ForumCommentsCreateReq 创建评论请求
type ForumCommentsCreateReq struct {
	g.Meta `path:"/forum/comments/create" method:"post" tags:"论坛评论" summary:"创建评论"`
	model.ForumComment
}

// ForumCommentsCreateRes 创建评论响应
type ForumCommentsCreateRes struct {
	Id uint64 `json:"id" description:"评论ID"`
}

// ForumCommentsUpdateReq 更新评论请求
type ForumCommentsUpdateReq struct {
	g.Meta `path:"/forum/comments/update" method:"put" tags:"论坛评论" summary:"更新评论"`
	Id     uint64 `json:"id" description:"评论ID"`
	model.ForumComment
}

// ForumCommentsUpdateRes 更新评论响应
type ForumCommentsUpdateRes struct {
	Id uint64 `json:"id" description:"评论ID"`
}

// ForumCommentsDeleteReq 删除评论请求
type ForumCommentsDeleteReq struct {
	g.Meta `path:"/forum/comments/delete" method:"delete" tags:"论坛评论" summary:"删除评论"`
	Id     uint64 `json:"id" description:"评论ID"`
}

// ForumCommentsDeleteRes 删除评论响应
type ForumCommentsDeleteRes struct {
}

// ForumCommentsViewReq 查看评论请求
type ForumCommentsViewReq struct {
	g.Meta `path:"/forum/comments/view" method:"get" tags:"论坛评论" summary:"查看评论"`
	Id     uint64 `json:"id" description:"评论ID"`
}

// ForumCommentsViewRes 查看评论响应
type ForumCommentsViewRes struct {
	model.ForumCommentViewParams
}

// ForumCommentsListReq 评论列表请求
type ForumCommentsListReq struct {
	response.PageResult
	g.Meta   `path:"/forum/comments/list" method:"get" tags:"论坛评论" summary:"评论列表"`
	UserId   uint64 `json:"userId,omitempty"   description:"评论发布者ID"`
	PostId   uint64 `json:"postId,omitempty"   description:"所属帖子ID"`
	ParentId uint64 `json:"parentId,omitempty" description:"父评论ID"`
	Status   string `json:"status,omitempty"   description:"评论状态"`
}

// ForumCommentsListRes 评论列表响应
type ForumCommentsListRes struct {
	response.PageResult
	List []*model.ForumCommentViewParams `json:"list" description:"评论列表"`
}
