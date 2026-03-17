// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "TrangleAgent/api/containment/v1"
	"context"
)

type (
	IContainment interface {
		ContainmentRepoList(ctx context.Context, req *v1.ContainmentRepoListReq) (res *v1.ContainmentRepoListRes, err error)
		ContainmentRepoView(ctx context.Context, req *v1.ContainmentRepoViewReq) (res *v1.ContainmentRepoViewRes, err error)
		ContainmentRepoUpdate(ctx context.Context, req *v1.ContainmentRepoUpdateReq) (res *v1.ContainmentRepoUpdateRes, err error)
		ContainmentRepoDelete(ctx context.Context, req *v1.ContainmentRepoDeleteReq) (res *v1.ContainmentRepoDeleteRes, err error)
	}
)

var (
	localContainment IContainment
)

func Containment() IContainment {
	if localContainment == nil {
		panic("implement not found for interface IContainment, forgot register?")
	}
	return localContainment
}

func RegisterContainment(i IContainment) {
	localContainment = i
}
