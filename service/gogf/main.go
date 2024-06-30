package gogf

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gsession"
)

func runServe(ctx context.Context) {
	s := g.Server()

	// session
	s.SetSessionMaxAge(time.Hour * 24 * 30)
	s.SetSessionStorage(gsession.NewStorageMemory())

	// 静态资源
	s.AddStaticPath("/assets", "resource/public/dist/assets") // 支持直接设置 gre 中的路径

	// 开发模式下，开启 openapi
	if IsDevelop() {
		runOpenApi(s)
	} else {
		// 生产模式下，关闭路由映射列表
		s.SetDumpRouterMap(false)
	}

	// 解决 vue 静态资源 hashRoute 刷新问题
	s.BindHandler("/*", func(r *ghttp.Request) {
		r.Response.ServeFile( "resource/public/dist/index.html")
	})

	// 项目接口
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			// compat.NewV1(),
		)
	})
	s.Run()
}

func runOpenApi(s *ghttp.Server) {
	s.SetOpenApiPath("/api.json")
	s.AddSearchPath("resource/public/resource") // 支持直接设置 gre 中的路径
	s.BindHandler("/apidoc", func(r *ghttp.Request) {
		r.Response.WriteTpl("/apidoc.html")
	})
}



func Main() {
	runServe(context.Background())
}