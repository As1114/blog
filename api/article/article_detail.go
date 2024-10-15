package article

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/service/search_ser"
)

func (a Article) ArticleDetail(c *gin.Context) {
	var req models.ArticleSearchRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	article, err := search_ser.GetDocumentById(req.ID)
	if err != nil {
		res.FailWithMessage("文章消失了", c)
		return
	}
	res.OkWithData(article, c)
}
