// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TrpgRoom is the golang structure of table trpg_room for DAO operations like Where/Data.
type TrpgRoom struct {
	g.Meta         `orm:"table:trpg_room, do:true"`
	Id             any         // 自增主键
	RoomId         any         // 房间全局唯一ID
	RoomCode       any         // 房间号
	RoomName       any         // 房间名称
	HostId         any         // 主持人用户ID
	HostNickname   any         // 主持人昵称
	MaxPlayers     any         // 最大玩家人数
	CurrentPlayers any         // 当前玩家人数
	HasPassword    any         // 是否有密码：0无 1有
	RoomPassword   any         // 房间密码
	IsPrivate      any         // 是否私密：0公开 1私密
	Status         any         // 状态：0未开始 1进行中 2已结束 3已关闭
	SystemName     any         // 规则系统
	ScenarioName   any         // 模组/剧本名称
	Description    any         // 房间简介
	CreatedAt      *gtime.Time // 创建时间
	StartedAt      *gtime.Time // 开团时间
	EndedAt        *gtime.Time // 结束时间
	UpdatedAt      *gtime.Time // 更新时间
}
