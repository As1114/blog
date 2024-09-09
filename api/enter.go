package api

import (
	"blog/api/article"
	"blog/api/captcha"
	"blog/api/comment"
	"blog/api/image"
	"blog/api/message"
	"blog/api/user"
)

type AppGroup struct {
	CaptchaApi captcha.Captcha
	UserApi    user.User
	ImageApi   image.Image
	ArticleApi article.Article
	CommentApi comment.Comment
	MessageApi message.Message
}

var AppGroupApp = new(AppGroup)
