// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFollows is the golang structure for table user_follows.
type UserFollows struct {
	Id         uint64      `json:"id"         orm:"id"          description:"自增主键"`            // 自增主键
	FollowerId uint64      `json:"followerId" orm:"follower_id" description:"关注者（粉丝）ID"`       // 关注者（粉丝）ID
	FollowedId uint64      `json:"followedId" orm:"followed_id" description:"被关注者ID"`          // 被关注者ID
	Status     int         `json:"status"     orm:"status"      description:"状态：1=关注中 0=取消关注"` // 状态：1=关注中 0=取消关注
	Remark     string      `json:"remark"     orm:"remark"      description:"备注名"`             // 备注名
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`            // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`            // 更新时间
}
