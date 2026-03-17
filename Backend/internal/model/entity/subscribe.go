// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Subscribe is the golang structure for table subscribe.
type Subscribe struct {
	Id        uint64      `json:"id"        orm:"id"         description:"自增主键"`                         // 自增主键
	UserId    uint64      `json:"userId"    orm:"user_id"    description:"用户ID（关注者，本人）"`                 // 用户ID（关注者，本人）
	FollowId  uint64      `json:"followId"  orm:"follow_id"  description:"被关注的用户ID"`                     // 被关注的用户ID
	Status    int         `json:"status"    orm:"status"     description:"状态：1=关注中 0=取消关注"`              // 状态：1=关注中 0=取消关注
	Remark    string      `json:"remark"    orm:"remark"     description:"备注名（user_id 对 follow_id 的备注）"` // 备注名（user_id 对 follow_id 的备注）
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                         // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                         // 更新时间
}
