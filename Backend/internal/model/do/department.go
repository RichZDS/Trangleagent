// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Department is the golang structure of table department for DAO operations like Where/Data.
type Department struct {
	g.Meta        `orm:"table:department, do:true"`
	Id            any         // 自增主键
	UserId        any         // 所属用户ID（对应 users.id）
	BranchName    any         // 分部名称
	TerminalCount any         // 分部散逸端的数量
	Weather       any         // 分部当前天气/气候描述
	ManagerName   any         // 分部经理名称
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	Location      any         // 地址
}
