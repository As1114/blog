package user_ser

import (
	"blog/global"
	"blog/models"
	"blog/models/ctypes"
	"blog/utils/pwd"
)

const Avatar = "/upload/avatar/default_avatar.jpg"

func (u UserService) CreateUser(nickname, account, password, ip string, role ctypes.Role) {
	//判断用户是否存在
	var user models.UserModel
	err := global.DB.Take(&user, "nickname=?", nickname).Error
	if err != nil {
		global.Log.Errorf("昵称%s已存在", nickname)
		return
	}
	//对密码进行加密
	hashpwd := pwd.HashPassword(password)
}
