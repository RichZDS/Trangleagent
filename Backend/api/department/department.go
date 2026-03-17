// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package department

import (
	"context"

	"TrangleAgent/api/department/v1"
)

type IDepartmentV1 interface {
	DepartmentList(ctx context.Context, req *v1.DepartmentListReq) (res *v1.DepartmentListRes, err error)
	DepartmentView(ctx context.Context, req *v1.DepartmentViewReq) (res *v1.DepartmentViewRes, err error)
	DepartmentCreate(ctx context.Context, req *v1.DepartmentCreateReq) (res *v1.DepartmentCreateRes, err error)
	DepartmentUpdate(ctx context.Context, req *v1.DepartmentUpdateReq) (res *v1.DepartmentUpdateRes, err error)
	DepartmentDelete(ctx context.Context, req *v1.DepartmentDeleteReq) (res *v1.DepartmentDeleteRes, err error)
}
