package article

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"time"
)

type ArticleUpdateRequest struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	Content  string `json:"content"`
	Category string `json:"category"`
	CoverID  uint   `json:"cover_id"`
}

func (a Article) ArticleUpdate(c *gin.Context) {
	var req ArticleUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	var coverUrl string
	if req.CoverID != 0 {
		err = global.DB.Model(models.ImageModel{}).Where("id = ?", req.CoverID).Select("path").Scan(&coverUrl).Error
		if err != nil {
			res.FailWithMessage("文章更新失败", c)
			return
		}
	}
	article := models.Article{
		ID:        req.ID,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     req.Title,
		Abstract:  req.Abstract,
		Content:   req.Content,
		Category:  req.Category,
		CoverID:   req.CoverID,
		CoverURL:  coverUrl,
	}
	err = article.UpdateDoc()
	if err != nil {
		res.FailWithMessage("文章更新失败", c)
		return
	}
	res.OkWithMessage("文章更新成功", c)
}
