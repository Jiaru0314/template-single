package controller

import (
	"context"

	"github.com/Jiaru0314/template-single/api"
	"github.com/Jiaru0314/template-single/internal/consts"
	"github.com/Jiaru0314/template-single/internal/service"
)

// Captcha 图形验证码
var Captcha = cCaptcha{}

type cCaptcha struct{}

func (a *cCaptcha) Index(ctx context.Context, req *api.CaptchaIndexReq) (res *api.CaptchaIndexRes, err error) {
	err = service.Captcha().NewAndStore(ctx, consts.CaptchaDefaultName)
	return
}
