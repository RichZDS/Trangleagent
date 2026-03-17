package department

import (
	"context"

	"TrangleAgent/api/department"
	v1 "TrangleAgent/api/department/v1"
	"TrangleAgent/internal/service"
)

type ControllerV1 struct{}

func NewV1() department.IDepartmentV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) DepartmentList(ctx context.Context, req *v1.DepartmentListReq) (res *v1.DepartmentListRes, err error) {
	return service.Department().DepartmentList(ctx, req)
}

func (c *ControllerV1) DepartmentView(ctx context.Context, req *v1.DepartmentViewReq) (res *v1.DepartmentViewRes, err error) {
	return service.Department().DepartmentView(ctx, req)
}

func (c *ControllerV1) DepartmentCreate(ctx context.Context, req *v1.DepartmentCreateReq) (res *v1.DepartmentCreateRes, err error) {
	return service.Department().DepartmentCreate(ctx, req)
}

func (c *ControllerV1) DepartmentUpdate(ctx context.Context, req *v1.DepartmentUpdateReq) (res *v1.DepartmentUpdateRes, err error) {
	return service.Department().DepartmentUpdate(ctx, req)
}

func (c *ControllerV1) DepartmentDelete(ctx context.Context, req *v1.DepartmentDeleteReq) (res *v1.DepartmentDeleteRes, err error) {
	return service.Department().DepartmentDelete(ctx, req)
}