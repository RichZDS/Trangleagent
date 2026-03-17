package user

import (
	v1 "TrangleAgent/api/user/v1"
	"TrangleAgent/internal/dao"
	"TrangleAgent/internal/model"
	"TrangleAgent/internal/model/entity"
	"TrangleAgent/internal/service"
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type sUser struct{}

func New() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(New())
}

func (s *sUser) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	// 创建数据库查询模型
	m := dao.Users.Ctx(ctx)

	// 按照账号搜索
	if req.Account != "" {
		m = m.WhereLike(dao.Users.Columns().Account, "%"+req.Account+"%")
	}

	// 按照昵称搜索
	if req.Nickname != "" {
		m = m.WhereLike(dao.Users.Columns().Nickname, "%"+req.Nickname+"%")
	}

	// 查询总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页处理
	m = m.Page(req.Page, req.PageSize)

	// 查询列表数据
	var users []*model.User
	err = m.Scan(&users)
	if err != nil {
		return nil, err
	}

	// 转换为视图参数
	userViewList := make([]*model.UserViewParams, 0, len(users))
	for _, user := range users {
		userViewList = append(userViewList, &model.UserViewParams{
			Id:         user.Id,
			Account:    user.Account,
			Nickname:   user.Nickname,
			Gender:     user.Gender,
			BirthDate:  user.BirthDate,
			UserType:   user.UserType,
			VipStartAt: user.VipStartAt,
			VipEndAt:   user.VipEndAt,
			CreatedAt:  user.CreatedAt,
		})
	}

	res = &v1.UserListRes{
		Users: userViewList,
	}

	// 设置分页信息
	req.PageResult.Total = int(total)
	req.PageResult.Page = req.Page
	req.PageResult.PageSize = req.PageSize

	return res, nil
}

func (s *sUser) UserView(ctx context.Context, req *v1.UserViewReq) (res *v1.UserViewRes, err error) {
	m := dao.Users.Ctx(ctx)
	if req.Id != 0 {
		m = m.Where(dao.Users.Columns().Id, req.Id)
	} else if req.Account != "" {
		m = m.Where(dao.Users.Columns().Account, req.Account)
	} else if req.Nickname != "" {
		m = m.Where(dao.Users.Columns().Nickname, req.Nickname)
	}

	res = &v1.UserViewRes{}
	err = m.Scan(&res.UserViewParams)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户信息失败")
	}

	// 若有当前选中角色，附带其 ARC
	if res.ActiveRoleId > 0 {
		var role entity.RoleCards
		err = dao.RoleCards.Ctx(ctx).Where(dao.RoleCards.Columns().Id, res.ActiveRoleId).Scan(&role)
		if err == nil {
			res.ArcAbnormal = role.ArcAbnormal
			res.ArcReality = role.ArcReality
			res.ArcPosition = role.ArcPosition
		}
	}

	return res, nil
}

func (s *sUser) UserUpdate(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	// 创建数据库查询模型
	m := dao.Users.Ctx(ctx)

	// 根据账号更新用户信息
	_, err = m.Data(req).Where(dao.Users.Columns().Account, req.Account).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新用户信息失败")
	}

	res = &v1.UserUpdateRes{}
	// 获取更新后的用户信息
	err = m.Where(dao.Users.Columns().Account, req.Account).Scan(&res)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户信息失败")
	}

	return res, nil
}

func (s *sUser) UserDelete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Account, req.Account).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除用户失败")
	}
	res = &v1.UserDeleteRes{}
	return res, nil
}

// ExpAdd 增加用户经验，等级 = 1 + exp/100
func (s *sUser) ExpAdd(ctx context.Context, req *v1.ExpAddReq) (res *v1.ExpAddRes, err error) {
	_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, req.UserId).Increment(dao.Users.Columns().Exp, int(req.Amount))
	if err != nil {
		return nil, gerror.Wrap(err, "增加经验失败")
	}
	var u entity.Users
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, req.UserId).Scan(&u)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户失败")
	}
	level := uint(1) + u.Exp/100
	_, _ = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, req.UserId).Data(map[string]interface{}{dao.Users.Columns().Level: level}).Update()
	return &v1.ExpAddRes{Exp: u.Exp, Level: level}, nil
}

const checkInExpAmount = 20 // 签到一次 +20 经验

// calcLevel 根据经验计算等级：每 100 经验升 1 级，level = 1 + exp/100
func calcLevel(exp uint) uint {
	if exp == 0 {
		return 1
	}
	return 1 + exp/100
}

