package models

import "blog/models/ctypes"

type UserModel struct {
	MODEL
	Nickname string      `json:"nickname" gorm:"comment:昵称"`
	Account  string      `json:"account" gorm:"comment:账号"`
	Password string      `json:"password" gorm:"comment:密码"`
	Avatar   string      `json:"avatar" gorm:"comment:头像"`
	Email    string      `json:"email" gorm:"comment:邮箱"`
	Address  string      `json:"address" gorm:"comment:地址"`
	Token    string      `json:"token"`
	Role     ctypes.Role `json:"role" gorm:"comment:身份,1管理员,2用户"`
}
