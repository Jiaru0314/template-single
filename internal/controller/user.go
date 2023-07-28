package controller

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/Jiaru0314/template-single/api"

	"github.com/Jiaru0314/template-single/internal/model"
	"github.com/Jiaru0314/template-single/internal/service"
)

var User = cUser{}

type cUser struct{}

func (*cUser) Add(ctx context.Context, req *api.AddUserReq) (res *api.AddUserRes, err error) {
	input := model.AddUserInput{}
	if err = gconv.Struct(req, &input); err != nil {
		return nil, err
	}

	out, err := service.User().Add(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.AddUserRes{Id: out.Id}, nil
}

func (*cUser) Update(ctx context.Context, req *api.UpdateUserReq) (res *api.UpdateUserRes, err error) {
	input := model.UpdateUserInput{}
	if err = gconv.Struct(req, &input); err != nil {
		return nil, err
	}

	if err = service.User().Update(ctx, input); err != nil {

		return nil, err
	}

	return &api.UpdateUserRes{
		Success: true,
	}, nil
}

func (*cUser) Logout(ctx context.Context, req *api.UserLogoutReq) (res *api.UserLogoutRes, err error) {
	service.User().Logout(ctx)
	return &api.UserLogoutRes{
		Success: true,
	}, nil
}

func (*cUser) GetCurrentUserInfo(ctx context.Context, req *api.GetCurrentUserInfoReq) (res *api.UserInfoRes, err error) {
	out, err := service.User().GetCurrentUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	return &api.UserInfoRes{Id: out.Id, Username: out.Username}, nil
}
