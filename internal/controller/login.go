package controller

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/Jiaru0314/template-single/api"
	"github.com/Jiaru0314/template-single/internal/model"
	"github.com/Jiaru0314/template-single/internal/service"
)

var Login = cLogin{}

type cLogin struct{}

func (*cLogin) Login(ctx context.Context, req *api.UserLoginReq) (res *api.UserLoginRes, err error) {
	input := model.UserLoginInput{}
	if err = gconv.Struct(req, &input); err != nil {
		return nil, err
	}

	out, err := service.User().Login(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.UserLoginRes{Success: out.Success, AccessToken: out.AccessToken}, nil
}
