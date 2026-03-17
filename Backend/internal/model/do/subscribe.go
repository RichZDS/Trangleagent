// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Subscribe is the golang structure of table subscribe for DAO operations like Where/Data.
type Subscribe struct {
	g.Meta    `orm:"table:subscribe, do:true"`
	Id        any         // 自增主键
	UserId    any         // 用户ID（关注者，本人）
	FollowId  any         // 被关注的用户ID
	Status    any         // 状态：1=关注中 0=取消关注
	Remark    any         // 备注名（user_id 对 follow_id 的备注）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
