// =================================================================================
// 三角机构 TRPG：房间成员 DAO
// 表需先通过 manifest/sql/001_optimize_trpg_role_structure.sql 创建
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TrpgRoomMemberDao is the data access object for the table trpg_room_member.
type TrpgRoomMemberDao struct {
	table    string
	group    string
	columns  TrpgRoomMemberColumns
	handlers []gdb.ModelHandler
}

// TrpgRoomMemberColumns defines and stores column names for the table trpg_room_member.
type TrpgRoomMemberColumns struct {
	Id         string
	RoomId     string
	UserId     string
	RoleCardId string
	JoinedAt   string
	LeftAt     string
	Status     string
}

var trpgRoomMemberColumns = TrpgRoomMemberColumns{
	Id:         "id",
	RoomId:     "room_id",
	UserId:     "user_id",
	RoleCardId: "role_card_id",
	JoinedAt:   "joined_at",
	LeftAt:     "left_at",
	Status:     "status",
}

// NewTrpgRoomMemberDao creates and returns a new DAO object for table data access.
func NewTrpgRoomMemberDao(handlers ...gdb.ModelHandler) *TrpgRoomMemberDao {
	return &TrpgRoomMemberDao{
		group:    "default",
		table:    "trpg_room_member",
		columns:  trpgRoomMemberColumns,
		handlers: handlers,
	}
}

func (dao *TrpgRoomMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *TrpgRoomMemberDao) Table() string {
	return dao.table
}

func (dao *TrpgRoomMemberDao) Columns() TrpgRoomMemberColumns {
	return dao.columns
}

func (dao *TrpgRoomMemberDao) Group() string {
	return dao.group
}

func (dao *TrpgRoomMemberDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table).Safe().Ctx(ctx)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model
}
