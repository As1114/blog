package user

import (
	"fmt"
	"github.com/axis1114/blog/global"
	"github.com/axis1114/blog/models/ctype"
	"github.com/axis1114/blog/models/res"
	"github.com/axis1114/blog/service/user_ser"
	"github.com/axis1114/blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserCreateRequest struct {
	Nickname string     `json:"nick_name" binding:"required" msg:"请输入昵称"`
	Password string     `json:"password" binding:"required" msg:"请输入密码"`
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择身份"`
}

func (u User) UserCreate(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	account := utils.GenerateID()
	err := user_ser.CreateUser(req.Nickname, strconv.FormatInt(account, 10), req.Password, c.ClientIP(), req.Role)
	if err != nil {
		global.Log.Error("create user error:", err.Error())
		res.FailWithError(err, &req, c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功", req.Nickname), c)
	return
}
