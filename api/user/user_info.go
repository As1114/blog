package user

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func (u User) Userinfo(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)

	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		global.Log.Warn("用户不存在", err.Error())
		res.FailWithMessage("用户不存在", c)
		return
	}
	res.OkWithData(user, c)
}
