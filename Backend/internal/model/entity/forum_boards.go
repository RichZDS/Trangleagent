// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumBoards is the golang structure for table forum_boards.
type ForumBoards struct {
	Id             uint64      `json:"id"             orm:"id"               description:"版块ID"`                     // 版块ID
	SectionId      uint64      `json:"sectionId"      orm:"section_id"       description:"所属频道ID"`                   // 所属频道ID
	Name           string      `json:"name"           orm:"name"             description:"版块名称"`                     // 版块名称
	Description    string      `json:"description"    orm:"description"      description:"版块简介"`                     // 版块简介
	CoverImage     string      `json:"coverImage"     orm:"cover_image"      description:"版块封面图URL"`                 // 版块封面图URL
	Status         string      `json:"status"         orm:"status"           description:"状态：normal/hidden/deleted"` // 状态：normal/hidden/deleted
	DisplayOrder   int         `json:"displayOrder"   orm:"display_order"    description:"排序值"`                      // 排序值
	PostCount      uint        `json:"postCount"      orm:"post_count"       description:"帖子总数"`                     // 帖子总数
	TodayPostCount uint        `json:"todayPostCount" orm:"today_post_count" description:"今日发帖数"`                    // 今日发帖数
	LastPostId     uint64      `json:"lastPostId"     orm:"last_post_id"     description:"最后一篇帖子ID"`                 // 最后一篇帖子ID
	LastPostAt     *gtime.Time `json:"lastPostAt"     orm:"last_post_at"     description:"最后发帖时间"`                   // 最后发帖时间
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`                     // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`                     // 更新时间
}
