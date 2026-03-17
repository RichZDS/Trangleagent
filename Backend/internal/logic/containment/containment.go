package containment

import (
	v1 "TrangleAgent/api/containment/v1"
	"TrangleAgent/internal/dao"
	"TrangleAgent/internal/model"
	"TrangleAgent/internal/model/response"
	"TrangleAgent/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

// 收容库相关方法
type sContainment struct{}

func New() *sContainment {
	return &sContainment{}
}

func init() {
	service.RegisterContainment(New())
}

// ContainmentRepoList 收容库列表查询（带分页与条件）
func (s *sContainment) ContainmentRepoList(ctx context.Context, req *v1.ContainmentRepoListReq) (res *v1.ContainmentRepoListRes, err error) {
	m := dao.ContainmentRepo.Ctx(ctx)

	if req.TerminalId != 0 {
		m = m.Where(dao.ContainmentRepo.Columns().TerminalId, req.TerminalId)
	}
	if req.AgentName != "" {
		m = m.Where(dao.ContainmentRepo.Columns().AgentName, req.AgentName)
	}
	if req.AnomalyName != "" {
		m = m.Where(dao.ContainmentRepo.Columns().AnomalyName, req.AnomalyName)
	}
	if req.RepoName != "" {
		m = m.Where(dao.ContainmentRepo.Columns().RepoName, req.RepoName)
	}

	// 基本分页参数兜底
	page := req.Page
	if page <= 0 {
		page = 1
	}
	size := req.PageSize
	if size <= 0 {
		size = 10
	}

	res = &v1.ContainmentRepoListRes{
		PageResult: response.PageResult{
			Page:     page,
			PageSize: size,
		},
	}

	// 查询总数
	if res.Total, err = m.Count(); err != nil {
		return nil, gerror.Wrap(err, "获取收容库总数失败")
	}

	// 分页查询数据
	var list []*model.ContainmentRepo
	if err = m.Page(page, size).Scan(&list); err != nil {
		return nil, gerror.Wrap(err, "查询收容库列表失败")
	}
	res.List = list
	return res, nil
}

// ContainmentRepoView 查看单条收容库详情
func (s *sContainment) ContainmentRepoView(ctx context.Context, req *v1.ContainmentRepoViewReq) (res *v1.ContainmentRepoViewRes, err error) {
	res = &v1.ContainmentRepoViewRes{}
	if err = dao.ContainmentRepo.Ctx(ctx).
		Where(dao.ContainmentRepo.Columns().Id, req.Id).
		Scan(&res.ContainmentRepoInfo); err != nil {
		return nil, gerror.Wrap(err, "查询收容库详情失败")
	}
	return res, nil
}

// ContainmentRepoUpdate 创建或更新收容库记录
func (s *sContainment) ContainmentRepoUpdate(ctx context.Context, req *v1.ContainmentRepoUpdateReq) (res *v1.ContainmentRepoUpdateRes, err error) {
	res = &v1.ContainmentRepoUpdateRes{}

	data := &model.ContainmentRepo{
		Id:          req.Id,
		TerminalId:  req.TerminalId,
		Abnormal:    0, // Abnormal 可根据业务补充计算或从 req 取值
		AnomalyName: req.AnomalyName,
		AgentName:   req.AgentName,
		RepoName:    req.RepoName,
	}

	if req.Id > 0 {
		// 更新
		if _, err = dao.ContainmentRepo.Ctx(ctx).
			Data(data).
			FieldsEx(dao.ContainmentRepo.Columns().Id).
			Where(dao.ContainmentRepo.Columns().Id, req.Id).
			Update(); err != nil {
			return nil, gerror.Wrap(err, "更新收容库失败")
		}
		res.Id = req.Id
		return res, nil
	}

	// 新增
	r, err := dao.ContainmentRepo.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return nil, gerror.Wrap(err, "创建收容库失败")
	}
	lastId, err := r.LastInsertId()
	if err != nil {
		return nil, gerror.Wrap(err, "获取收容库ID失败")
	}
	res.Id = uint64(lastId)
	return res, nil
}

// ContainmentRepoDelete 删除收容库记录
func (s *sContainment) ContainmentRepoDelete(ctx context.Context, req *v1.ContainmentRepoDeleteReq) (res *v1.ContainmentRepoDeleteRes, err error) {
	res = &v1.ContainmentRepoDeleteRes{}
	if _, err = dao.ContainmentRepo.Ctx(ctx).
		Where(dao.ContainmentRepo.Columns().Id, req.Id).
		Delete(); err != nil {
		return nil, gerror.Wrap(err, "删除收容库失败")
	}
	return res, nil
}
