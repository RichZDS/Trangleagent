// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFollows is the golang structure of table user_follows for DAO operations like Where/Data.
type UserFollows struct {
	g.Meta     `orm:"table:user_follows, do:true"`
	Id         any         // 自增主键
	FollowerId any         // 关注者（粉丝）ID
	FollowedId any         // 被关注者ID
	Status     any         // 状态：1=关注中 0=取消关注
	Remark     any         // 备注名
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
