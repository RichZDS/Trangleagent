// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumPosts is the golang structure for table forum_posts.
type ForumPosts struct {
	Id            uint64      `json:"id"            orm:"id"              description:"帖子ID"`                           // 帖子ID
	BoardId       uint64      `json:"boardId"       orm:"board_id"        description:"所属版块ID"`                         // 所属版块ID
	UserId        uint64      `json:"userId"        orm:"user_id"         description:"发帖用户ID"`                         // 发帖用户ID
	Title         string      `json:"title"         orm:"title"           description:"帖子标题"`                           // 帖子标题
	Content       string      `json:"content"       orm:"content"         description:"帖子正文"`                           // 帖子正文
	CoverImage    string      `json:"coverImage"    orm:"cover_image"     description:"封面图URL"`                         // 封面图URL
	Status        string      `json:"status"        orm:"status"          description:"状态：normal/deleted/audit/reject"` // 状态：normal/deleted/audit/reject
	IsPinned      int         `json:"isPinned"      orm:"is_pinned"       description:"是否置顶：0否 1是"`                     // 是否置顶：0否 1是
	IsEssence     int         `json:"isEssence"     orm:"is_essence"      description:"是否精华：0否 1是"`                     // 是否精华：0否 1是
	ViewCount     uint        `json:"viewCount"     orm:"view_count"      description:"浏览量"`                            // 浏览量
	LikeCount     uint        `json:"likeCount"     orm:"like_count"      description:"点赞数"`                            // 点赞数
	CommentCount  uint        `json:"commentCount"  orm:"comment_count"   description:"评论数"`                            // 评论数
	CollectCount  uint        `json:"collectCount"  orm:"collect_count"   description:"收藏数"`                            // 收藏数
	LastCommentId uint64      `json:"lastCommentId" orm:"last_comment_id" description:"最后评论ID"`                         // 最后评论ID
	LastCommentAt *gtime.Time `json:"lastCommentAt" orm:"last_comment_at" description:"最后评论时间"`                         // 最后评论时间
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"发帖时间"`                           // 发帖时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`                           // 更新时间
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"      description:"软删除时间"`                          // 软删除时间
}
