package api

import (
	"github.com/axis1114/blog/api/article"
	"github.com/axis1114/blog/api/captcha"
	"github.com/axis1114/blog/api/comment"
	"github.com/axis1114/blog/api/image"
	"github.com/axis1114/blog/api/message"
	"github.com/axis1114/blog/api/user"
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
