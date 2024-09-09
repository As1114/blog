package article

import (
	"blog/models"
	"blog/models/res"
	"blog/service/search_ser"
	"github.com/gin-gonic/gin"
)

type ArticleSearchRequest struct {
	models.PageInfo
}

func (a Article) ArticleList(c *gin.Context) {
	var req ArticleSearchRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	fields := []string{"category"}
	articles := search_ser.SearchDocumentMultiMatch(fields, req.Key, req.PageInfo)
	res.OkWithList(articles, int64(len(articles)), c)
}
