// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDepartment is the golang structure of table user_department for DAO operations like Where/Data.
type UserDepartment struct {
	g.Meta    `orm:"table:user_department, do:true"`
	Id        any         //
	UserId    any         // 用户ID（对应 users.id）
	DeptId    any         // 部门ID（对应 department.id）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
