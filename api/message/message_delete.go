package message

import (
	"github.com/axis1114/blog/global"
	"github.com/axis1114/blog/models"
	"github.com/axis1114/blog/models/res"
	"github.com/gin-gonic/gin"
)

type MessageDeleteRequest struct {
	ID uint `json:"id" uri:"id" form:"id"`
}

func (m Message) MessageDelete(c *gin.Context) {
	var req MessageDeleteRequest
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
