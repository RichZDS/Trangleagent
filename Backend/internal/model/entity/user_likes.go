// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLikes is the golang structure for table user_likes.
type UserLikes struct {
	Id         uint64      `json:"id"         orm:"id"          description:"自增主键"`                  // 自增主键
	UserId     uint64      `json:"userId"     orm:"user_id"     description:"点赞用户ID"`                // 点赞用户ID
	TargetType string      `json:"targetType" orm:"target_type" description:"被点赞对象类型（post/comment）"` // 被点赞对象类型（post/comment）
	TargetId   uint64      `json:"targetId"   orm:"target_id"   description:"被点赞对象ID"`               // 被点赞对象ID
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`                  // 创建时间
}
