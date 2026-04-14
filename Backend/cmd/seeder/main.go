package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "TrangleAgent/internal/packed"
	_ "TrangleAgent/internal/logic"

	"TrangleAgent/internal/dao"
	"TrangleAgent/internal/model/entity"
	
	"github.com/gogf/gf/v2/frame/g"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/brianvoe/gofakeit/v6"
)

// Seeder 用于填充假数据
func main() {
	// Initialize gofakeit
	gofakeit.Seed(time.Now().UnixNano())
	ctx := gctx.GetInitCtx()

	fmt.Println("开始生成假数据...")

	// 1. 生成用户 (Users)
	fmt.Println("-> 生成用户数据...")
	var userIDs []uint64
	for i := 0; i < 20; i++ {
		password, _ := gmd5.Encrypt("123456")
		
		extraInfo := fmt.Sprintf(`{"realityRole":"%s", "abnormalRole":"%s", "jobTitle":"%s"}`,
			gofakeit.JobDescriptor(), gofakeit.HackerAdjective(), gofakeit.JobTitle())
			
		user := entity.Users{
			Account:      gofakeit.Username() + fmt.Sprintf("%d", gofakeit.Number(10,999)),
			Password:     password,
			Nickname:     gofakeit.Name(),
			Gender:       gofakeit.Number(0, 2),
			UserType:     "user",
			Email:        gofakeit.Email(),
			Exp:          uint(gofakeit.Number(0, 5000)),
			Level:        uint(gofakeit.Number(1, 50)),
			ExtraInfo:    extraInfo,
			CreatedAt:    gtime.Now(),
			UpdatedAt:    gtime.Now(),
		}

		res, err := dao.Users.Ctx(ctx).Data(user).Insert()
		if err != nil {
			log.Printf("插入用户失败: %v", err)
			continue
		}
		
		id, _ := res.LastInsertId()
		userIDs = append(userIDs, uint64(id))
	}
	fmt.Printf("成功插入 %d 个用户\n", len(userIDs))

	if len(userIDs) < 2 {
		log.Fatal("需要至少2个用户才能生成关注和评论关系")
	}

	// 2. 生成关注关系 (UserFollows)
	fmt.Println("-> 生成关注关系数据...")
	followCount := 0
	for i := 0; i < 30; i++ {
		followerIdx := rand.Intn(len(userIDs))
		followedIdx := rand.Intn(len(userIDs))
		
		if followerIdx == followedIdx {
			continue // 不能关注自己
		}

		follow := entity.UserFollows{
			FollowerId: userIDs[followerIdx],
			FollowedId: userIDs[followedIdx],
			Status:     1,
			Remark:     gofakeit.Adjective(),
			CreatedAt:  gtime.Now(),
			UpdatedAt:  gtime.Now(),
		}
		
		_, err := dao.UserFollows.Ctx(ctx).Data(follow).Insert()
		if err == nil {
			followCount++
		}
	}
	fmt.Printf("成功插入 %d 条关注关系\n", followCount)

	// 3. 生成角色卡 (RoleCards)
	fmt.Println("-> 生成角色卡数据...")
	var roleIDs []uint64
	for _, uid := range userIDs {
		// 给一半的用户生成角色卡
		if gofakeit.Bool() {
			extraInfo := fmt.Sprintf(`{"backstory":"%s"}`, gofakeit.Paragraph(1, 2, 5, " "))
			role := entity.RoleCards{
				UserId:         uid,
				DepartmentId:   1, // 假定一个默认部门
				Commendation:   uint(gofakeit.Number(0, 10)),
				Reprimand:      uint(gofakeit.Number(0, 5)),
				BlueTrack:      uint(gofakeit.Number(0, 40)),
				YellowTrack:    uint(gofakeit.Number(0, 40)),
				RedTrack:       uint(gofakeit.Number(0, 40)),
				ArcAbnormal:    gofakeit.HackerAdjective(),
				ArcReality:     gofakeit.JobDescriptor(),
				ArcPosition:    gofakeit.JobTitle(),
				AgentName:      gofakeit.Name(),
				CodeName:       gofakeit.Word(),
				Gender:         gofakeit.Gender(),
				QaMeticulous:   uint(gofakeit.Number(0, 100)),
				QaDeception:    uint(gofakeit.Number(0, 100)),
				QaVigor:        uint(gofakeit.Number(0, 100)),
				QaEmpathy:      uint(gofakeit.Number(0, 100)),
				QaInitiative:   uint(gofakeit.Number(0, 100)),
				QaResilience:   uint(gofakeit.Number(0, 100)),
				QaPresence:     uint(gofakeit.Number(0, 100)),
				QaProfessional: uint(gofakeit.Number(0, 100)),
				QaDiscretion:   uint(gofakeit.Number(0, 100)),
				ExtraInfo:      extraInfo,
				CreatedAt:      gtime.Now(),
				UpdatedAt:      gtime.Now(),
			}

			res, err := dao.RoleCards.Ctx(ctx).Data(role).Insert()
			if err == nil {
				id, _ := res.LastInsertId()
				roleIDs = append(roleIDs, uint64(id))
				
				// 顺便更新用户的ActiveRoleId
				dao.Users.Ctx(ctx).Data(g.Map{"active_role_id": id}).Where("id", uid).Update()
			}
		}
	}
	fmt.Printf("成功插入 %d 张角色卡\n", len(roleIDs))

	// 4. 生成帖子版块 (为了能够生成帖子)
	fmt.Println("-> 生成论坛版块数据...")
	section := entity.ForumSections{
		Name:        "闲聊灌水",
		Description: "随便聊聊",
		Status:      "normal",
		CreatedAt:   gtime.Now(),
	}
	res, _ := dao.ForumSections.Ctx(ctx).Data(section).Insert()
	sectionId, _ := res.LastInsertId()

	board := entity.ForumBoards{
		SectionId:   uint64(sectionId),
		Name:        "主板块",
		Description: "欢迎来到主板块",
		Status:      "normal",
		CreatedAt:   gtime.Now(),
	}
	res, _ = dao.ForumBoards.Ctx(ctx).Data(board).Insert()
	boardId, _ := res.LastInsertId()

	// 5. 生成帖子与评论 (多态)
	fmt.Println("-> 生成帖子与多态评论数据...")
	postCount := 0
	commentCount := 0
	for i := 0; i < 10; i++ {
		authorIdx := rand.Intn(len(userIDs))
		post := entity.ForumPosts{
			BoardId:   uint64(boardId),
			UserId:    userIDs[authorIdx],
			Title:     gofakeit.Sentence(5),
			Content:   gofakeit.Paragraph(1, 3, 10, "\n"),
			Status:    "normal",
			ViewCount: uint(gofakeit.Number(10, 1000)),
			LikeCount: uint(gofakeit.Number(0, 100)),
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		}

		res, err := dao.ForumPosts.Ctx(ctx).Data(post).Insert()
		if err == nil {
			postCount++
			postId, _ := res.LastInsertId()

			// 为每个帖子生成几条评论 (多态关联)
			for j := 0; j < gofakeit.Number(1, 5); j++ {
				commenterIdx := rand.Intn(len(userIDs))
				comment := entity.Comments{
					UserId:     userIDs[commenterIdx],
					TargetType: "post",
					TargetId:   uint64(postId),
					ParentId:   0, // 暂作一级评论
					RootId:     0,
					Content:    gofakeit.Sentence(10),
					Status:     1,
					LikeCount:  uint(gofakeit.Number(0, 20)),
					CreatedAt:  gtime.Now(),
					UpdatedAt:  gtime.Now(),
				}
				_, err := dao.Comments.Ctx(ctx).Data(comment).Insert()
				if err == nil {
					commentCount++
				}
			}
			
			// 同步帖子评论数
			dao.ForumPosts.Ctx(ctx).Data(g.Map{"comment_count": commentCount}).Where("id", postId).Update()
		}
	}
	fmt.Printf("成功插入 %d 篇帖子，%d 条多态评论\n", postCount, commentCount)
	
	fmt.Println("假数据生成完成！")
}
