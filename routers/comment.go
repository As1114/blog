package routers

import "blog/api"

func (router RouterGroup) CommentRouter() {
	commentRouter := router.Group("/comment")
	commentApi := api.AppGroupApp.CommentApi
}
