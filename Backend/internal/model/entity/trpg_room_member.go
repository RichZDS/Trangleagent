// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TrpgRoomMember is the golang structure for table trpg_room_member.
type TrpgRoomMember struct {
	Id         uint64      `json:"id"         orm:"id"           description:"主键"`          // 主键
	RoomId     uint64      `json:"roomId"     orm:"room_id"      description:"房间ID"`        // 房间ID
	UserId     uint64      `json:"userId"     orm:"user_id"      description:"用户ID"`        // 用户ID
	RoleCardId uint64      `json:"roleCardId" orm:"role_card_id" description:"使用的角色卡ID"`    // 使用的角色卡ID
	JoinedAt   *gtime.Time `json:"joinedAt"   orm:"joined_at"    description:"加入时间"`        // 加入时间
	LeftAt     *gtime.Time `json:"leftAt"     orm:"left_at"      description:"离开时间"`        // 离开时间
	Status     int         `json:"status"     orm:"status"       description:"状态：0已离开 1在线"` // 状态：0已离开 1在线
}
