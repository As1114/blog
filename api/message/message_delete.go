package message

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
)

func (m Message) MessageDelete(c *gin.Context) {
	var req models.MessageSearchRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	var message models.MessageModel
	err = global.DB.Where("id=?", req.ID).First(&message).Error
	if err != nil {
		res.FailWithMessage("删除失败", c)
		return
	}
	global.DB.Model(&message).Delete("id=?", req.ID)
	res.OkWithMessage("删除成功", c)
}
