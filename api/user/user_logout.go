package user

import (
	"blog/global"
	"blog/models/res"
	"blog/service/user_ser"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func (u User) UserLogout(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	token := c.Request.Header.Get("token")
	err := user_ser.Logout(claims, token)
	if err != nil {
		global.Log.Warn("注销失败", err.Error())
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)
}
