package models

type MessageModel struct {
	MODEL
	SendUserID       uint      `gorm:"comment:发送人id" json:"send_user_id"`
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName string    `gorm:"comment:发送人昵称" json:"send_user_nick_name"`
	SendUserAvatar   string    `gorm:"comment:发送人头像" json:"send_user_avatar"`

	RevUserID       uint      `gorm:"comment:接收人id" json:"rev_user_id"`
	RevUserModel    UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	RevUserNickName string    `gorm:"comment:接收人昵称" json:"rev_user_nick_name"`
	RevUserAvatar   string    `gorm:"comment:接收人头像" json:"rev_user_avatar"`

	IsRead  bool   `gorm:"default:false;comment:接收人是否查看" json:"is_read"`
	Content string `gorm:"comment:消息内容" json:"content"`
}
