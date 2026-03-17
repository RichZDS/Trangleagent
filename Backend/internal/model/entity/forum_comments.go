// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumComments is the golang structure for table forum_comments.
type ForumComments struct {
	Id            uint64      `json:"id"            orm:"id"               description:"评论ID（主键）"`                                 // 评论ID（主键）
	PostId        uint64      `json:"postId"        orm:"post_id"          description:"所属帖子ID（关联 forum_posts.id，无外键）"`            // 所属帖子ID（关联 forum_posts.id，无外键）
	UserId        uint64      `json:"userId"        orm:"user_id"          description:"评论发布者ID（关联 users.id，无外键）"`                 // 评论发布者ID（关联 users.id，无外键）
	ParentId      uint64      `json:"parentId"      orm:"parent_id"        description:"父评论ID：NULL=一级评论；非NULL=二级评论（指向一级评论ID，无外键）"` // 父评论ID：NULL=一级评论；非NULL=二级评论（指向一级评论ID，无外键）
	ReplyToUserId uint64      `json:"replyToUserId" orm:"reply_to_user_id" description:"回复的用户ID（可选，用于展示“回复@xxx”，无外键）"`             // 回复的用户ID（可选，用于展示“回复@xxx”，无外键）
	Content       string      `json:"content"       orm:"content"          description:"评论内容（支持emoji）"`                            // 评论内容（支持emoji）
	Status        string      `json:"status"        orm:"status"           description:"状态：1=正常 2=软删 3=审核中 4=驳回"`                  // 状态：1=正常 2=软删 3=审核中 4=驳回
	LikeCount     uint        `json:"likeCount"     orm:"like_count"       description:"点赞数（冗余）"`                                  // 点赞数（冗余）
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"       description:"创建时间"`                                     // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"       description:"更新时间"`                                     // 更新时间
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"       description:"软删除时间"`                                    // 软删除时间
}
