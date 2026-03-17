// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumComments is the golang structure for table forum_comments.
type ForumComments struct {
	Id            uint64      `json:"id"            orm:"id"               description:"评论ID"`                           // 评论ID
	PostId        uint64      `json:"postId"        orm:"post_id"          description:"所属帖子ID"`                         // 所属帖子ID
	UserId        uint64      `json:"userId"        orm:"user_id"          description:"评论用户ID"`                         // 评论用户ID
	ParentId      uint64      `json:"parentId"      orm:"parent_id"        description:"父评论ID：0=一级评论"`                   // 父评论ID：0=一级评论
	ReplyToUserId uint64      `json:"replyToUserId" orm:"reply_to_user_id" description:"回复的用户ID"`                        // 回复的用户ID
	Content       string      `json:"content"       orm:"content"          description:"评论内容"`                           // 评论内容
	Status        string      `json:"status"        orm:"status"           description:"状态：normal/deleted/audit/reject"` // 状态：normal/deleted/audit/reject
	LikeCount     uint        `json:"likeCount"     orm:"like_count"       description:"点赞数"`                            // 点赞数
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"       description:"创建时间"`                           // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"       description:"更新时间"`                           // 更新时间
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"       description:"软删除时间"`                          // 软删除时间
}
