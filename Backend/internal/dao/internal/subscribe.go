// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SubscribeDao is the data access object for the table subscribe.
type SubscribeDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SubscribeColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SubscribeColumns defines and stores column names for the table subscribe.
type SubscribeColumns struct {
	Id        string // 自增主键
	UserId    string // 用户ID（关注者，本人）
	FollowId  string // 被关注的用户ID
	Status    string // 状态：1=关注中 0=取消关注
	Remark    string // 备注名（user_id 对 follow_id 的备注）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// subscribeColumns holds the columns for the table subscribe.
var subscribeColumns = SubscribeColumns{
	Id:        "id",
	UserId:    "user_id",
	FollowId:  "follow_id",
	Status:    "status",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSubscribeDao creates and returns a new DAO object for table data access.
func NewSubscribeDao(handlers ...gdb.ModelHandler) *SubscribeDao {
	return &SubscribeDao{
		group:    "default",
		table:    "subscribe",
		columns:  subscribeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SubscribeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SubscribeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SubscribeDao) Columns() SubscribeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SubscribeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SubscribeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SubscribeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
