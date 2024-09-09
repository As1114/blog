package models

type MessageModel struct {
	MODEL
	SendUserID       uint      `gorm:"comment:发送人id" json:"send_user_id"`
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName string    `gorm:"comment:发送人昵称" json:"send_user_nick_name"`
	SendUserAvatar   string    `gorm:"comment:发送人头像" json:"send_user_avatar"`

	RecUserID       uint      `gorm:"comment:接收人id" json:"rec_user_id"`
	RecUserModel    UserModel `gorm:"foreignKey:RecUserID" json:"-"`
	RecUserNickName string    `gorm:"comment:接收人昵称" json:"rec_user_nick_name"`
	RecUserAvatar   string    `gorm:"comment:接收人头像" json:"rec_user_avatar"`

	IsRead  bool   `gorm:"default:false;comment:接收人是否查看" json:"is_read"`
	Content string `gorm:"comment:消息内容" json:"content"`
}
