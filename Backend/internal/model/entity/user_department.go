// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDepartment is the golang structure for table user_department.
type UserDepartment struct {
	Id        uint64      `json:"id"        orm:"id"         description:""`                       //
	UserId    uint64      `json:"userId"    orm:"user_id"    description:"用户ID（对应 users.id）"`      // 用户ID（对应 users.id）
	DeptId    uint64      `json:"deptId"    orm:"dept_id"    description:"部门ID（对应 department.id）"` // 部门ID（对应 department.id）
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                   // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                   // 更新时间
}
