package user_ser

import (
	"blog/global"
	"blog/models"
	"blog/models/ctypes"
	"blog/utils"
)

const Avatar = "/upload/avatar/default_avatar.jpg"

func (u UserService) CreateUser(nickname, account, password, ip string, role ctypes.Role) (err error) {
	//判断用户是否存在
	var user models.UserModel
	err = global.DB.Take(&user, "nickname=?", nickname).Error
	if err == nil {
		global.Log.Errorf("昵称%s已存在", nickname)
		return err
	}
	//对密码进行加密
	hashpwd := utils.HashPassword(password)

	addr := utils.GetAddrByIp(ip)

	err = global.DB.Create(&models.UserModel{
		Avatar:   Avatar,
		Nickname: nickname,
		Account:  account,
		Password: hashpwd,
		Role:     role,
		Address:  addr,
	}).Error
	if err != nil {
		global.Log.Error("create user err:", err)
		return err
	}
	return nil
}
