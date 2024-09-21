package article

import (
	"github.com/axis1114/blog/models"
	"github.com/axis1114/blog/models/res"
	"github.com/gin-gonic/gin"
)

func (a Article) ArticleDelete(c *gin.Context) {
	var req models.ArticleSearchRequest
	err := c.ShouldBindUri(&req)
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
