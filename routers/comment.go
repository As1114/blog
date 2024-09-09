package routers

import "blog/api"

func (router RouterGroup) CommentRouter() {
	commentRouter := router.Group("comment")
	commentApi := api.AppGroupApp.CommentApi
	commentRouter.POST("create", commentApi.CommentCreate)
	commentRouter.DELETE("delete/:id", commentApi.CommentDelete)
	commentRouter.GET("list/:id", commentApi.CommentList)
}
