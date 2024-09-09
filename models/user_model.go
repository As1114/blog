package models

import "blog/models/ctype"

type UserModel struct {
	MODEL
	Nickname string     `json:"nick_name" gorm:"comment:昵称,select(c)"`
	Account  string     `json:"account" gorm:"comment:账号"`
	Password string     `json:"password" gorm:"comment:密码"`
	Avatar   string     `json:"avatar" gorm:"comment:头像,select(c)"`
	Email    string     `json:"email" gorm:"comment:邮箱"`
	Address  string     `json:"address" gorm:"comment:地址,select(c)"`
	Token    string     `json:"token"`
	Role     ctype.Role `json:"role" gorm:"comment:身份,1管理员,2用户"`
}
