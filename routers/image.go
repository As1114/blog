package routers

import (
	"github.com/axis1114/blog/api"
	"github.com/axis1114/blog/middleware"
)

func (router RouterGroup) ImageRouter() {
	imageApi := api.AppGroupApp.ImageApi
	imageRouter := router.Group("image")
	imageRouter.POST("upload", middleware.JwtAdmin(), imageApi.ImageUpload)
	imageRouter.GET("list", middleware.JwtAdmin(), imageApi.ImageList)
	imageRouter.DELETE("delete", middleware.JwtAdmin(), imageApi.ImageDelete)
}
