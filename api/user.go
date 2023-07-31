package api

import "github.com/gogf/gf/v2/frame/g"

// UserBase User基类
type UserBase struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
}

type UserInfoRes struct {
	Id       int    `json:"id" `
	Username string `json:"username"`
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
}

// UserRegisterReq 新增
type UserRegisterReq struct {
	g.Meta `path:"/user/register" method:"post" tags:"用户模型" summary:"注册"`
	UserBase
}

type UserRegisterRes struct {
	Id int `json:"id" dc:"id"`
}

// UpdateUserReq 更新
type UpdateUserReq struct {
	g.Meta `path:"/user" method:"put" tags:"用户模型" summary:"更新"`
	Id     int `json:"id" v:"required#用户模型ID不能为空" dc:"id"`
	UserBase
}

type UpdateUserRes struct {
	Success bool `json:"success"      description:"是否成功"`
}

// UserLoginReq 用户登录请求内容
type UserLoginReq struct {
	g.Meta   `path:"/user/login" method:"post" tags:"用户模型" summary:"用户登录"`
	UserName string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

// UserLoginRes 用户登录请求响应
type UserLoginRes struct {
	Success     bool
	AccessToken string
}

// UserLogoutReq 用户退出登录请求
type UserLogoutReq struct {
	g.Meta `path:"/user/logout" method:"post" tags:"用户模型" summary:"用户退出登录"`
}

// UserLogoutRes 用户退出登录响应
type UserLogoutRes struct {
	Success bool
}

type GetCurrentUserInfoReq struct {
	g.Meta `path:"/user/currentUser" method:"get" tags:"用户模型" summary:"获取当前用户信息"`
}

// UpdatePasswordReq 更新
type UpdatePasswordReq struct {
	g.Meta   `path:"/user/password/update" method:"put" tags:"用户模型" summary:"修改密码"`
	Id       int    `json:"id" v:"required#用户模型ID不能为空" dc:"id"`
	Password string `json:"password" v:"required#用户密码不能为空" dc:"密码"`
}
