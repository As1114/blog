package routers

import "blog/api"

func (router RouterGroup) MessageRouter() {
	messageRouter := router.Group("message")
	messageApi := api.AppGroupApp.MessageApi
	messageRouter.POST("create", messageApi.MessageCreate)
	messageRouter.GET("list/:id", messageApi.MessageList)
	messageRouter.DELETE("delete/:id", messageApi.MessageDelete)
}
