// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLikes is the golang structure of table user_likes for DAO operations like Where/Data.
type UserLikes struct {
	g.Meta     `orm:"table:user_likes, do:true"`
	Id         any         // 自增主键
	UserId     any         // 点赞用户ID
	TargetType any         // 被点赞对象类型（post/comment）
	TargetId   any         // 被点赞对象ID
	CreatedAt  *gtime.Time // 创建时间
}
