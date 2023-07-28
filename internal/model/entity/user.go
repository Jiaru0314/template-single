// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int64       `json:"id"         description:"编码"`
	Username  string      `json:"username"   description:"用户名"`
	Password  string      `json:"password"   description:"密码"`
	NickName  string      `json:"nick_name"  description:"昵称"`
	Phone     string      `json:"phone"      description:"手机号"`
	Avatar    string      `json:"avatar"     description:"头像"`
	Sex       string      `json:"sex"        description:"性别"`
	Email     string      `json:"email"      description:"邮箱"`
	Remark    string      `json:"remark"     description:"备注"`
	Status    string      `json:"status"     description:"状态"`
	CreateBy  int64       `json:"create_by"  description:"创建者"`
	UpdateBy  int64       `json:"update_by"  description:"更新者"`
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"最后更新时间"`
	DeletedAt *gtime.Time `json:"deleted_at" description:"删除时间"`
}
