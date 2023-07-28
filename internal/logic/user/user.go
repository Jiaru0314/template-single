package User

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/Jiaru0314/template-single/internal/dao"
	"github.com/Jiaru0314/template-single/internal/model"
	"github.com/Jiaru0314/template-single/internal/service"
)

type sUser struct{}

func (u *sUser) GetCurrentUserInfo(ctx context.Context) (out *model.UserInfo, err error) {
	payload := service.Auth().GetPayload(ctx)

	var res model.UserInfo
	var payloadMap map[string]interface{}
	err = json.Unmarshal([]byte(payload), &payloadMap)
	if err != nil {
		fmt.Println("解析 payload 失败:", err)
		return nil, err
	}

	res = model.UserInfo{
		Id:       gconv.Int(payloadMap["id"]),
		Username: payloadMap["userName"].(string),
		NickName: payloadMap["nickname"].(string),
		Phone:    payloadMap["phone"].(string),
		Avatar:   payloadMap["avatar"].(string),
		Email:    payloadMap["email"].(string),
		Sex:      payloadMap["sex"].(string),
	}

	return &res, nil
}

func (u *sUser) Login(ctx context.Context, in model.UserLoginInput) (out *model.UserLoginOutput, err error) {
	tokenString, _ := service.Auth().LoginHandler(ctx)
	return &model.UserLoginOutput{Success: true, AccessToken: tokenString}, nil
}

func (u *sUser) Logout(ctx context.Context) (err error) {
	service.Auth().LogoutHandler(ctx)
	return nil
}

func init() {
	service.RegisterUser(&sUser{})
}

func (*sUser) Add(ctx context.Context, in model.AddUserInput) (out *model.AddUserOutput, err error) {
	id, err := dao.User.Ctx(ctx).Data(in.UserBase).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.AddUserOutput{Id: int(id)}, err
}

func (*sUser) Update(ctx context.Context, in model.UpdateUserInput) (err error) {
	if _, err = dao.User.Ctx(ctx).Data(in).FieldsEx(in.Id).Where(dao.User.Columns().Id, in.Id).Update(); err != nil {
		return err
	}

	return nil
}
