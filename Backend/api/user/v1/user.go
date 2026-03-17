package v1

import (
	"TrangleAgent/internal/model"
	"TrangleAgent/internal/model/response"

	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	response.PageResult
	g.Meta   `path:"/user/list" method:"get" tags:"用户" summary:"用户列表"`
	Account  string `json:"account"    orm:"account"      description:"账号"`
	Nickname string `json:"nickname"   orm:"nickname"     description:"昵称"`
}

type UserListRes struct {
	response.PageResult
	Users []*model.UserViewParams
}

type UserViewReq struct {
	g.Meta   `path:"/user/view" method:"get" tags:"用户" summary:"用户详情"`
	Account  string `json:"account"    orm:"account"      description:"账号"`
	Nickname string `json:"nickname"   orm:"nickname"     description:"昵称"`
	Id       uint64 `json:"id"         orm:"id"           description:"用户ID"`
}
type UserViewRes struct {
	model.UserViewParams
}
type UserUpdateReq struct {
	g.Meta  `path:"/user/update" method:"put" tags:"用户" summary:"更新用户"`
	Account string `json:"account"    orm:"account"      description:"账号"`
	model.User
}
type UserUpdateRes struct {
	Id uint64 `json:"id"         orm:"id"           description:"用户ID"`
}
type UserDeleteReq struct {
	g.Meta  `path:"/user/delete" method:"delete" tags:"用户" summary:"删除用户"`
	Account string `json:"account"    orm:"account"      description:"账号"`
}
type UserDeleteRes struct {
}

// ExpAddReq 增加经验请求
type ExpAddReq struct {
	g.Meta `path:"/user/exp/add" method:"post" tags:"用户" summary:"增加经验"`
	UserId uint64 `json:"userId" v:"required" description:"用户ID"`
	Amount uint   `json:"amount" v:"required|min:1" description:"增加的经验值"`
}

// ExpAddRes 增加经验响应
type ExpAddRes struct {
	Exp   uint `json:"exp"   description:"当前经验值"`
	Level uint `json:"level" description:"当前等级"`
}

// CheckInReq 签到请求（每日一次，+20 经验）
type CheckInReq struct {
	g.Meta `path:"/user/checkin" method:"post" tags:"用户" summary:"每日签到"`
	UserId uint64 `json:"userId" v:"required" description:"用户ID"`
}

// CheckInRes 签到响应
type CheckInRes struct {
	Exp       uint   `json:"exp"       description:"当前经验值"`
	Level     uint   `json:"level"     description:"当前等级"`
	AddedExp  uint   `json:"addedExp"  description:"本次获得经验"`
	Message   string `json:"message"   description:"提示信息"`
}
