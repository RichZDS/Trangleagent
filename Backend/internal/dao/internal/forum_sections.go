// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForumSectionsDao is the data access object for the table forum_sections.
type ForumSectionsDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  ForumSectionsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// ForumSectionsColumns defines and stores column names for the table forum_sections.
type ForumSectionsColumns struct {
	Id           string // 频道ID
	Name         string // 频道名称
	Description  string // 频道简介
	Icon         string // 频道图标URL
	Status       string // 状态：normal/hidden/deleted
	DisplayOrder string // 排序值
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// forumSectionsColumns holds the columns for the table forum_sections.
var forumSectionsColumns = ForumSectionsColumns{
	Id:           "id",
	Name:         "name",
	Description:  "description",
	Icon:         "icon",
	Status:       "status",
	DisplayOrder: "display_order",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewForumSectionsDao creates and returns a new DAO object for table data access.
func NewForumSectionsDao(handlers ...gdb.ModelHandler) *ForumSectionsDao {
	return &ForumSectionsDao{
		group:    "default",
		table:    "forum_sections",
		columns:  forumSectionsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForumSectionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForumSectionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForumSectionsDao) Columns() ForumSectionsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForumSectionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForumSectionsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ForumSectionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
