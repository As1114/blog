package comment

import (
	"github.com/axis1114/blog/global"
	"github.com/axis1114/blog/models"
	"github.com/axis1114/blog/models/res"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

func (comment Comment) CommentDelete(c *gin.Context) {
	var req CommentIDRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, req.ID).Error
	if err != nil {
		res.FailWithMessage("评论不存在", c)
	}

	subCommentList := models.FindAllSubComment(commentModel)
	count := len(subCommentList) + 1

	if commentModel.ParentCommentID != nil {
		global.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}

	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(models.CommentModel{}).Delete("id = ?", id)
	}

	res.OkWithMessage("评论删除成功", c)
}
