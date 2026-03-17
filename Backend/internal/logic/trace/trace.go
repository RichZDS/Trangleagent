package trace

import (
	"context"
	"fmt"
	v1 "TrangleAgent/api/user/v1"
	"TrangleAgent/internal/dao"
	"TrangleAgent/internal/model/entity"
	"TrangleAgent/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

type sTrace struct{}

func NewTrace() *sTrace {
	return &sTrace{}
}

func init() {
	service.RegisterTrace(NewTrace())
}

// TraceList 获取轨迹列表
func (s *sTrace) TraceList(ctx context.Context, req *v1.TraceListReq) (res *v1.TraceListRes, err error) {
	// 创建数据库查询模型
	m := dao.RoleCards.Ctx(ctx)

	// 根据角色ID查询
	if req.RoleId != 0 {
		m = m.Where(dao.RoleCards.Columns().Id, req.RoleId)
	}

	// 查询总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页处理
	m = m.Page(req.Page, req.PageSize)

	// 查询列表数据 (Scan 到实体)
	var roleCards []*entity.RoleCards
	err = m.Scan(&roleCards)
	if err != nil {
		return nil, err
	}

	// 转换为视图参数
	traceList := make([]*v1.TraceViewParams, 0, len(roleCards))
	for _, card := range roleCards {
		traceList = append(traceList, &v1.TraceViewParams{
			UserId:       card.UserId,
			RoleId:       card.Id,
			DepartmentId: card.DepartmentId,
			RedTrace:     int(card.RedTrack),
			YellowTrace:  int(card.YellowTrack),
			BlueTrace:    int(card.BlueTrack),
		})
	}

	res = &v1.TraceListRes{
		List: traceList,
	}

	// 设置分页信息
	req.PageResult.Total = int(total)
	req.PageResult.Page = req.Page
	req.PageResult.PageSize = req.PageSize

	return
}

// TraceView 查看轨迹详情
func (s *sTrace) TraceView(ctx context.Context, req *v1.TraceViewReq) (res *v1.TraceViewRes, err error) {
	// 创建数据库查询模型
	m := dao.RoleCards.Ctx(ctx)

	// 根据角色ID查询
	var card entity.RoleCards
	err = m.Where(dao.RoleCards.Columns().Id, req.RoleId).Scan(&card)
	if err != nil {
		return nil, err
	}

	res = &v1.TraceViewRes{
		TraceViewParams: v1.TraceViewParams{
			UserId:       card.UserId,
			RoleId:       card.Id,
			DepartmentId: card.DepartmentId,
			RedTrace:     int(card.RedTrack),
			YellowTrace:  int(card.YellowTrack),
			BlueTrace:    int(card.BlueTrack),
		},
	}

	return
}

// TraceUpdate 更新轨迹信息 (增加)
func (s *sTrace) TraceUpdate(ctx context.Context, req *v1.TraceUpdateReq) (res *v1.TraceUpdateRes, err error) {
	// 创建数据库查询模型
	m := dao.RoleCards.Ctx(ctx)

	// 准备更新数据
	data := make(map[string]interface{})

	// 只有当数值不为0时才进行更新操作
	if req.RedTrace != 0 {
		data[dao.RoleCards.Columns().RedTrack] = gdb.Raw(fmt.Sprintf("%s + %d", dao.RoleCards.Columns().RedTrack, req.RedTrace))
	}

	if req.YellowTrace != 0 {
		data[dao.RoleCards.Columns().YellowTrack] = gdb.Raw(fmt.Sprintf("%s + %d", dao.RoleCards.Columns().YellowTrack, req.YellowTrace))
	}

	if req.BlueTrace != 0 {
		data[dao.RoleCards.Columns().BlueTrack] = gdb.Raw(fmt.Sprintf("%s + %d", dao.RoleCards.Columns().BlueTrack, req.BlueTrace))
	}

	// 如果没有需要更新的字段，直接返回
	if len(data) == 0 {
		return &v1.TraceUpdateRes{
			RoleId:      req.RoleId,
			RedTrace:    req.RedTrace,
			YellowTrace: req.YellowTrace,
			BlueTrace:   req.BlueTrace,
		}, nil
	}

	// 添加更新时间
	data[dao.RoleCards.Columns().UpdatedAt] = gtime.Now()

	// 执行更新
	_, err = m.Data(data).Where(dao.RoleCards.Columns().Id, req.RoleId).Update()
	if err != nil {
		return nil, err
	}

	res = &v1.TraceUpdateRes{
		RoleId:      req.RoleId,
		RedTrace:    req.RedTrace,
		YellowTrace: req.YellowTrace,
		BlueTrace:   req.BlueTrace,
	}

	return
}

// TraceReduce 减少轨迹数值
func (s *sTrace) TraceReduce(ctx context.Context, req *v1.TraceReduceReq) (res *v1.TraceReduceRes, err error) {
	// 创建数据库查询模型
	m := dao.RoleCards.Ctx(ctx)

	// 准备更新数据
	data := make(map[string]interface{})

	// 只有当数值不为0时才进行减少操作
	if req.RedTrace != 0 {
		data[dao.RoleCards.Columns().RedTrack] = gdb.Raw(fmt.Sprintf("GREATEST(0, %s - %d)", dao.RoleCards.Columns().RedTrack, req.RedTrace))
	}

	if req.YellowTrace != 0 {
		data[dao.RoleCards.Columns().YellowTrack] = gdb.Raw(fmt.Sprintf("GREATEST(0, %s - %d)", dao.RoleCards.Columns().YellowTrack, req.YellowTrace))
	}

	if req.BlueTrace != 0 {
		data[dao.RoleCards.Columns().BlueTrack] = gdb.Raw(fmt.Sprintf("GREATEST(0, %s - %d)", dao.RoleCards.Columns().BlueTrack, req.BlueTrace))
	}

	// 如果没有需要更新的字段，直接返回
	if len(data) == 0 {
		return &v1.TraceReduceRes{}, nil
	}

	// 添加更新时间
	data[dao.RoleCards.Columns().UpdatedAt] = gtime.Now()

	// 执行更新
	_, err = m.Data(data).Where(dao.RoleCards.Columns().Id, req.RoleId).Update()
	if err != nil {
		return nil, err
	}

	res = &v1.TraceReduceRes{}
	return
}
