// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FansDao is the data access object for the table fans.
type FansDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  FansColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// FansColumns defines and stores column names for the table fans.
type FansColumns struct {
	Id        string // 自增主键
	UserId    string // 用户ID（被关注者，本人）
	FanId     string // 粉丝用户ID
	Status    string // 状态：1=粉丝 0=取关/无效
	Remark    string // 备注名（user_id 对 fan_id 的备注/分组）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// fansColumns holds the columns for the table fans.
var fansColumns = FansColumns{
	Id:        "id",
	UserId:    "user_id",
	FanId:     "fan_id",
	Status:    "status",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewFansDao creates and returns a new DAO object for table data access.
func NewFansDao(handlers ...gdb.ModelHandler) *FansDao {
	return &FansDao{
		group:    "default",
		table:    "fans",
		columns:  fansColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FansDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FansDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FansDao) Columns() FansColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FansDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FansDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FansDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
