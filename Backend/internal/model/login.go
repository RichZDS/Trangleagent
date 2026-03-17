package model

type LoginField struct {
	Account  string `json:"account"  orm:"account"   v:"required#账号不能为空"`
	Password string `json:"password" orm:"password"  v:"required#密码不能为空"`
	Nickname string `json:"nickname" orm:"nickname"  v:"required#昵称不能为空"`
	Email    string `json:"email"    orm:"email"`
	UserType string `json:"userType" orm:"user_type"` // 用户类型：user/admin，用于前端权限判断
}
