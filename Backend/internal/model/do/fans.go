// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Fans is the golang structure of table fans for DAO operations like Where/Data.
type Fans struct {
	g.Meta    `orm:"table:fans, do:true"`
	Id        any         // 自增主键
	UserId    any         // 用户ID（被关注者，本人）
	FanId     any         // 粉丝用户ID
	Status    any         // 状态：1=粉丝 0=取关/无效
	Remark    any         // 备注名（user_id 对 fan_id 的备注/分组）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
