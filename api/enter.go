package api

import (
	"blog/api/captcha"
	"blog/api/user"
)

type AppGroup struct {
	CaptchaApi captcha.Captcha
	UserApi    user.User
}

var AppGroupApp = new(AppGroup)
