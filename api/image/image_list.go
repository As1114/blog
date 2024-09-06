package image

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"blog/service/search_ser"
	"github.com/gin-gonic/gin"
)

func (i Image) ImageList(c *gin.Context) {
	list, count, err := search_ser.SqlSearch(models.ImageModel{}, search_ser.Option{})
	if err != nil {
		global.Log.Error("select image list error:", err.Error())
		res.FailWithMessage("图片加载失败", c)
		return
	}
	res.OkWithList(list, count, c)
}
