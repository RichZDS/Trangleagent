// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumPosts is the golang structure of table forum_posts for DAO operations like Where/Data.
type ForumPosts struct {
	g.Meta        `orm:"table:forum_posts, do:true"`
	Id            any         // 帖子ID（主键）
	BoardId       any         // 所属版块ID（关联 forum_boards.id，无外键）
	UserId        any         // 发帖用户ID（关联 users.id，无外键）
	Title         any         // 帖子标题
	Content       any         // 帖子正文（支持富文本/emoji）
	CoverImage    any         // 帖子封面图URL
	Status        any         // 状态：normal=正常 deleted=软删 audit=审核中 reject=驳回
	IsPinned      any         // 是否置顶：0否 1是
	IsEssence     any         // 是否精华：0否 1是
	ViewCount     any         // 浏览量（冗余）
	LikeCount     any         // 点赞数（冗余）
	CommentCount  any         // 评论数（冗余）
	CollectCount  any         // 收藏数（冗余，可选）
	LastCommentId any         // 最后一条评论ID（冗余，可选）
	LastCommentAt *gtime.Time // 最后评论时间（冗余，可选）
	CreatedAt     *gtime.Time // 发帖时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 软删除时间
}
