// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TrpgRoom is the golang structure for table trpg_room.
type TrpgRoom struct {
	Id             uint64      `json:"id"             orm:"id"              description:"自增主键"`                   // 自增主键
	RoomId         string      `json:"roomId"         orm:"room_id"         description:"房间全局唯一ID"`               // 房间全局唯一ID
	RoomCode       string      `json:"roomCode"       orm:"room_code"       description:"房间号"`                    // 房间号
	RoomName       string      `json:"roomName"       orm:"room_name"       description:"房间名称"`                   // 房间名称
	HostId         uint64      `json:"hostId"         orm:"host_id"         description:"主持人用户ID"`                // 主持人用户ID
	HostNickname   string      `json:"hostNickname"   orm:"host_nickname"   description:"主持人昵称"`                  // 主持人昵称
	MaxPlayers     uint        `json:"maxPlayers"     orm:"max_players"     description:"最大玩家人数"`                 // 最大玩家人数
	CurrentPlayers uint        `json:"currentPlayers" orm:"current_players" description:"当前玩家人数"`                 // 当前玩家人数
	HasPassword    int         `json:"hasPassword"    orm:"has_password"    description:"是否有密码：0无 1有"`            // 是否有密码：0无 1有
	RoomPassword   string      `json:"roomPassword"   orm:"room_password"   description:"房间密码"`                   // 房间密码
	IsPrivate      int         `json:"isPrivate"      orm:"is_private"      description:"是否私密：0公开 1私密"`           // 是否私密：0公开 1私密
	Status         int         `json:"status"         orm:"status"          description:"状态：0未开始 1进行中 2已结束 3已关闭"` // 状态：0未开始 1进行中 2已结束 3已关闭
	SystemName     string      `json:"systemName"     orm:"system_name"     description:"规则系统"`                   // 规则系统
	ScenarioName   string      `json:"scenarioName"   orm:"scenario_name"   description:"模组/剧本名称"`                // 模组/剧本名称
	Description    string      `json:"description"    orm:"description"     description:"房间简介"`                   // 房间简介
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`                   // 创建时间
	StartedAt      *gtime.Time `json:"startedAt"      orm:"started_at"      description:"开团时间"`                   // 开团时间
	EndedAt        *gtime.Time `json:"endedAt"        orm:"ended_at"        description:"结束时间"`                   // 结束时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`                   // 更新时间
}
