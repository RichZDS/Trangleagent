// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TrpgRoomMemberDao is the data access object for the table trpg_room_member.
type TrpgRoomMemberDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  TrpgRoomMemberColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// TrpgRoomMemberColumns defines and stores column names for the table trpg_room_member.
type TrpgRoomMemberColumns struct {
	Id         string // 主键
	RoomId     string // 房间ID
	UserId     string // 用户ID
	RoleCardId string // 使用的角色卡ID
	JoinedAt   string // 加入时间
	LeftAt     string // 离开时间
	Status     string // 状态：0已离开 1在线
}

// trpgRoomMemberColumns holds the columns for the table trpg_room_member.
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

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TrpgRoomMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TrpgRoomMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TrpgRoomMemberDao) Columns() TrpgRoomMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TrpgRoomMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TrpgRoomMemberDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TrpgRoomMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
