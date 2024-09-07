package article

import (
	"blog/models"
	"blog/models/res"
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
	var article models.Article
	fields := []string{"category"}
	articles := article.SearchDocumentMultiMatch(fields, req.Key, req.PageInfo)
	res.OkWithList(articles, int64(len(articles)), c)
}
