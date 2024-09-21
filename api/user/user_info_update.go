package user

import (
	"github.com/axis1114/blog/global"
	"github.com/axis1114/blog/models"
	"github.com/axis1114/blog/models/res"
	"github.com/axis1114/blog/utils"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserInfoUpdateRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Avatar   string `json:"avatar" structs:"avatar"`
	Email    string `json:"email"  structs:"email"`
}

func (u User) UserInfoUpdate(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	var req UserInfoUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	var newMap = map[string]interface{}{}
	maps := structs.Map(req)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMap[key] = val
		}
	}
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("更新失败", c)
		return
	}
	err = global.DB.Model(&user).Updates(newMap).Error
	if err != nil {
		res.FailWithMessage("更新失败", c)
		return
	}
	res.OkWithMessage("更新成功", c)
}
