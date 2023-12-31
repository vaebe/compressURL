// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortUrlCodeDao is the data access object for table short_url_code.
type ShortUrlCodeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ShortUrlCodeColumns // columns contains all the column names of Table for convenient usage.
}

// ShortUrlCodeColumns defines and stores column names for table short_url_code.
type ShortUrlCodeColumns struct {
	Id        string // 唯一标识，自增长整数
	Code      string // 短链,唯一，不能为空
	Status    string // 是否使用
	CreatedAt string // 创建时间，默认为当前时间戳
}

// shortUrlCodeColumns holds the columns for table short_url_code.
var shortUrlCodeColumns = ShortUrlCodeColumns{
	Id:        "id",
	Code:      "code",
	Status:    "status",
	CreatedAt: "created_at",
}

// NewShortUrlCodeDao creates and returns a new DAO object for table data access.
func NewShortUrlCodeDao() *ShortUrlCodeDao {
	return &ShortUrlCodeDao{
		group:   "default",
		table:   "short_url_code",
		columns: shortUrlCodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortUrlCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortUrlCodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortUrlCodeDao) Columns() ShortUrlCodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortUrlCodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortUrlCodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortUrlCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
