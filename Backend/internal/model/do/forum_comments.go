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
	Id            any         // 评论ID
	PostId        any         // 所属帖子ID
	UserId        any         // 评论用户ID
	ParentId      any         // 父评论ID：0=一级评论
	ReplyToUserId any         // 回复的用户ID
	Content       any         // 评论内容
	Status        any         // 状态：normal/deleted/audit/reject
	LikeCount     any         // 点赞数
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 软删除时间
}
