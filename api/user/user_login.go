package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/api/captcha"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/utils"
)

type UserLoginRequest struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captcha_id"`
}

func (u User) UserLogin(c *gin.Context) {
	var req UserLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	if req.Captcha != "" && req.CaptchaId != "" && captcha.Store.Verify(req.CaptchaId, req.Captcha, true) {
		var user models.UserModel
		err = global.DB.Take(&user, "account=?", req.Account).Error
		if err != nil {
			res.FailWithMessage("用户名或密码错误", c)
			return
		}
		check := utils.CheckPassword(user.Password, req.Password)
		if !check {
			res.FailWithMessage("用户名或密码错误", c)
			return
		}
		token, err := utils.GetToken(utils.PayLoad{
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
