package comment

import (
	"blog/models"
	"blog/models/res"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `form:"id" uri:"id" json:"id"`
}

func (comment Comment) CommentList(c *gin.Context) {
	var req CommentListRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	commentList := models.FindArticleComment(req.ArticleID)
	data := filter.Select("c", commentList)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list := make([]models.CommentModel, 0)
		res.OkWithList(list, 0, c)
		return
	}
	res.OkWithList(data, int64(len(commentList)), c)
}
