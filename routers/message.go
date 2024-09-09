package routers

import (
	"blog/api"
	"blog/middleware"
)

func (router RouterGroup) MessageRouter() {
	messageRouter := router.Group("message")
	messageApi := api.AppGroupApp.MessageApi
	messageRouter.POST("create", middleware.JwtAuth(), messageApi.MessageCreate)
	messageRouter.GET("list/:id", middleware.JwtAuth(), messageApi.MessageList)
	messageRouter.DELETE("delete/:id", middleware.JwtAdmin(), messageApi.MessageDelete)
}
