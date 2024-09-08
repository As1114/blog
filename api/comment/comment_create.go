package comment

import (
	"blog/models/res"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

type CommentCreateRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"`
}

func (comment *Comment) CommentCreate(c *gin.Context) {
	var req CommentCreateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)

}
