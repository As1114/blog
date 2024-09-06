package models

import "blog/models/ctype"

type ImageModel struct {
	MODEL
	Path      string          `gorm:"comment:图片路径" json:"path"`
	Hash      string          `gorm:"comment:图片的hash值" json:"hash"`
	Name      string          `gorm:"comment:图片名称" json:"name"`
	ImageType ctype.ImageType `gorm:"default:1;comment:图片的类型，1本地，2云端，默认是1" json:"image_type"`
}
