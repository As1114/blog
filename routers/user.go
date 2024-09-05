package routers

import (
	"blog/api"
	"blog/middleware"
)

func (router RouterGroup) UserRouter() {
	userApi := api.AppGroupApp.UserApi
	userRouter := router.Group("user")
	userRouter.GET("info", middleware.JwtAuth(), userApi.Userinfo)
	userRouter.POST("create", userApi.UserCreate)
	userRouter.POST("login", userApi.UserLogin)
	userRouter.POST("logout", middleware.JwtAuth(), userApi.UserLogout)
}