// CheckIn 每日签到，+20 经验，等级 = 1 + exp/100
func (s *sUser) CheckIn(ctx context.Context, req *v1.CheckInReq) (res *v1.CheckInRes, err error) {
	var u entity.Users
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, req.UserId).Scan(&u)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户失败")
	}
	// 检查今日是否已签到（按本地日期）
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if u.LastCheckinAt != nil {
		last := u.LastCheckinAt.Time
		lastDay := time.Date(last.Year(), last.Month(), last.Day(), 0, 0, 0, 0, last.Location())
		if !today.After(lastDay) {
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "今日已签到，请明日再来")
		}
	}
	// 使用事务确保原子性，并用 Raw SQL 避免 ORM schema 缓存导致的 "input data match no fields"
	err = dao.Users.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 先更新签到时间（防止重复签到）
		_, e := tx.Exec("UPDATE users SET last_checkin_at = ? WHERE id = ?", gtime.New(now).Format("Y-m-d H:i:s"), req.UserId)
		if e != nil {
			return gerror.Wrap(e, "更新签到时间失败")
		}
		// 2. 增加经验
		_, e = tx.Model("users").Where("id", req.UserId).Increment("exp", checkInExpAmount)
		if e != nil {
			return gerror.Wrap(e, "签到失败")
		}
		// 3. 查询最新 exp 并更新 level
		var u2 entity.Users
		e = tx.Model("users").Where("id", req.UserId).Scan(&u2)
		if e != nil {
			return gerror.Wrap(e, "查询用户失败")
		}
		level := calcLevel(u2.Exp)
		_, e = tx.Exec("UPDATE users SET level = ? WHERE id = ?", level, req.UserId)
		if e != nil {
			return gerror.Wrap(e, "更新等级失败")
		}
		// 写回外部变量供返回
		u = u2
		return nil
	})
	if err != nil {
		return nil, err
	}
	level := calcLevel(u.Exp)
	return &v1.CheckInRes{
		Exp:      u.Exp,
		Level:    level,
		AddedExp: checkInExpAmount,
		Message:  "签到成功，获得 20 经验",
	}, nil
}

// RoleCreate 创建角色（每个用户可创建多个角色，每个角色拥有不同的素质保障）
func (s *sUser) RoleCreate(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleCreateRes, err error) {
	data := &entity.RoleCards{
		UserId:         req.UserId,
		DepartmentId:   req.DepartmentId,
		AgentName:      req.AgentName,
		CodeName:       req.CodeName,
		Gender:         req.Gender,
		ArcAbnormal:    req.ArcAbnormal,
		ArcReality:     req.ArcReality,
		ArcPosition:    req.ArcPosition,
		QaMeticulous:   req.QaMeticulous,
		QaDeception:    req.QaDeception,
		QaVigor:        req.QaVigor,
		QaEmpathy:      req.QaEmpathy,
		QaInitiative:   req.QaInitiative,
		QaResilience:   req.QaResilience,
		QaPresence:     req.QaPresence,
		QaProfessional: req.QaProfessional,
		QaDiscretion:   req.QaDiscretion,
	}

	id, err := dao.RoleCards.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "创建角色失败")
	}

	res = &v1.RoleCreateRes{
		Id: uint64(id),
	}
	return res, nil
}

// RoleUpdate 更新角色（含素质保障 QA 字段）
func (s *sUser) RoleUpdate(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	data := &entity.RoleCards{
		DepartmentId:   req.DepartmentId,
		AgentName:      req.AgentName,
		CodeName:       req.CodeName,
		Gender:         req.Gender,
		ArcAbnormal:    req.ArcAbnormal,
		ArcReality:     req.ArcReality,
		ArcPosition:    req.ArcPosition,
		Commendation:   req.Commendation,
		Reprimand:      req.Reprimand,
		BlueTrack:      req.BlueTrack,
		YellowTrack:    req.YellowTrack,
		RedTrack:       req.RedTrack,
		QaMeticulous:   req.QaMeticulous,
		QaDeception:    req.QaDeception,
		QaVigor:        req.QaVigor,
		QaEmpathy:      req.QaEmpathy,
		QaInitiative:   req.QaInitiative,
		QaResilience:   req.QaResilience,
		QaPresence:     req.QaPresence,
		QaProfessional: req.QaProfessional,
		QaDiscretion:   req.QaDiscretion,
	}

	_, err = dao.RoleCards.Ctx(ctx).Data(data).Where(dao.RoleCards.Columns().Id, req.Id).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "更新角色失败")
	}

	res = &v1.RoleUpdateRes{
		Id: req.Id,
	}
	return res, nil
}

