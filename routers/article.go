package routers

import (
	"blog/api"
	"blog/middleware"
)

func (router RouterGroup) ArticleRouter() {
	articleApi := api.AppGroupApp.ArticleApi
	articleRouter := router.Group("article")
	articleRouter.POST("create", middleware.JwtAdmin(), articleApi.ArticleCreate)
	articleRouter.GET("list", articleApi.ArticleList)
	articleRouter.DELETE("delete", middleware.JwtAdmin(), articleApi.ArticleDelete)
}
