package article

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type ArticleRequest struct {
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	Category string `json:"category"`
	Content  string `json:"content" `
	CoverID  uint   `json:"cover_id"`
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
	userId := claims.UserID
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

	if req.CoverID == 0 {
		var imageIDList []uint
		global.DB.Model(models.ImageModel{}).Select("id").Scan(&imageIDList)
		if len(imageIDList) == 0 {
			res.FailWithMessage("暂无图片", c)
			return
		}
		rand.New(rand.NewSource(time.Now().UnixNano()))
		req.CoverID = imageIDList[rand.Intn(len(imageIDList))]
	}

	var coverUrl string
	err = global.DB.Model(models.ImageModel{}).Where("id = ?", req.CoverID).Select("path").Scan(&coverUrl).Error
	if err != nil {
		res.FailWithMessage("图片不存在", c)
		return
	}
	var user models.UserModel
	err = global.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	article := models.Article{
		Title:      req.Title,
		Abstract:   req.Abstract,
		Category:   req.Category,
		Content:    content,
		CoverID:    req.CoverID,
		CoverURL:   coverUrl,
		UserID:     userId,
		CreatedAt:  now,
		UpdatedAt:  now,
		UserName:   user.Nickname,
		UserAvatar: user.Avatar,
	}
	articleItem := models.ArticleItem{
		Article: article,
	}
	exist := articleItem.DocumentExist(req.Title)
	if exist {
		res.FailWithMessage("文章已存在", c)
		return
	}
	articleItem.CreateDocument()
	res.OkWithMessage("文章发布成功", c)
}
