// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumComments is the golang structure of table forum_comments for DAO operations like Where/Data.
type ForumComments struct {
	g.Meta        `orm:"table:forum_comments, do:true"`
	Id            any         // 评论ID（主键）
	PostId        any         // 所属帖子ID（关联 forum_posts.id，无外键）
	UserId        any         // 评论发布者ID（关联 users.id，无外键）
	ParentId      any         // 父评论ID：NULL=一级评论；非NULL=二级评论（指向一级评论ID，无外键）
	ReplyToUserId any         // 回复的用户ID（可选，用于展示“回复@xxx”，无外键）
	Content       any         // 评论内容（支持emoji）
	Status        any         // 状态：1=正常 2=软删 3=审核中 4=驳回
	LikeCount     any         // 点赞数（冗余）
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 软删除时间
}
