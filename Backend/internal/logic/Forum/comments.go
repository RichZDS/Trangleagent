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

	"github.com/gogf/gf/v2/database/gdb"
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
	userId, err := getUserIdFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	var isLiked bool
	err = dao.UserLikes.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		count, e := dao.UserLikes.Ctx(ctx).
			Where(dao.UserLikes.Columns().UserId, userId).
			Where(dao.UserLikes.Columns().TargetType, "comment").
			Where(dao.UserLikes.Columns().TargetId, req.Id).
			Count()
		if e != nil {
			return e
		}
		if count > 0 {
			_, e = dao.UserLikes.Ctx(ctx).
				Where(dao.UserLikes.Columns().UserId, userId).
				Where(dao.UserLikes.Columns().TargetType, "comment").
				Where(dao.UserLikes.Columns().TargetId, req.Id).
				Delete()
			if e != nil {
				return e
			}
			_, e = dao.Comments.Ctx(ctx).
				Where(dao.Comments.Columns().Id, req.Id).
				Decrement(dao.Comments.Columns().LikeCount, 1)
			isLiked = false
			return e
		}
		_, e = dao.UserLikes.Ctx(ctx).Data(entity.UserLikes{
			UserId:     userId,
			TargetType: "comment",
			TargetId:   req.Id,
		}).Insert()
		if e != nil {
			return e
		}
		_, e = dao.Comments.Ctx(ctx).
			Where(dao.Comments.Columns().Id, req.Id).
			Increment(dao.Comments.Columns().LikeCount, 1)
		isLiked = true
		return e
	})
	if err != nil {
		return nil, gerror.Wrap(err, "点赞操作失败")
	}
	return &v1.ForumCommentsLikeRes{IsLiked: isLiked}, nil
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
	insert := entity.Comments{
		TargetId:      req.PostId,
		TargetType:    "post",
		UserId:        req.UserId,
		ParentId:      req.ParentId,
		Content:       req.Content,
		Status:        1,
		LikeCount:     0,
	}
	result, err := dao.Comments.Ctx(ctx).Data(insert).Insert()
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
		data[dao.Comments.Columns().Content] = req.Content
	}
	if req.Status != "" {
		data[dao.Comments.Columns().Status] = req.Status
	}
	if len(data) == 0 {
		return &v1.ForumCommentsUpdateRes{Id: req.Id}, nil
	}
	_, err = dao.Comments.Ctx(ctx).Where(dao.Comments.Columns().Id, req.Id).Data(data).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新评论失败")
	}
	return &v1.ForumCommentsUpdateRes{Id: req.Id}, nil
}

func (s *sForumComments) Delete(ctx context.Context, req *v1.ForumCommentsDeleteReq) (res *v1.ForumCommentsDeleteRes, err error) {
	var c entity.Comments
	err = dao.Comments.Ctx(ctx).Where(dao.Comments.Columns().Id, req.Id).Scan(&c)
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论失败")
	}
	_, err = dao.Comments.Ctx(ctx).Where(dao.Comments.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除评论失败")
	}
	// 帖子评论数 -1
	if c.TargetType == "post" && c.TargetId > 0 {
		dao.ForumPosts.Ctx(ctx).Where(dao.ForumPosts.Columns().Id, c.TargetId).Decrement(dao.ForumPosts.Columns().CommentCount, 1)
	}
	return &v1.ForumCommentsDeleteRes{}, nil
}

func (s *sForumComments) View(ctx context.Context, req *v1.ForumCommentsViewReq) (res *v1.ForumCommentsViewRes, err error) {
	var c entity.Comments
	err = dao.Comments.Ctx(ctx).Where(dao.Comments.Columns().Id, req.Id).Scan(&c)
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论失败")
	}
	status := "normal"
	if c.Status == 0 {
		status = "hidden"
	}

	var isLiked bool
	if userId, e := getUserIdFromCtx(ctx); e == nil && userId > 0 {
		cnt, _ := dao.UserLikes.Ctx(ctx).
			Where(dao.UserLikes.Columns().UserId, userId).
			Where(dao.UserLikes.Columns().TargetType, "comment").
			Where(dao.UserLikes.Columns().TargetId, req.Id).
			Count()
		isLiked = cnt > 0
	}

	res = &v1.ForumCommentsViewRes{
		ForumCommentViewParams: model.ForumCommentViewParams{
			Id:        c.Id,
			UserId:    c.UserId,
			PostId:    c.TargetId,
			ParentId:  c.ParentId,
			Content:   c.Content,
			Status:    status,
			LikeCount: c.LikeCount,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			IsLiked:   isLiked,
		},
	}
	return res, nil
}

func (s *sForumComments) List(ctx context.Context, req *v1.ForumCommentsListReq) (res *v1.ForumCommentsListRes, err error) {
	m := dao.Comments.Ctx(ctx).Where(dao.Comments.Columns().TargetType, "post")
	if req.PostId > 0 {
		m = m.Where(dao.Comments.Columns().TargetId, req.PostId)
	}
	if req.UserId > 0 {
		m = m.Where(dao.Comments.Columns().UserId, req.UserId)
	}
	if req.ParentId > 0 {
		m = m.Where(dao.Comments.Columns().ParentId, req.ParentId)
	}
	// 如果 req.Status 有传，映射处理一下，这里简化
	total, err := m.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论总数失败")
	}
	m = m.Page(req.Page, req.PageSize).OrderDesc(dao.Comments.Columns().CreatedAt)
	var list []*entity.Comments
	err = m.Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论列表失败")
	}
	// 批量查询当前用户的评论点赞状态
	var likedCommentIds map[uint64]struct{}
	if userId, e := getUserIdFromCtx(ctx); e == nil && userId > 0 {
		commentIds := make([]uint64, 0, len(list))
		for _, c := range list {
			commentIds = append(commentIds, c.Id)
		}
		var likedRecords []entity.UserLikes
		_ = dao.UserLikes.Ctx(ctx).
			Where(dao.UserLikes.Columns().UserId, userId).
			Where(dao.UserLikes.Columns().TargetType, "comment").
			Where(dao.UserLikes.Columns().TargetId, commentIds).
			Scan(&likedRecords)
		likedCommentIds = make(map[uint64]struct{}, len(likedRecords))
		for _, r := range likedRecords {
			likedCommentIds[r.TargetId] = struct{}{}
		}
	}

	items := make([]*model.ForumCommentViewParams, 0, len(list))
	for _, c := range list {
		statusStr := "normal"
		if c.Status == 0 {
			statusStr = "hidden"
		}
		var liked bool
		if likedCommentIds != nil {
			_, liked = likedCommentIds[c.Id]
		}
		items = append(items, &model.ForumCommentViewParams{
			Id:        c.Id,
			UserId:    c.UserId,
			PostId:    c.TargetId,
			ParentId:  c.ParentId,
			Content:   c.Content,
			Status:    statusStr,
			LikeCount: c.LikeCount,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			IsLiked:   liked,
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
