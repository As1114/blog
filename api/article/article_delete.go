package article

import (
	"blog/models"
	"blog/models/res"
	"github.com/gin-gonic/gin"
)

type IDListRequest struct {
	ID string `json:"id"`
}

func (a Article) ArticleDelete(c *gin.Context) {
	var req IDListRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	err = models.DeleteDoc(req.ID)
	if err != nil {
		res.FailWithMessage("文章删除失败", c)
		return
	}
	res.OkWithMessage("文章删除成功", c)
}
