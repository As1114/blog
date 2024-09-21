package models

import "time"

type PageInfo struct {
	Page     int    `form:"page"`
	Key      string `form:"key"`
	PageSize int    `form:"page_size"`
	Sort     string `form:"sort"`
}

type MODEL struct {
	ID        uint      `gorm:"primaryKey;comment:id" json:"id,select($any)" structs:"-"`
	CreatedAt time.Time `gorm:"comment:创建时间" json:"created_at,select($any)" structs:"-"`
	UpdatedAt time.Time `gorm:"comment:更新时间" json:"-" structs:"-"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type ArticleSearchRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type MessageSearchRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}
