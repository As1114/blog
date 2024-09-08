package models

type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"`
	ParentCommentID    *uint           `gorm:"comment:父评论id" json:"parent_comment_id"`
	Content            string          `gorm:"comment:评论内容" json:"content"`
	DiggCount          int             `gorm:"default:0;comment:点赞数" json:"digg_count"`
	CommentCount       int             `gorm:"default:0;comment:子评论数" json:"comment_count"`
	ArticleID          string          `gorm:"comment:文章id" json:"article_id"`
	User               UserModel       `gorm:"comment:关联的用户" json:"user"`
	UserID             uint            `gorm:"comment:关联用户的id" json:"user_id"`
}
