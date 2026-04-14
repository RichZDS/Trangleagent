package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumBoard 论坛版块
type ForumBoard struct {
	Id           uint64 `json:"id"             orm:"id"             description:"版块ID"`
	SectionId    uint64 `json:"sectionId"      orm:"section_id"     description:"所属频道ID"`
	Name         string `json:"name"           orm:"name"           description:"版块名称"`
	Description  string `json:"description"    orm:"description"    description:"版块简介"`
	CoverImage   string `json:"coverImage"     orm:"cover_image"    description:"版块封面图URL"`
	Status       string `json:"status"         orm:"status"         description:"版块状态"`
	DisplayOrder int    `json:"displayOrder"   orm:"display_order"  description:"排序值"`
}

// ForumBoardViewParams 版块查看参数
type ForumBoardViewParams struct {
	Id             uint64      `json:"id"             description:"版块ID"`
	SectionId      uint64      `json:"sectionId"      description:"所属频道ID"`
	Name           string      `json:"name"           description:"版块名称"`
	Description    string      `json:"description"    description:"版块简介"`
	CoverImage     string      `json:"coverImage"     description:"版块封面图URL"`
	Status         string      `json:"status"         description:"版块状态"`
	DisplayOrder   int         `json:"displayOrder"   description:"排序值"`
	PostCount      uint        `json:"postCount"      description:"帖子总数"`
	TodayPostCount uint        `json:"todayPostCount" description:"今日发帖数"`
	LastPostId     uint64      `json:"lastPostId"     description:"最后一篇帖子ID"`
	LastPostAt     *gtime.Time `json:"lastPostAt"     description:"最后发帖时间"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"     description:"更新时间"`
}

// ForumPost 论坛帖子
type ForumPost struct {
	Id            uint64      `json:"id"            orm:"id"              description:"帖子ID"`
	BoardId       uint64      `json:"boardId"       orm:"board_id"        description:"所属版块ID"`
	UserId        uint64      `json:"userId"        orm:"user_id"         description:"发帖用户ID"`
	Title         string      `json:"title"         orm:"title"           description:"帖子标题"`
	Content       string      `json:"content"       orm:"content"         description:"帖子正文"`
	CoverImage    string      `json:"coverImage"    orm:"cover_image"     description:"帖子封面图URL"`
	Status        string      `json:"status"        orm:"status"          description:"帖子状态"`
	IsPinned      int         `json:"isPinned"      orm:"is_pinned"       description:"是否置顶"`
	IsEssence     int         `json:"isEssence"     orm:"is_essence"      description:"是否精华"`
	ViewCount     uint        `json:"viewCount"     orm:"view_count"      description:"浏览量"`
	LikeCount     uint        `json:"likeCount"     orm:"like_count"      description:"点赞数"`
	CommentCount  uint        `json:"commentCount"  orm:"comment_count"   description:"评论数"`
	CollectCount  uint        `json:"collectCount"  orm:"collect_count"   description:"收藏数"`
	LastCommentId uint64      `json:"lastCommentId" orm:"last_comment_id" description:"最后一条评论ID"`
	LastCommentAt *gtime.Time `json:"lastCommentAt" orm:"last_comment_at" description:"最后评论时间"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"发帖时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`
}

// ForumPostViewParams 帖子查看参数
type ForumPostViewParams struct {
	Id            uint64      `json:"id"            description:"帖子ID"`
	BoardId       uint64      `json:"boardId"       description:"所属版块ID"`
	UserId        uint64      `json:"userId"        description:"发帖用户ID"`
	Name          string      `json:"name"          description:"发帖用户名称"`
	Title         string      `json:"title"         description:"帖子标题"`
	Content       string      `json:"content"       description:"帖子正文"`
	CoverImage    string      `json:"coverImage"    description:"帖子封面图URL"`
	Status        string      `json:"status"        description:"帖子状态"`
	IsPinned      int         `json:"isPinned"      description:"是否置顶"`
	IsEssence     int         `json:"isEssence"     description:"是否精华"`
	ViewCount     uint        `json:"viewCount"     description:"浏览量"`
	LikeCount     uint        `json:"likeCount"     description:"点赞数"`
	CommentCount  uint        `json:"commentCount"  description:"评论数"`
	CollectCount  uint        `json:"collectCount"  description:"收藏数"`
	LastCommentId uint64      `json:"lastCommentId" description:"最后一条评论ID"`
	LastCommentAt *gtime.Time `json:"lastCommentAt" description:"最后评论时间"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"发帖时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
	IsLiked       bool        `json:"isLiked"       description:"当前用户是否已点赞" orm:"-"`
}

// ForumComment 论坛评论
type ForumComment struct {
	Id            uint64      `json:"id"            orm:"id"               description:"评论ID"`
	PostId        uint64      `json:"postId"        orm:"post_id"          description:"所属帖子ID"`
	UserId        uint64      `json:"userId"        orm:"user_id"          description:"评论发布者ID"`
	ParentId      uint64      `json:"parentId"      orm:"parent_id"        description:"父评论ID"`
	ReplyToUserId uint64      `json:"replyToUserId" orm:"reply_to_user_id" description:"回复的用户ID"`
	Content       string      `json:"content"       orm:"content"          description:"评论内容"`
	Status        string      `json:"status"        orm:"status"           description:"评论状态"`
	LikeCount     uint        `json:"likeCount"     orm:"like_count"       description:"点赞数"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"       description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"    orm:"updated_at"       description:"更新时间"`
}

// ForumCommentViewParams 评论查看参数
type ForumCommentViewParams struct {
	Id            uint64      `json:"id"            description:"评论ID"`
	UserId        uint64      `json:"userId"        description:"评论发布者ID"`
	PostId        uint64      `json:"postId"        description:"所属帖子ID"`
	ParentId      uint64      `json:"parentId"      description:"父评论ID"`
	ReplyToUserId uint64      `json:"replyToUserId" description:"回复的用户ID"`
	Content       string      `json:"content"       description:"评论内容"`
	Status        string      `json:"status"        description:"评论状态"`
	LikeCount     uint        `json:"likeCount"     description:"点赞数"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"评论创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"评论更新时间"`
	IsLiked       bool        `json:"isLiked"       description:"当前用户是否已点赞"`
}
