package models

import (
	"github.com/axis1114/blog/global"
)

type CommentModel struct {
	MODEL              `json:",select(c)"`
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments,select(c)"`
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"parent_comment_model"`
	ParentCommentID    *uint           `gorm:"comment:父评论id" json:"parent_comment_id,select(c)"`
	Content            string          `gorm:"comment:评论内容" json:"content,select(c)"`
	DiggCount          int             `gorm:"default:0;comment:点赞数" json:"digg_count,select(c)"`
	CommentCount       int             `gorm:"default:0;comment:子评论数" json:"comment_count,select(c)"`
	ArticleID          string          `gorm:"comment:文章id" json:"article_id,select(c)"`
	User               UserModel       `gorm:"foreignKey:UserID" json:"user,select(c)"`
	UserID             uint            `gorm:"comment:关联用户的id" json:"user_id,select(c)"`
}

func FindAllSubComment(com CommentModel) (subList []CommentModel) {
	global.DB.Preload("SubComments").Preload("User").Take(&com)
	for _, model := range com.SubComments {
		subList = append(subList, FindAllSubComment(*model)...)
		subList = append(subList, *model)
	}
	return subList
}

func GetCommentTree(rootComment *CommentModel) *CommentModel {
	global.DB.Preload("User").Preload("SubComments").Find(rootComment)

	for _, subComment := range rootComment.SubComments {
		GetCommentTree(subComment)
	}

	return rootComment
}

func FindArticleComment(articleID string) (RootCommentList []*CommentModel) {
	global.DB.Preload("User").Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)

	for _, model := range RootCommentList {
		GetCommentTree(model)
	}

	return RootCommentList
}
