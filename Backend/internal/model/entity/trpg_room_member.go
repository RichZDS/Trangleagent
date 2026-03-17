// =================================================================================
// 三角机构 TRPG：房间成员表（用户-房间-角色关联）
// 执行 manifest/sql/001_optimize_trpg_role_structure.sql 后生效
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TrpgRoomMember 房间成员：记录用户加入房间时选用的角色卡
type TrpgRoomMember struct {
	Id         uint64      `json:"id"          orm:"id"           description:"主键"`
	RoomId     uint64      `json:"roomId"      orm:"room_id"      description:"房间ID"`
	UserId     uint64      `json:"userId"      orm:"user_id"      description:"用户ID"`
	RoleCardId uint64      `json:"roleCardId"  orm:"role_card_id"  description:"使用的角色卡ID"`
	JoinedAt   *gtime.Time `json:"joinedAt"    orm:"joined_at"    description:"加入时间"`
	LeftAt     *gtime.Time `json:"leftAt"     orm:"left_at"      description:"离开时间"`
	Status     int         `json:"status"      orm:"status"       description:"状态：0已离开 1在线"`
}
