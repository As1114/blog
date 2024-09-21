package article

import (
	"github.com/axis1114/blog/models"
	"github.com/axis1114/blog/models/res"
	"github.com/axis1114/blog/service/search_ser"
	"github.com/gin-gonic/gin"
)

type ArticleListRequest struct {
	models.PageInfo
}

func (a Article) ArticleList(c *gin.Context) {
	var req ArticleListRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	articles := search_ser.SearchAllDocuments(req.PageInfo)
	res.OkWithList(articles, int64(len(articles)), c)
}
