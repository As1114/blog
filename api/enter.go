package api

import (
	"blog/api/article"
	"blog/api/captcha"
	"blog/api/comment"
	"blog/api/image"
	"blog/api/user"
)

type AppGroup struct {
	CaptchaApi captcha.Captcha
	UserApi    user.User
	ImageApi   image.Image
	ArticleApi article.Article
	CommentApi comment.Comment
}

var AppGroupApp = new(AppGroup)
