// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumBoards is the golang structure for table forum_boards.
type ForumBoards struct {
	Id             uint64      `json:"id"             orm:"id"               description:"版块ID（主键）"`                                 // 版块ID（主键）
	SectionId      uint64      `json:"sectionId"      orm:"section_id"       description:"所属频道ID（关联 forum_sections.id，无外键；可为空=不分区）"` // 所属频道ID（关联 forum_sections.id，无外键；可为空=不分区）
	Name           string      `json:"name"           orm:"name"             description:"版块名称"`                                     // 版块名称
	Description    string      `json:"description"    orm:"description"      description:"版块简介"`                                     // 版块简介
	CoverImage     string      `json:"coverImage"     orm:"cover_image"      description:"版块封面图URL"`                                 // 版块封面图URL
	Status         string      `json:"status"         orm:"status"           description:"状态：normal=正常 hidden=隐藏 deleted=删除"`        // 状态：normal=正常 hidden=隐藏 deleted=删除
	DisplayOrder   int         `json:"displayOrder"   orm:"display_order"    description:"排序值（越大越靠前）"`                               // 排序值（越大越靠前）
	PostCount      uint        `json:"postCount"      orm:"post_count"       description:"帖子总数（冗余）"`                                 // 帖子总数（冗余）
	TodayPostCount uint        `json:"todayPostCount" orm:"today_post_count" description:"今日发帖数（冗余，可选）"`                             // 今日发帖数（冗余，可选）
	LastPostId     uint64      `json:"lastPostId"     orm:"last_post_id"     description:"最后一篇帖子ID（冗余，可选）"`                          // 最后一篇帖子ID（冗余，可选）
	LastPostAt     *gtime.Time `json:"lastPostAt"     orm:"last_post_at"     description:"最后发帖时间（冗余，可选）"`                            // 最后发帖时间（冗余，可选）
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`                                     // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`                                     // 更新时间
}
