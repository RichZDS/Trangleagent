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
	Id            any         // 帖子ID
	BoardId       any         // 所属版块ID
	UserId        any         // 发帖用户ID
	Title         any         // 帖子标题
	Content       any         // 帖子正文
	CoverImage    any         // 封面图URL
	Status        any         // 状态：normal/deleted/audit/reject
	IsPinned      any         // 是否置顶：0否 1是
	IsEssence     any         // 是否精华：0否 1是
	ViewCount     any         // 浏览量
	LikeCount     any         // 点赞数
	CommentCount  any         // 评论数
	CollectCount  any         // 收藏数
	LastCommentId any         // 最后评论ID
	LastCommentAt *gtime.Time // 最后评论时间
	CreatedAt     *gtime.Time // 发帖时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 软删除时间
}
