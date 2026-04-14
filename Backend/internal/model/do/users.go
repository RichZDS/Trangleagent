// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta        `orm:"table:users, do:true"`
	Id            any         // 用户ID
	Account       any         // 账号
	Password      any         // 密码哈希
	Nickname      any         // 昵称
	Gender        any         // 性别：0未知 1男 2女
	BirthDate     *gtime.Time // 生日
	UserType      any         // 用户类型：user/admin
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	Email         any         // 邮箱
	VipStartAt    *gtime.Time // VIP开始时间
	VipEndAt      *gtime.Time // VIP结束时间
	ActiveRoleId  any         // 当前选中的角色卡ID
	Exp           any         // 经验值（经验条）
	Level         any         // 等级（1+exp/100）
	LastCheckinAt *gtime.Time // 上次签到时间
	ExtraInfo     any         // 扩展配置信息
}
