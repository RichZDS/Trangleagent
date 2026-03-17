// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ContainmentRepo is the golang structure for table containment_repo.
type ContainmentRepo struct {
	Id          uint64 `json:"id"          orm:"id"           description:"主键"`       // 主键
	AnomalyName string `json:"anomalyName" orm:"anomaly_name" description:"异常体名称"`    // 异常体名称
	AgentName   string `json:"agentName"   orm:"agent_name"   description:"特工名称"`     // 特工名称
	RepoName    string `json:"repoName"    orm:"repo_name"    description:"收容库名称/代码"` // 收容库名称/代码
	Department  int    `json:"department"  orm:"department"   description:"部门"`       // 部门
}
