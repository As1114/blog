package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
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
		res.FailWithMessage("图片删除失败", c)
		return
	}
	err = global.DB.Delete(&imageList).Error
	if err != nil {
		res.FailWithMessage("图片删除失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("已删除 %d 张图片", count), c)

}
