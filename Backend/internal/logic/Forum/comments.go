package Forum

import (
	v1 "TrangleAgent/api/forum/v1"
	"TrangleAgent/internal/dao"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sForumComments struct{}

func NewForumComments() *sForumComments {
	return &sForumComments{}
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
	return nil, nil
}

func (s *sForumComments) Update(ctx context.Context, req *v1.ForumCommentsUpdateReq) (res *v1.ForumCommentsUpdateRes, err error) {
	return nil, nil
}

func (s *sForumComments) Delete(ctx context.Context, req *v1.ForumCommentsDeleteReq) (res *v1.ForumCommentsDeleteRes, err error) {
	return nil, nil
}

func (s *sForumComments) View(ctx context.Context, req *v1.ForumCommentsViewReq) (res *v1.ForumCommentsViewRes, err error) {
	return nil, nil
	}
func (s *sForumComments) List(ctx context.Context, req *v1.ForumCommentsListReq) (res *v1.ForumCommentsListRes, err error) {
	return nil, nil
}
