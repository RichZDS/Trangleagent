package Forum

import (
	v1 "TrangleAgent/api/forum/v1"
	"TrangleAgent/internal/dao"
	"TrangleAgent/internal/model"
	"TrangleAgent/internal/model/entity"
	"TrangleAgent/internal/model/response"
	"TrangleAgent/internal/service"
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sForumComments struct{}

func NewForumComments() *sForumComments {
	return &sForumComments{}
}

func init() {
	service.RegisterForumComments(NewForumComments())
}

func (s *sForumComments) Like(ctx context.Context, req *v1.ForumCommentsLikeReq) (res *v1.ForumCommentsLikeRes, err error) {
	// 评论点赞数+1
	_, err = dao.ForumComments.Ctx(ctx).Where(dao.ForumComments.Columns().Id, req.Id).Increment(dao.ForumComments.Columns().LikeCount, 1)
	if err != nil {
		return nil, gerror.Wrap(err, "点赞失败")
	}
	return &v1.ForumCommentsLikeRes{}, nil
}

func (s *sForumComments) Create(ctx context.Context, req *v1.ForumCommentsCreateReq) (res *v1.ForumCommentsCreateRes, err error) {
	if req.PostId == 0 {
		return nil, gerror.New("帖子ID不能为空")
	}
	if req.UserId == 0 {
		return nil, gerror.New("用户ID不能为空")
	}
	if len(req.Content) == 0 {
		return nil, gerror.New("评论内容不能为空")
	}
	status := req.Status
	if status == "" {
		status = "normal"
	} else {
		status = strings.ToLower(status)
	}
	insert := entity.ForumComments{
		PostId:        req.PostId,
		UserId:        req.UserId,
		ParentId:      req.ParentId,
		ReplyToUserId: req.ReplyToUserId,
		Content:       req.Content,
		Status:        status,
		LikeCount:     0,
	}
	result, err := dao.ForumComments.Ctx(ctx).Data(insert).Insert()
	if err != nil {
		return nil, gerror.Wrap(err, "创建评论失败")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.Wrap(err, "获取评论ID失败")
	}
	// 帖子评论数 +1
	_, _ = dao.ForumPosts.Ctx(ctx).Where(dao.ForumPosts.Columns().Id, req.PostId).Increment(dao.ForumPosts.Columns().CommentCount, 1)
	return &v1.ForumCommentsCreateRes{Id: uint64(id)}, nil
}

func (s *sForumComments) Update(ctx context.Context, req *v1.ForumCommentsUpdateReq) (res *v1.ForumCommentsUpdateRes, err error) {
	data := map[string]interface{}{}
	if req.Content != "" {
		data[dao.ForumComments.Columns().Content] = req.Content
	}
	if req.Status != "" {
		data[dao.ForumComments.Columns().Status] = req.Status
	}
	if len(data) == 0 {
		return &v1.ForumCommentsUpdateRes{Id: req.Id}, nil
	}
	_, err = dao.ForumComments.Ctx(ctx).Where(dao.ForumComments.Columns().Id, req.Id).Data(data).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新评论失败")
	}
	return &v1.ForumCommentsUpdateRes{Id: req.Id}, nil
}

func (s *sForumComments) Delete(ctx context.Context, req *v1.ForumCommentsDeleteReq) (res *v1.ForumCommentsDeleteRes, err error) {
	var c entity.ForumComments
	err = dao.ForumComments.Ctx(ctx).Where(dao.ForumComments.Columns().Id, req.Id).Scan(&c)
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论失败")
	}
	_, err = dao.ForumComments.Ctx(ctx).Where(dao.ForumComments.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除评论失败")
	}
	// 帖子评论数 -1
	if c.PostId > 0 {
		dao.ForumPosts.Ctx(ctx).Where(dao.ForumPosts.Columns().Id, c.PostId).Decrement(dao.ForumPosts.Columns().CommentCount, 1)
	}
	return &v1.ForumCommentsDeleteRes{}, nil
}

func (s *sForumComments) View(ctx context.Context, req *v1.ForumCommentsViewReq) (res *v1.ForumCommentsViewRes, err error) {
	var c entity.ForumComments
	err = dao.ForumComments.Ctx(ctx).Where(dao.ForumComments.Columns().Id, req.Id).Scan(&c)
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论失败")
	}
	res = &v1.ForumCommentsViewRes{
		ForumCommentViewParams: model.ForumCommentViewParams{
			Id:            c.Id,
			UserId:        c.UserId,
			PostId:        c.PostId,
			ParentId:      c.ParentId,
			ReplyToUserId: c.ReplyToUserId,
			Content:       c.Content,
			Status:        c.Status,
			LikeCount:     c.LikeCount,
			CreatedAt:     c.CreatedAt,
			UpdatedAt:     c.UpdatedAt,
		},
	}
	return res, nil
}

func (s *sForumComments) List(ctx context.Context, req *v1.ForumCommentsListReq) (res *v1.ForumCommentsListRes, err error) {
	m := dao.ForumComments.Ctx(ctx)
	if req.PostId > 0 {
		m = m.Where(dao.ForumComments.Columns().PostId, req.PostId)
	}
	if req.UserId > 0 {
		m = m.Where(dao.ForumComments.Columns().UserId, req.UserId)
	}
	if req.ParentId > 0 {
		m = m.Where(dao.ForumComments.Columns().ParentId, req.ParentId)
	}
	if req.Status != "" {
		m = m.Where(dao.ForumComments.Columns().Status, req.Status)
	}
	total, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论总数失败")
	}
	m = m.Page(req.Page, req.PageSize).OrderDesc(dao.ForumComments.Columns().CreatedAt)
	var list []*entity.ForumComments
	err = m.Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论列表失败")
	}
	items := make([]*model.ForumCommentViewParams, 0, len(list))
	for _, c := range list {
		items = append(items, &model.ForumCommentViewParams{
			Id:            c.Id,
			UserId:        c.UserId,
			PostId:        c.PostId,
			ParentId:      c.ParentId,
			ReplyToUserId: c.ReplyToUserId,
			Content:       c.Content,
			Status:        c.Status,
			LikeCount:     c.LikeCount,
			CreatedAt:     c.CreatedAt,
			UpdatedAt:     c.UpdatedAt,
		})
	}
	res = &v1.ForumCommentsListRes{
		PageResult: response.PageResult{
			Total:    int(total),
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		List: items,
	}
	return res, nil
}
