// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForumBoardsDao is the data access object for the table forum_boards.
type ForumBoardsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ForumBoardsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ForumBoardsColumns defines and stores column names for the table forum_boards.
type ForumBoardsColumns struct {
	Id             string // 版块ID（主键）
	SectionId      string // 所属频道ID（关联 forum_sections.id，无外键；可为空=不分区）
	Name           string // 版块名称
	Description    string // 版块简介
	CoverImage     string // 版块封面图URL
	Status         string // 状态：normal=正常 hidden=隐藏 deleted=删除
	DisplayOrder   string // 排序值（越大越靠前）
	PostCount      string // 帖子总数（冗余）
	TodayPostCount string // 今日发帖数（冗余，可选）
	LastPostId     string // 最后一篇帖子ID（冗余，可选）
	LastPostAt     string // 最后发帖时间（冗余，可选）
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// forumBoardsColumns holds the columns for the table forum_boards.
var forumBoardsColumns = ForumBoardsColumns{
	Id:             "id",
	SectionId:      "section_id",
	Name:           "name",
	Description:    "description",
	CoverImage:     "cover_image",
	Status:         "status",
	DisplayOrder:   "display_order",
	PostCount:      "post_count",
	TodayPostCount: "today_post_count",
	LastPostId:     "last_post_id",
	LastPostAt:     "last_post_at",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewForumBoardsDao creates and returns a new DAO object for table data access.
func NewForumBoardsDao(handlers ...gdb.ModelHandler) *ForumBoardsDao {
	return &ForumBoardsDao{
		group:    "default",
		table:    "forum_boards",
		columns:  forumBoardsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForumBoardsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForumBoardsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForumBoardsDao) Columns() ForumBoardsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForumBoardsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForumBoardsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ForumBoardsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
