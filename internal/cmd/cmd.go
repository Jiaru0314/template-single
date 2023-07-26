package cmd

import (
	"context"
	"github.com/Jiaru0314/template-single/internal/consts"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
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
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.GET("/swagger", func(r *ghttp.Request) {
					r.Response.Write(consts.SwaggerUIPageContent)
				})

				// 业务路由绑定
				group.Bind()
			})
			s.Run()
			return nil
		},
	}
)
