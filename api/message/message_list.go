package message

import (
	"fmt"
	"github.com/axis1114/blog/global"
	"github.com/axis1114/blog/models"
	"github.com/axis1114/blog/models/res"
	"github.com/axis1114/blog/utils"
	"github.com/gin-gonic/gin"
)

type MessageListRequest struct {
	ID uint `json:"id" uri:"id" form:"id"`
}

func (m Message) MessageList(c *gin.Context) {
	var req MessageListRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	fmt.Println(claims.UserID, req.ID)
	var messageList []models.MessageModel
	err = global.DB.Order("created_at asc").
		Find(&messageList, "(send_user_id = ? and rec_user_id = ?) or (rec_user_id = ? and send_user_id = ?)",
			claims.UserID, req.ID, claims.UserID, req.ID).Error
	if err != nil {
		res.FailWithMessage("暂无消息", c)
		return
	}
	res.OkWithList(messageList, int64(len(messageList)), c)
}
