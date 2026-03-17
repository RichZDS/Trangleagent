// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"TrangleAgent/api/user/v1"
)

type IUserV1 interface {
	RoleCreate(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleCreateRes, err error)
	RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error)
	RoleView(ctx context.Context, req *v1.RoleViewReq) (res *v1.RoleViewRes, err error)
	RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error)
	RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error)
	RolePermissionCheck(ctx context.Context, req *v1.RolePermissionCheckReq) (res *v1.RolePermissionCheckRes, err error)
	TraceList(ctx context.Context, req *v1.TraceListReq) (res *v1.TraceListRes, err error)
	TraceView(ctx context.Context, req *v1.TraceViewReq) (res *v1.TraceViewRes, err error)
	TraceUpdate(ctx context.Context, req *v1.TraceUpdateReq) (res *v1.TraceUpdateRes, err error)
	TraceReduce(ctx context.Context, req *v1.TraceReduceReq) (res *v1.TraceReduceRes, err error)
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	UserView(ctx context.Context, req *v1.UserViewReq) (res *v1.UserViewRes, err error)
	UserUpdate(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error)
	UserDelete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error)
	ExpAdd(ctx context.Context, req *v1.ExpAddReq) (res *v1.ExpAddRes, err error)
	CheckIn(ctx context.Context, req *v1.CheckInReq) (res *v1.CheckInRes, err error)
}
