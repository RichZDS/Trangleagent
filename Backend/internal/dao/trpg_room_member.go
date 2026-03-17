// =================================================================================
// 三角机构 TRPG：房间成员 DAO 入口
// =================================================================================

package dao

import (
	"TrangleAgent/internal/dao/internal"
)

type trpgRoomMemberDao struct {
	*internal.TrpgRoomMemberDao
}

var (
	TrpgRoomMember = trpgRoomMemberDao{internal.NewTrpgRoomMemberDao()}
)
