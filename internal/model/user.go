package model

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

// UserInfo User全部内容
type UserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
}

// AddUserInput 创建用户模型请求内容
type AddUserInput struct {
	UserBase
}

// AddUserOutput 创建用户模型返回响应
type AddUserOutput struct {
	Id int `json:"id"`
}

// UpdateUserInput 修改用户模型请求内容
type UpdateUserInput struct {
	UserBase
	Id int
}

// UpdateUserOutPut 修改用户模型响应
type UpdateUserOutPut struct {
	Success bool
}

// DetailUserInput 根据ID查询用户模型 详情请求内容
type DetailUserInput struct {
	Id int
}

// DetailUserOutput 查询用户模型 详情响应
type DetailUserOutput struct {
	UserInfo
}

// UserLoginInput 用户登录请求内容
type UserLoginInput struct {
	UserName string
	Password string
}

// UserLoginOutput 用户登录请求响应
type UserLoginOutput struct {
	Id          uint
	UserName    string
	Success     bool
	AccessToken string
}
