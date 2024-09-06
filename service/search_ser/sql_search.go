package search_ser

import (
	"blog/global"
	"blog/models"
	"fmt"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo          // 分页查询
	Likes           []string // 需要模糊匹配的字段列表
	Debug           bool     // 是否打印sql
	Where           *gorm.DB // 额外的查询
	Preload         []string // 预加载的字段列表
}

func SqlSearch[T any](model T, option Option) (list []T, count int64, err error) {
	query := global.DB.Model(&model).Where(model)
	if option.Debug {
		query = query.Debug()
	}

	// 设置排序
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}

	// 设置分页
	if option.Limit == 0 {
		option.Limit = 10
	}
	offset := (option.Page - 1) * option.Limit

	// 高级查询
	if option.Where != nil {
		query = query.Where(option.Where)
	}

	// 模糊匹配
	if option.Key != "" {
		likeQuery := global.DB
		for index, column := range option.Likes {
			//SELECT * FROM users WHERE name LIKE '%john%' OR email LIKE '%john%' OR bio LIKE '%john%'
			if index == 0 {
				likeQuery = likeQuery.Where(fmt.Sprintf("%s LIKE ?", column), fmt.Sprintf("%%%s%%", option.Key))
			} else {
				likeQuery = likeQuery.Or(fmt.Sprintf("%s LIKE ?", column), fmt.Sprintf("%%%s%%", option.Key))
			}
		}
		query = query.Where(likeQuery)
	}

	// 预加载
	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}

	// 获取总数
	countQuery := query.Session(&gorm.Session{NewDB: true})
	countQuery.Model(&model).Count(&count)

	// 获取列表
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return
}
