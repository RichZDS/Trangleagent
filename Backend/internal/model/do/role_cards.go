// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleCards is the golang structure of table role_cards for DAO operations like Where/Data.
type RoleCards struct {
	g.Meta         `orm:"table:role_cards, do:true"`
	Id             any         // 角色卡ID
	UserId         any         // 所属用户ID
	DepartmentId   any         // 所属部门ID
	Commendation   any         // 嘉奖次数
	Reprimand      any         // 申戒次数
	BlueTrack      any         // 蓝轨（0-40）
	YellowTrack    any         // 黄轨（0-40）
	RedTrack       any         // 红轨（0-40）
	ArcAbnormal    any         // ARC：异常
	ArcReality     any         // ARC：现实
	ArcPosition    any         // ARC：职位
	AgentName      any         // 特工名字
	CodeName       any         // 代号
	Gender         any         // 性别
	QaMeticulous   any         // QA Meticulousness (0-100)
	QaDeception    any         // QA Deception (0-100)
	QaVigor        any         // QA Vigor (0-100)
	QaEmpathy      any         // QA Empathy (0-100)
	QaInitiative   any         // QA Initiative (0-100)
	QaResilience   any         // QA Resilience (0-100)
	QaPresence     any         // QA Presence (0-100)
	QaProfessional any         // QA Professionalism (0-100)
	QaDiscretion   any         // QA Discretion (0-100)
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
