package routers

import "github.com/axis1114/blog/api"

func (router RouterGroup) CaptchaRouter() {
	captchaApi := api.AppGroupApp.CaptchaApi
	router.GET("captcha", captchaApi.CaptchaCreate)
}
