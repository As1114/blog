package image

import (
	"fmt"
	"github.com/axis1114/blog/global"
	"github.com/axis1114/blog/models"
	"github.com/axis1114/blog/models/res"
	"github.com/gin-gonic/gin"
)

func (i Image) ImageDelete(c *gin.Context) {
	var req models.RemoveRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	var imageList []models.ImageModel
	count := global.DB.Find(&imageList, req.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("图片不存在", c)
		return
	}
	err = global.DB.Delete(&imageList).Error
	if err != nil {
		res.FailWithMessage("图片删除失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("已删除 %d 张图片", count), c)

}
