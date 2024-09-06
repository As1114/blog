package article

import (
	"blog/global"
	"blog/models/res"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

type ArticleRequest struct {
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	Category string `json:"category"`
	Content  string `json:"content" `
	CoverUrl string `json:"cover_url"`
}

func (a Article) ArticleCreate(c *gin.Context) {
	var req ArticleRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	user_id := claims.UserID
	html, err := utils.ConvertMarkdownToHTML(req.Content)
	if err != nil {
		global.Log.Error("ConvertMarkdownToHTML err:", err)
		return
	}
	content, err := utils.ConvertHTMLToMarkdown(html)
	if err != nil {
		global.Log.Error("ConvertHTMLToMarkdown err:", err)
		return
	}
}
