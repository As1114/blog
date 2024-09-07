package user

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"blog/utils"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		res.FailWithMessage("用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(newMap).Error
	if err != nil {
		global.Log.Error("fail to update user info", zap.Error(err))
		res.FailWithMessage("修改信息失败", c)
		return
	}
	res.OkWithMessage("修改信息成功", c)
}
