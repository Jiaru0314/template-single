package cmd

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/Jiaru0314/template-single/internal/consts"
	"github.com/Jiaru0314/template-single/internal/controller"
	"github.com/Jiaru0314/template-single/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 国际化 全局配置
			i18nConfig()

			// 路由绑定
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.GET("/swagger", func(r *ghttp.Request) {
					r.Response.Write(consts.SwaggerUIPageContent)
				})

				// 业务路由绑定(不鉴权)
				group.Bind(
					controller.Captcha, // 验证码
				)

				// 业务路由绑定(鉴权)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind()
				})
			})
			s.Run()
			return nil
		},
	}
)

func i18nConfig() {
	err := g.I18n().SetPath(consts.I18NPath)
	if err != nil {
		g.Log().Panic(context.Background(), "I18N 全局初始化失败")
	}
	// g.I18n().SetLanguage(consts.LanguageZH)
}
