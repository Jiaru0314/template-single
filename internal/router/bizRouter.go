package router

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/Jiaru0314/template-single/internal/controller"
	"github.com/Jiaru0314/template-single/internal/service"
)

func RegisterAuthBizRouter(group *ghttp.RouterGroup) {
	group.Middleware(service.Middleware().Auth)
	group.Bind(
		controller.User, // 用户接口
	)
}
