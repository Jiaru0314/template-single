package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/Jiaru0314/template-single/internal/service"
)

type sMiddleware struct{}

func (s sMiddleware) ResponseHandler(r *ghttp.Request) {
	// TODO implement me
	panic("implement me")
}

func (s sMiddleware) Ctx(r *ghttp.Request) {
	// TODO implement me
	panic("implement me")
}

func (s sMiddleware) Auth(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

func init() {
	service.RegisterMiddleware(New())
}

// New Captcha 验证码管理服务
func New() *sMiddleware {
	return &sMiddleware{}
}
