package user

import (
	"github.com/axis1114/blog/global"
	"github.com/axis1114/blog/models"
	"github.com/axis1114/blog/models/res"
	"github.com/axis1114/blog/utils"
	"github.com/gin-gonic/gin"
)

func (u User) Userinfo(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)

	var user models.UserModel
	if err := global.DB.First(&user, claims.UserID).Error; err != nil {
		global.Log.Warn("用户不存在", err.Error())
		res.FailWithMessage("用户不存在", c)
		return
	}

	res.OkWithData(user, c)
}
