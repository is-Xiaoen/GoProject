package models

import "github.com/infraboard/mcube/v2/tools/pretty"

type BookSet struct {
	// 总共多少个
	Total int64 `json:"total"`
	// book清单
	Items []*Book `json:"items"`
}

type Book struct {
	// 对象Id
	Id uint `json:"id" gorm:"primaryKey;column:id"`

	BookSpec
}

func (b *Book) String() string {
	return pretty.ToJSON(b)
}

type BookSpec struct {
	// type 用于要使用gorm 来自动创建和更新表的时候 才需要定义
	Title  string  `json:"title"  gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author"  gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price"  gorm:"column:price" validate:"required"`
	// bool false
	// nil 是零值, false
	IsSale *bool `json:"is_sale"  gorm:"column:is_sale"`
}

// books
func (b *Book) TableName() string {
	return "books"
}
