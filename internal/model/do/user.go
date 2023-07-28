// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        interface{} // 编码
	Username  interface{} // 用户名
	Password  interface{} // 密码
	NickName  interface{} // 昵称
	Phone     interface{} // 手机号
	Avatar    interface{} // 头像
	Sex       interface{} // 性别
	Email     interface{} // 邮箱
	Remark    interface{} // 备注
	Status    interface{} // 状态
	CreateBy  interface{} // 创建者
	UpdateBy  interface{} // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 最后更新时间
	DeletedAt *gtime.Time // 删除时间
}
