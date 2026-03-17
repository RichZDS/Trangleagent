// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumBoards is the golang structure of table forum_boards for DAO operations like Where/Data.
type ForumBoards struct {
	g.Meta         `orm:"table:forum_boards, do:true"`
	Id             any         // 版块ID（主键）
	SectionId      any         // 所属频道ID（关联 forum_sections.id，无外键；可为空=不分区）
	Name           any         // 版块名称
	Description    any         // 版块简介
	CoverImage     any         // 版块封面图URL
	Status         any         // 状态：normal=正常 hidden=隐藏 deleted=删除
	DisplayOrder   any         // 排序值（越大越靠前）
	PostCount      any         // 帖子总数（冗余）
	TodayPostCount any         // 今日发帖数（冗余，可选）
	LastPostId     any         // 最后一篇帖子ID（冗余，可选）
	LastPostAt     *gtime.Time // 最后发帖时间（冗余，可选）
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
