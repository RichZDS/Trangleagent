package department

import (
	v1 "TrangleAgent/api/department/v1"
	"TrangleAgent/internal/dao"
	"TrangleAgent/internal/model"
	"TrangleAgent/internal/service"
	"context"
)

// 部门相关方法
type sDepartment struct{}

func New() *sDepartment {
	return &sDepartment{}
}

func init() {
	service.RegisterDepartment(New())
}

// DepartmentList 获取部门列表
func (s *sDepartment) DepartmentList(ctx context.Context, req *v1.DepartmentListReq) (res *v1.DepartmentListRes, err error) {
	// 创建数据库查询模型
	m := dao.Department.Ctx(ctx)

	// 根据分部名称查询
	if req.BranchName != "" {
		m = m.WhereLike(dao.Department.Columns().BranchName, "%"+req.BranchName+"%")
	}

	// 根据经理名称查询
	if req.ManagerName != "" {
		m = m.WhereLike(dao.Department.Columns().ManagerName, "%"+req.ManagerName+"%")
	}

	// 根据用户ID查询
	if req.UserId != 0 {
		m = m.Where(dao.Department.Columns().UserId, req.UserId)
	}

	// 查询总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页处理
	m = m.Page(req.Page, req.PageSize)

	// 查询列表数据
	var departmentList []*v1.Department
	err = m.Scan(&departmentList)
	if err != nil {
		return nil, err
	}

	res = &v1.DepartmentListRes{
		List: departmentList,
	}

	// 设置分页信息
	req.PageResult.Total = int(total)
	req.PageResult.Page = req.Page
	req.PageResult.PageSize = req.PageSize

	return
}

// DepartmentView 获取部门详情
func (s *sDepartment) DepartmentView(ctx context.Context, req *v1.DepartmentViewReq) (res *v1.DepartmentViewRes, err error) {
	// 创建数据库查询模型
	m := dao.Department.Ctx(ctx)

	// 根据ID查询
	m = m.Where(dao.Department.Columns().Id, req.Id)

	res = &v1.DepartmentViewRes{}
	err = m.Scan(&res.Department)
	if err != nil {
		return nil, err
	}

	return
}

// DepartmentCreate 创建部门
func (s *sDepartment) DepartmentCreate(ctx context.Context, req *v1.DepartmentCreateReq) (res *v1.DepartmentCreateRes, err error) {
	// 准备数据
	departmentData := model.Department{
		UserId:        req.UserId,
		BranchName:    req.BranchName,
		TerminalCount: req.TerminalCount,
		Weather:       req.Weather,
		ManagerName:   req.ManagerName,
		Location:      req.Location,
	}

	// 插入数据
	result, err := dao.Department.Ctx(ctx).Data(departmentData).Insert()
	if err != nil {
		return nil, err
	}

	// 获取插入的ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	res = &v1.DepartmentCreateRes{
		Id: uint64(lastInsertId),
	}

	return
}

// DepartmentUpdate 更新部门
func (s *sDepartment) DepartmentUpdate(ctx context.Context, req *v1.DepartmentUpdateReq) (res *v1.DepartmentUpdateRes, err error) {
	// 更新数据
	_, err = dao.Department.Ctx(ctx).Data(req).Where(dao.Department.Columns().Id, req.Id).Update()
	if err != nil {
		return nil, err
	}

	res = &v1.DepartmentUpdateRes{
		Id: req.Id,
	}

	return
}

// DepartmentDelete 删除部门
func (s *sDepartment) DepartmentDelete(ctx context.Context, req *v1.DepartmentDeleteReq) (res *v1.DepartmentDeleteRes, err error) {
	// 根据ID删除
	_, err = dao.Department.Ctx(ctx).Where(dao.Department.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, err
	}

	res = &v1.DepartmentDeleteRes{}

	return
}