// RoleView 查看角色详情
func (s *sUser) RoleView(ctx context.Context, req *v1.RoleViewReq) (res *v1.RoleViewRes, err error) {
	// 查询角色详情
	var role entity.RoleCards
	err = dao.RoleCards.Ctx(ctx).Where(dao.RoleCards.Columns().Id, req.Id).Scan(&role)
	if err != nil {
		return nil, gerror.Wrap(err, "查询角色详情失败")
	}

	res = &v1.RoleViewRes{
		RoleViewParams: v1.RoleViewParams{
			Id:             role.Id,
			UserId:         role.UserId,
			DepartmentId:   role.DepartmentId,
			Commendation:   role.Commendation,
			Reprimand:      role.Reprimand,
			BlueTrack:      role.BlueTrack,
			YellowTrack:    role.YellowTrack,
			RedTrack:       role.RedTrack,
			ArcAbnormal:    role.ArcAbnormal,
			ArcReality:     role.ArcReality,
			ArcPosition:    role.ArcPosition,
			AgentName:      role.AgentName,
			CodeName:       role.CodeName,
			Gender:         role.Gender,
			QaMeticulous:   role.QaMeticulous,
			QaDeception:    role.QaDeception,
			QaVigor:        role.QaVigor,
			QaEmpathy:      role.QaEmpathy,
			QaInitiative:   role.QaInitiative,
			QaResilience:   role.QaResilience,
			QaPresence:     role.QaPresence,
			QaProfessional: role.QaProfessional,
			QaDiscretion:   role.QaDiscretion,
		},
	}
	return res, nil
}

// RoleList 获取角色列表
func (s *sUser) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	// 创建数据库查询模型
	m := dao.RoleCards.Ctx(ctx)

	// 根据用户ID查询
	if req.UserId != 0 {
		m = m.Where(dao.RoleCards.Columns().UserId, req.UserId)
	}

	// 根据部门ID查询
	if req.DepartmentId != 0 {
		m = m.Where(dao.RoleCards.Columns().DepartmentId, req.DepartmentId)
	}

	// 查询总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页处理
	m = m.Page(req.Page, req.PageSize)

	// 查询列表数据
	var roles []*entity.RoleCards
	err = m.Scan(&roles)
	if err != nil {
		return nil, err
	}

	// 转换为视图参数
	roleViewList := make([]*v1.RoleViewParams, 0, len(roles))
	for _, role := range roles {
		roleViewList = append(roleViewList, &v1.RoleViewParams{
			Id:             role.Id,
			UserId:         role.UserId,
			DepartmentId:   role.DepartmentId,
			Commendation:   role.Commendation,
			Reprimand:      role.Reprimand,
			BlueTrack:      role.BlueTrack,
			YellowTrack:    role.YellowTrack,
			RedTrack:       role.RedTrack,
			ArcAbnormal:    role.ArcAbnormal,
			ArcReality:     role.ArcReality,
			ArcPosition:    role.ArcPosition,
			AgentName:      role.AgentName,
			CodeName:       role.CodeName,
			Gender:         role.Gender,
			QaMeticulous:   role.QaMeticulous,
			QaDeception:    role.QaDeception,
			QaVigor:        role.QaVigor,
			QaEmpathy:      role.QaEmpathy,
			QaInitiative:   role.QaInitiative,
			QaResilience:   role.QaResilience,
			QaPresence:     role.QaPresence,
			QaProfessional: role.QaProfessional,
			QaDiscretion:   role.QaDiscretion,
		})
	}

	res = &v1.RoleListRes{
		List:       roleViewList,
		PageResult: req.PageResult, // 设置分页信息
	}

	// 设置分页信息
	req.PageResult.Total = int(total)
	req.PageResult.Page = req.Page
	req.PageResult.PageSize = req.PageSize

	return res, nil
}

// RoleDelete 删除角色
func (s *sUser) RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	// 根据ID删除角色
	_, err = dao.RoleCards.Ctx(ctx).Where(dao.RoleCards.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除角色失败")
	}

	res = &v1.RoleDeleteRes{
		Id: req.Id,
	}
	return res, nil
}

// RolePermissionCheck 权限查询
func (s *sUser) RolePermissionCheck(ctx context.Context, req *v1.RolePermissionCheckReq) (res *v1.RolePermissionCheckRes, err error) {
	// 查询角色信息
	var role entity.RoleCards
	err = dao.RoleCards.Ctx(ctx).Where(dao.RoleCards.Columns().Id, req.RoleId).Scan(&role)
	if err != nil {
		return nil, gerror.Wrap(err, "查询角色信息失败")
	}

	// 根据轨道类型获取对应的轨道值
	var trackValue uint
	switch req.TrackType {
	case "blue":
		trackValue = role.BlueTrack
	case "yellow":
		trackValue = role.YellowTrack
	case "red":
		trackValue = role.RedTrack
	default:
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "无效的轨道类型")
	}

	// 判断权限：如果角色的轨道值 >= 文件需要的值，则权限验证通过
	if trackValue >= req.FileValue {
		// 权限验证通过
		res = &v1.RolePermissionCheckRes{
			Code: 333,
			Mes:  "权限验证通过",
		}
		return
	}

	// 权限不足，返回错误
	res = &v1.RolePermissionCheckRes{
		Code: 403,
		Mes:  "权限不足",
	}
	return res, nil
}
