package service

import (
	"context"

	"github.com/Jiaru0314/template-single/internal/model"
)

type (
	IUser interface {
		Add(ctx context.Context, in model.AddUserInput) (out *model.AddUserOutput, err error)
		Update(ctx context.Context, in model.UpdateUserInput) (err error)

		Login(ctx context.Context, in model.UserLoginInput) (out *model.UserLoginOutput, err error)
		Logout(ctx context.Context) (err error)
		GetCurrentUserInfo(ctx context.Context) (out *model.UserInfo, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
