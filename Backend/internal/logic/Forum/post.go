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

// sForumPosts 帖子逻辑
type sForumPosts struct{}

func NewForumPosts() *sForumPosts {
	return &sForumPosts{}
}

func init() {
	service.RegisterForumPosts(NewForumPosts())
}

func (s *sForumPosts) Like(ctx context.Context, req *v1.ForumPostsLikeReq) (res *v1.ForumPostsLikeRes, err error) {
	// 帖子点赞数+1
	_, err = dao.ForumPosts.Ctx(ctx).Where(dao.ForumPosts.Columns().Id, req.Id).Increment(dao.ForumPosts.Columns().LikeCount, 1)
	if err != nil {
		return nil, gerror.Wrap(err, "点赞失败")
	}
	return &v1.ForumPostsLikeRes{}, nil
}

func (s *sForumPosts) Create(ctx context.Context, req *v1.ForumPostsCreateReq) (res *v1.ForumPostsCreateRes, err error) {

	//校验参数
	if len(req.Title) < 5 || len(req.Title) > 100 {
		return nil, gerror.New("标题长度要在5-100之间")
	}
	if req.BoardId == 0 {
		return nil, gerror.New("板块ID不能为空")
	}

	//创建帖子
	insert := entity.ForumPosts{
		BoardId:    req.BoardId,
		UserId:     req.UserId,
		Title:      req.Title,
		Content:    req.Content,
		CoverImage: req.CoverImage,
		Status:     req.Status,
		CreatedAt:  gtime.Now(),
		UpdatedAt:  gtime.Now(),
	}

	if insert.Status == "" {
		insert.Status = "normal"
	}

	result, err := dao.ForumPosts.Ctx(ctx).Data(insert).Insert()
	if err != nil {
		return nil, gerror.Wrap(err, "创建帖子失败")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.Wrap(err, "获取帖子ID失败")
	}

	return &v1.ForumPostsCreateRes{Id: uint64(id)}, nil

}

func (s *sForumPosts) Update(ctx context.Context, req *v1.ForumPostsUpdateReq) (res *v1.ForumPostsUpdateRes, err error) {
	//校验参数
	if len(req.Title) < 5 || len(req.Title) > 100 {
		return nil, gerror.New("标题长度要在5-100之间")
	}
	if len(req.Content) < 5 || len(req.Content) > 10000 {
		return nil, gerror.New("内容长度要在5-10000之间")
	}
	//只能修改 帖子标题 帖子正文 帖子封面图URL
	//更新帖子
	_, err = dao.ForumPosts.Ctx(ctx).Data(entity.ForumPosts{
		Id:         req.Id,
		Title:      req.Title,
		Content:    req.Content,
		CoverImage: req.CoverImage,
		UpdatedAt:  gtime.Now(),
	}).Fields("title", "content", "cover_image", "updated_at").OmitEmpty().Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新帖子失败")
	}

	return &v1.ForumPostsUpdateRes{Id: req.Id}, nil
}

func (s *sForumPosts) Delete(ctx context.Context, req *v1.ForumPostsDeleteReq) (res *v1.ForumPostsDeleteRes, err error) {
	//根据id删除帖子
	_, err = dao.ForumPosts.Ctx(ctx).Where(dao.ForumPosts.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除帖子失败")
	}
	return &v1.ForumPostsDeleteRes{}, nil
}

func (s *sForumPosts) View(ctx context.Context, req *v1.ForumPostsViewReq) (res *v1.ForumPostsViewRes, err error) {
	// 根据id查询帖子
	var post model.ForumPostViewParams
	err = dao.ForumPosts.Ctx(ctx).Where(dao.ForumPosts.Columns().Id, req.Id).Scan(&post)
	if err != nil {
		return nil, gerror.Wrap(err, "查询帖子失败")
	}
	// 浏览量 +1
	_, _ = dao.ForumPosts.Ctx(ctx).Where(dao.ForumPosts.Columns().Id, req.Id).Increment(dao.ForumPosts.Columns().ViewCount, 1)
	post.ViewCount++
	// 查询用户名称
	type User struct {
		Name string `json:"name" orm:"nickname" description:"用户昵称"`
	}
	var user User
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, post.UserId).Scan(&user)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户昵称失败")
	}
	return &v1.ForumPostsViewRes{
		UserName:            user.Name,
		ForumPostViewParams: post,
	}, nil
}

func (s *sForumPosts) List(ctx context.Context, req *v1.ForumPostsListReq) (res *v1.ForumPostsListRes, err error) {
	// 查询帖子列表

	mod := dao.ForumPosts.Ctx(ctx)

	if req.Title != "" {
		mod = mod.WhereLike(dao.ForumPosts.Columns().Title, "%"+req.Title+"%")
	}
	if req.BoardId != 0 {
		mod = mod.Where(dao.ForumPosts.Columns().BoardId, req.BoardId)
	}
	if req.UserId != 0 {
		mod = mod.Where(dao.ForumPosts.Columns().UserId, req.UserId)
	}
	if req.Status != "" {
		mod = mod.Where(dao.ForumPosts.Columns().Status, req.Status)
	}
	if req.IsPinned != 0 {
		mod = mod.Where(dao.ForumPosts.Columns().IsPinned, req.IsPinned)
	}
	if req.IsEssence != 0 {
		mod = mod.Where(dao.ForumPosts.Columns().IsEssence, req.IsEssence)
	}

	// 排序：优先置顶，然后按创建时间倒序
	mod = mod.OrderDesc(dao.ForumPosts.Columns().IsPinned).OrderDesc(dao.ForumPosts.Columns().CreatedAt)

	// 分页
	page := req.Page
	if page <= 0 {
		page = 1
	}
	size := req.PageSize
	if size <= 0 {
		size = 10
	}

	total, err := mod.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取帖子总数失败")
	}

	var list []*model.ForumPostViewParams
	err = mod.Page(page, size).Scan(&list)
	if err != nil {
		return nil, gerror.Wrap(err, "查询帖子列表失败")
	}
	// 查询用户名称
	if len(list) > 0 {
		type User struct {
			Id   uint64 `json:"id"   orm:"id"`
			Name string `json:"name" orm:"nickname"`
		}
		var users []User

		// 构造去重后的用户 ID 切片
		idSet := make(map[uint64]struct{}, len(list))
		for _, item := range list {
			idSet[item.UserId] = struct{}{}
		}

		userIds := make([]uint64, 0, len(idSet))
		for id := range idSet {
			userIds = append(userIds, id)
		}

		err = dao.Users.Ctx(ctx).
			Fields(dao.Users.Columns().Id, dao.Users.Columns().Nickname).
			Where(dao.Users.Columns().Id, userIds).
			Scan(&users)
		if err != nil {
			return nil, gerror.Wrap(err, "查询用户名称失败")
		}

		// 建立 userId -> name 映射
		userMap := make(map[uint64]string, len(users))
		for _, user := range users {
			userMap[user.Id] = user.Name
		}

		// 关联用户名称
		for _, item := range list {
			if name, ok := userMap[item.UserId]; ok {
				item.Name = name
			}
		}
	}
	return &v1.ForumPostsListRes{
		PageResult: response.PageResult{
			Total:    total,
			Page:     page,
			PageSize: size,
		},
		List: list,
	}, nil
}
