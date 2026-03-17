// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumSections is the golang structure of table forum_sections for DAO operations like Where/Data.
type ForumSections struct {
	g.Meta       `orm:"table:forum_sections, do:true"`
	Id           any         // 频道ID
	Name         any         // 频道名称
	Description  any         // 频道简介
	Icon         any         // 频道图标URL
	Status       any         // 状态：normal/hidden/deleted
	DisplayOrder any         // 排序值
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
