package Forum

import (
	v1 "TrangleAgent/api/forum/v1"
	"TrangleAgent/internal/dao"
	"TrangleAgent/internal/model"
	"TrangleAgent/internal/model/entity"
	"TrangleAgent/internal/model/response"
	"TrangleAgent/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type sForumBoards struct{}

func NewForumBoards() *sForumBoards {
	return &sForumBoards{}
}

func init() {
	service.RegisterForumBoards(NewForumBoards())
}

func (s *sForumBoards) Create(ctx context.Context, req *v1.ForumBoardsCreateReq) (res *v1.ForumBoardsCreateRes, err error) {
	// 校验参数
	if req.Name == "" {
		return nil, gerror.New("版块名称不能为空")
	}

	// 创建版块
	insert := entity.ForumBoards{
		SectionId:    req.SectionId,
		Name:         req.Name,
		Description:  req.Description,
		CoverImage:   req.CoverImage,
		Status:       req.Status,
		DisplayOrder: req.DisplayOrder,
		CreatedAt:    gtime.Now(),
		UpdatedAt:    gtime.Now(),
	}

	// 如果 Status 为空，默认为 normal
	if insert.Status == "" {
		insert.Status = "normal"
	}

	result, err := dao.ForumBoards.Ctx(ctx).Data(insert).Insert()
	if err != nil {
		return nil, gerror.Wrap(err, "创建版块失败")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.Wrap(err, "获取版块ID失败")
	}

	return &v1.ForumBoardsCreateRes{Id: uint64(id)}, nil
}

func (s *sForumBoards) Update(ctx context.Context, req *v1.ForumBoardsUpdateReq) (res *v1.ForumBoardsUpdateRes, err error) {
	if req.Id == 0 {
		return nil, gerror.New("版块ID不能为空")
	}

	// 只能修改部分字段

	data := entity.ForumBoards{
		Id:           req.Id,
		SectionId:    req.SectionId,
		Name:         req.Name,
		Description:  req.Description,
		CoverImage:   req.CoverImage,
		Status:       req.Status,
		DisplayOrder: req.DisplayOrder,
		UpdatedAt:    gtime.Now(),
	}

	_, err = dao.ForumBoards.Ctx(ctx).Data(data).OmitEmpty().Where(dao.ForumBoards.Columns().Id, req.Id).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新版块失败")
	}

	return &v1.ForumBoardsUpdateRes{Id: req.Id}, nil
}

func (s *sForumBoards) Delete(ctx context.Context, req *v1.ForumBoardsDeleteReq) (res *v1.ForumBoardsDeleteRes, err error) {
	if req.Id == 0 {
		return nil, gerror.New("版块ID不能为空")
	}

	// 物理删除
	_, err = dao.ForumBoards.Ctx(ctx).Where(dao.ForumBoards.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除版块失败")
	}

	return &v1.ForumBoardsDeleteRes{}, nil
}

func (s *sForumBoards) View(ctx context.Context, req *v1.ForumBoardsViewReq) (res *v1.ForumBoardsViewRes, err error) {
	if req.Id == 0 {
		return nil, gerror.New("版块ID不能为空")
	}

	var board model.ForumBoardViewParams
	err = dao.ForumBoards.Ctx(ctx).Where(dao.ForumBoards.Columns().Id, req.Id).Scan(&board)
	if err != nil {
		return nil, gerror.Wrap(err, "查询版块失败")
	}

	return &v1.ForumBoardsViewRes{
		ForumBoardViewParams: board,
	}, nil
}

func (s *sForumBoards) List(ctx context.Context, req *v1.ForumBoardsListReq) (res *v1.ForumBoardsListRes, err error) {
	m := dao.ForumBoards.Ctx(ctx)

	if req.SectionId != 0 {
		m = m.Where(dao.ForumBoards.Columns().SectionId, req.SectionId)
	}
	if req.Name != "" {
		m = m.WhereLike(dao.ForumBoards.Columns().Name, "%"+req.Name+"%")
	}
	if req.Status != "" {
		m = m.Where(dao.ForumBoards.Columns().Status, req.Status)
	}

	// 默认按 DisplayOrder 降序，Id 升序
	m = m.OrderDesc(dao.ForumBoards.Columns().DisplayOrder).OrderAsc(dao.ForumBoards.Columns().Id)

	// 分页
	total, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取版块总数失败")
	}

	var list []*model.ForumBoardViewParams

	page := req.Page
	if page <= 0 {
		page = 1
	}
	size := req.PageSize
	if size <= 0 {
		size = 10
	}

	err = m.Page(page, size).Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, "查询版块列表失败")
	}

	return &v1.ForumBoardsListRes{
		PageResult: response.PageResult{
			Total:    total,
			Page:     page,
			PageSize: size,
		},
		List: list,
	}, nil
}
