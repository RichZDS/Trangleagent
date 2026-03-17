// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TrpgRoomMember is the golang structure of table trpg_room_member for DAO operations like Where/Data.
type TrpgRoomMember struct {
	g.Meta     `orm:"table:trpg_room_member, do:true"`
	Id         any         // 主键
	RoomId     any         // 房间ID
	UserId     any         // 用户ID
	RoleCardId any         // 使用的角色卡ID
	JoinedAt   *gtime.Time // 加入时间
	LeftAt     *gtime.Time // 离开时间
	Status     any         // 状态：0已离开 1在线
}
