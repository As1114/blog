package api

import "blog/api/captcha"

type AppGroup struct {
	CaptchaApi captcha.Captcha
}

var AppGroupApp = new(AppGroup)
