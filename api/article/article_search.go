package article

import (
	"github.com/axis1114/blog/models/res"
	"github.com/axis1114/blog/service/search_ser"
	"github.com/gin-gonic/gin"
)

type ArticleSearchRequest struct {
	Key string `json:"key" form:"key"`
}

func (a Article) ArticleSearch(c *gin.Context) {
	var req ArticleSearchRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	fields := []string{"title", "id"}
	articles := search_ser.SearchDocumentMultiMatchByTitle(fields, req.Key)
	res.OkWithList(articles, int64(len(articles)), c)
}
