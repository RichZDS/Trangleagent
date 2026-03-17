// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ContainmentRepo is the golang structure of table containment_repo for DAO operations like Where/Data.
type ContainmentRepo struct {
	g.Meta      `orm:"table:containment_repo, do:true"`
	Id          any // 主键
	AnomalyName any // 异常体名称
	AgentName   any // 特工名称
	RepoName    any // 收容库名称/代码
	Department  any // 部门
}
