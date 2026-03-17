package cmd

import (
	"TrangleAgent/internal/controller"
	"TrangleAgent/internal/controller/websocket"
	"TrangleAgent/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// CORS 中间件 - 排除 swagger 和 openapi 路径
			s.Use(func(r *ghttp.Request) {
				// 排除 swagger 相关路径
				path := r.URL.Path
				if path == "/swagger" || path == "/api.json" || path == "/swagger/" || path == "/api.json/" {
					r.Middleware.Next()
					return
				}

				r.Response.CORS(ghttp.CORSOptions{
					AllowOrigin:      "http://localhost:3000,http://localhost:8080",
					AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
					AllowHeaders:     "Content-Type,Authorization,X-Requested-With",
					AllowCredentials: "true",
					MaxAge:           3600,
				})
				if r.Method == "OPTIONS" {
					r.Response.WriteStatusExit(200)
					return
				}
				r.Middleware.Next()
			})

			// 注册 API 路由组
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.ResponseHandler)
				controller.RegisterControllers(group)
			})

			// 注册 WebSocket 聊天室路由（不在 /api 组下，因为 WebSocket 不需要中间件）
			s.BindHandler("/ws/chat", websocket.HandleChatConnections)

			s.SetDumpRouterMap(false)

			s.Run()
			return nil
		},
	}
)
