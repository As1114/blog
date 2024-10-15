package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/service/user_ser"
	"github.com/nsxz1114/blog/utils"
)

func (u User) UserLogout(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	token := c.Request.Header.Get("token")
	err := user_ser.Logout(claims, token)
	if err != nil {
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)
}
