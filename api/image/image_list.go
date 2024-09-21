package image

import (
	"github.com/axis1114/blog/models"
	"github.com/axis1114/blog/models/res"
	"github.com/axis1114/blog/service/search_ser"
	"github.com/gin-gonic/gin"
)

func (i Image) ImageList(c *gin.Context) {
	list, count, err := search_ser.SqlSearch(models.ImageModel{}, search_ser.Option{})
	if err != nil {
		res.FailWithMessage("图片加载失败", c)
		return
	}
	res.OkWithList(list, count, c)
}
