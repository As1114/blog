package routers

import "blog/api"

func (router RouterGroup) UserRouter() {
	userApi := api.AppGroupApp.UserApi
	userRouter := router.Group("user")
	userRouter.GET("info", userApi.Userinfo)
	userRouter.POST("create", userApi.UserCreate)
	userRouter.POST("login", userApi.UserLogin)
	userRouter.POST("logout", userApi.UserLogout)
}
