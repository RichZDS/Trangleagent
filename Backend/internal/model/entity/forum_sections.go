// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumSections is the golang structure for table forum_sections.
type ForumSections struct {
	Id           uint64      `json:"id"           orm:"id"            description:"频道/分区ID（主键）"`                       // 频道/分区ID（主键）
	Name         string      `json:"name"         orm:"name"          description:"频道名称"`                              // 频道名称
	Description  string      `json:"description"  orm:"description"   description:"频道简介"`                              // 频道简介
	Icon         string      `json:"icon"         orm:"icon"          description:"频道图标URL"`                           // 频道图标URL
	Status       string      `json:"status"       orm:"status"        description:"状态：normal=正常 hidden=隐藏 deleted=删除"` // 状态：normal=正常 hidden=隐藏 deleted=删除
	DisplayOrder int         `json:"displayOrder" orm:"display_order" description:"排序值（越大越靠前）"`                        // 排序值（越大越靠前）
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`                              // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`                              // 更新时间
}
