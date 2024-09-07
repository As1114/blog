package article

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
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
			res.FailWithMessage("图片不存在", c)
			return
		}
	}
	article := models.Article{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     req.Title,
		Abstract:  req.Abstract,
		Content:   req.Content,
		Category:  req.Category,
		CoverID:   req.CoverID,
		CoverURL:  coverUrl,
	}
	maps := structs.Map(&article)
	var DataMap = map[string]any{}

	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		DataMap[key] = v
	}

	err = article.UpdateDocument(req.ID, DataMap)
	if err != nil {
		res.FailWithMessage("文章更新失败", c)
		return
	}
	res.OkWithMessage("文章更新成功", c)
}
