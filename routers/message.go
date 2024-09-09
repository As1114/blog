package routers

import "blog/api"

func (router RouterGroup) MessageRouter() {
	messageRouter := router.Group("message")
	messageApi := api.AppGroupApp.MessageApi
}
