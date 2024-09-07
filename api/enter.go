package api

import (
	"blog/api/article"
	"blog/api/captcha"
	"blog/api/image"
	"blog/api/user"
)

type AppGroup struct {
	CaptchaApi captcha.Captcha
	UserApi    user.User
	ImageApi   image.Image
	ArticleApi article.Article
}

var AppGroupApp = new(AppGroup)
