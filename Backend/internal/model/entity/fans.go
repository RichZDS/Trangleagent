// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Fans is the golang structure for table fans.
type Fans struct {
	Id        uint64      `json:"id"        orm:"id"         description:"自增主键"`                         // 自增主键
	UserId    uint64      `json:"userId"    orm:"user_id"    description:"用户ID（被关注者，本人）"`                // 用户ID（被关注者，本人）
	FanId     uint64      `json:"fanId"     orm:"fan_id"     description:"粉丝用户ID"`                       // 粉丝用户ID
	Status    int         `json:"status"    orm:"status"     description:"状态：1=粉丝 0=取关/无效"`              // 状态：1=粉丝 0=取关/无效
	Remark    string      `json:"remark"    orm:"remark"     description:"备注名（user_id 对 fan_id 的备注/分组）"` // 备注名（user_id 对 fan_id 的备注/分组）
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                         // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                         // 更新时间
}
