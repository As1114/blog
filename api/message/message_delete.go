package message

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"github.com/gin-gonic/gin"
)

type MessageDeleteRequest struct {
	ID string `json:"id" url:"id" form:"id"`
}

func (m Message) MessageDelete(c *gin.Context) {
	var req MessageDeleteRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}

	err = global.DB.Take(&models.MessageModel{}, req.ID).Error
	if err == nil {
		res.FailWithMessage("删除失败", c)
		return
	}
	global.DB.Delete(&models.MessageModel{})
	res.OkWithMessage("删除成功", c)
}
