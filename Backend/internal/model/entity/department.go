// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Department is the golang structure for table department.
type Department struct {
	Id            uint64      `json:"id"            orm:"id"             description:"自增主键"`    // 自增主键
	UserId        uint64      `json:"userId"        orm:"user_id"        description:"所属用户ID"`  // 所属用户ID
	BranchName    string      `json:"branchName"    orm:"branch_name"    description:"分部名称"`    // 分部名称
	TerminalCount int         `json:"terminalCount" orm:"terminal_count" description:"分部散逸端数量"` // 分部散逸端数量
	Weather       string      `json:"weather"       orm:"weather"        description:"分部天气/气候"` // 分部天气/气候
	ManagerName   string      `json:"managerName"   orm:"manager_name"   description:"分部经理名称"`  // 分部经理名称
	Location      string      `json:"location"      orm:"location"       description:"地址"`      // 地址
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`    // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`    // 更新时间
}
