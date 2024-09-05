package routers

import "blog/api"

func (router RouterGroup) CaptchaRouter() {
	captchaApi := api.AppGroupApp.CaptchaApi
	router.GET("captcha", captchaApi.CaptchaCreate)
}
