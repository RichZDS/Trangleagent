// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserFollowsDao is the data access object for the table user_follows.
type UserFollowsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserFollowsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserFollowsColumns defines and stores column names for the table user_follows.
type UserFollowsColumns struct {
	Id         string // 自增主键
	FollowerId string // 关注者（粉丝）ID
	FollowedId string // 被关注者ID
	Status     string // 状态：1=关注中 0=取消关注
	Remark     string // 备注名
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// userFollowsColumns holds the columns for the table user_follows.
var userFollowsColumns = UserFollowsColumns{
	Id:         "id",
	FollowerId: "follower_id",
	FollowedId: "followed_id",
	Status:     "status",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewUserFollowsDao creates and returns a new DAO object for table data access.
func NewUserFollowsDao(handlers ...gdb.ModelHandler) *UserFollowsDao {
	return &UserFollowsDao{
		group:    "default",
		table:    "user_follows",
		columns:  userFollowsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserFollowsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserFollowsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserFollowsDao) Columns() UserFollowsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserFollowsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserFollowsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserFollowsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
