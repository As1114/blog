package user

import (
	"blog/api/captcha"
	"blog/global"
	"blog/models"
	"blog/models/res"
	"blog/utils/jwt"
	"blog/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

func (u User) UserLogin(c *gin.Context) {
	var req UserLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
	}

	if req.Captcha != "" && req.CaptchaId != "" && captcha.Store.Verify(req.CaptchaId, req.Captcha, true) {
		var user models.UserModel
		err = global.DB.Take(&user, "account=?", req.Account).Error
		if err != nil {
			global.Log.Warn("用户名或密码错误", err.Error())
			res.FailWithMessage("用户名或密码错误", c)
			return
		}
		check := pwd.CheckPassword(user.Password, req.Password)
		if !check {
			global.Log.Warn("用户名或密码错误")
			res.FailWithMessage("用户名或密码错误", c)
			return
		}
		token, err := jwt.GetToken(jwt.PayLoad{
			Account: req.Account,
			UserID:  user.ID,
			Role:    int(user.Role)})
		if err != nil {
			global.Log.Error("token生成失败", err.Error())
			res.FailWithMessage("登录失败", c)
			return
		}
		c.Request.Header.Set("token", token)
		res.OkWithData(token, c)
		return
	}
	res.FailWithMessage("验证码错误", c)
}
