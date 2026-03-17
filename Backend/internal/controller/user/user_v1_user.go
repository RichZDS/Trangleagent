package user

import (
	"context"

	v1 "TrangleAgent/api/user/v1"
	"TrangleAgent/internal/service"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	return service.User().UserList(ctx, req)
}
func (c *ControllerV1) UserView(ctx context.Context, req *v1.UserViewReq) (res *v1.UserViewRes, err error) {
	return service.User().UserView(ctx, req)
}
func (c *ControllerV1) UserUpdate(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	return service.User().UserUpdate(ctx, req)
}
func (c *ControllerV1) UserDelete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	return service.User().UserDelete(ctx, req)
}

// Trace 相关方法
func (c *ControllerV1) TraceList(ctx context.Context, req *v1.TraceListReq) (res *v1.TraceListRes, err error) {
	return service.Trace().TraceList(ctx, req)
}

func (c *ControllerV1) TraceView(ctx context.Context, req *v1.TraceViewReq) (res *v1.TraceViewRes, err error) {
	return service.Trace().TraceView(ctx, req)
}

func (c *ControllerV1) TraceUpdate(ctx context.Context, req *v1.TraceUpdateReq) (res *v1.TraceUpdateRes, err error) {
	return service.Trace().TraceUpdate(ctx, req)
}

func (c *ControllerV1) TraceReduce(ctx context.Context, req *v1.TraceReduceReq) (res *v1.TraceReduceRes, err error) {
	return service.Trace().TraceReduce(ctx, req)
}
