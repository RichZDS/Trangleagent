package model

import "github.com/gogf/gf/v2/os/gtime"

type User struct {
	Id           uint64      `json:"id"           orm:"id"             description:"用户ID"`
	Account      string      `json:"account"      orm:"account"        description:"账号"`
	Password     string      `json:"password"     orm:"password"       description:"密码哈希"`
	Nickname     string      `json:"nickname"     orm:"nickname"        description:"昵称"`
	Gender       int         `json:"gender"       orm:"gender"         description:"性别：0未知 1男 2女"`
	BirthDate    *gtime.Time `json:"birthDate"    orm:"birth_date"     description:"生日"`
	UserType     string      `json:"userType"     orm:"user_type"       description:"用户类型"`
	VipStartAt   *gtime.Time `json:"vipStartAt"   orm:"vip_start_at"   description:"VIP开始时间"`
	VipEndAt     *gtime.Time `json:"vipEndAt"     orm:"vip_end_at"      description:"VIP结束时间"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"     description:"更新时间"`
	ExtraInfo    string      `json:"extraInfo"    orm:"extra_info"     description:"扩展配置信息"`
	ActiveRoleId uint64      `json:"activeRoleId" orm:"active_role_id" description:"当前选中的角色卡ID"`
	Exp          uint        `json:"exp"          orm:"exp"            description:"经验值"`
	Level        uint        `json:"level"        orm:"level"          description:"等级"`
}

type UserViewParams struct {
	Id            uint64      `json:"id"           orm:"id"             description:"用户ID"`                       // 用户ID
	Account       string      `json:"account"      orm:"account"        description:"账号"`                         // 账号
	Nickname      string      `json:"nickname"     orm:"nickname"        description:"昵称"`                        // 昵称
	Gender        int         `json:"gender"       orm:"gender"          description:"性别：0未知 1男 2女"`              // 性别：0未知 1男 2女
	BirthDate     *gtime.Time `json:"birthDate"    orm:"birth_date"      description:"生日"`                        // 生日
	UserType      string      `json:"userType"     orm:"user_type"       description:"用户类型：normal普通用户，vip为VIP用户"` // 用户类型：normal普通用户，vip为VIP用户
	VipStartAt    *gtime.Time `json:"vipStartAt"   orm:"vip_start_at"    description:"VIP开始时间"`                   // VIP开始时间
	VipEndAt      *gtime.Time `json:"vipEndAt"     orm:"vip_end_at"      description:"VIP结束时间"`                   // VIP结束时间
	CreatedAt     *gtime.Time `json:"createdAt"    orm:"created_at"      description:"创建时间"`                      // 创建时间
	ExtraInfo     string      `json:"extraInfo"    orm:"extra_info"      description:"扩展配置信息"`
	ActiveRoleId  uint64      `json:"activeRoleId"  orm:"active_role_id"   description:"当前选中的角色卡ID"`
	Exp           uint        `json:"exp"           orm:"exp"             description:"经验值"`
	Level         uint        `json:"level"         orm:"level"            description:"等级"`
	LastCheckinAt *gtime.Time `json:"lastCheckinAt"  orm:"last_checkin_at"  description:"上次签到时间"`
	// ARC 来自当前选中角色卡，非 users 表字段
	ArcAbnormal string `json:"arcAbnormal,omitempty" description:"ARC：异常"`
	ArcReality  string `json:"arcReality,omitempty"  description:"ARC：现实"`
	ArcPosition string `json:"arcPosition,omitempty" description:"ARC：职位"`
}
