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

// Register 用户注册
func (*cUser) Register(ctx context.Context, req *api.UserRegisterReq) (res *api.UserRegisterRes, err error) {
	input := model.AddUserInput{}
	if err = gconv.Struct(req, &input); err != nil {
		return nil, err
	}

	out, err := service.User().Add(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.UserRegisterRes{Id: out.Id}, nil
}

// UpdateUserInfo 更新用户信息
func (*cUser) UpdateUserInfo(ctx context.Context, req *api.UpdateUserReq) (res *api.UpdateUserRes, err error) {
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

// Logout 用户注销
func (*cUser) Logout(ctx context.Context, req *api.UserLogoutReq) (res *api.UserLogoutRes, err error) {
	service.User().Logout(ctx)
	return &api.UserLogoutRes{
		Success: true,
	}, nil
}

// GetCurrentUserInfo 获取当前用户信息
func (*cUser) GetCurrentUserInfo(ctx context.Context, req *api.GetCurrentUserInfoReq) (res *api.UserInfoRes, err error) {
	out, err := service.User().GetCurrentUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	return &api.UserInfoRes{Id: out.Id, Username: out.Username}, nil
}

// UpdatePassword 更新用户密码
func (*cUser) UpdatePassword(ctx context.Context, req *api.UpdatePasswordReq) (res *api.UpdateUserRes, err error) {
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
