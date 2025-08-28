package view

import (
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/modules/iam/apps"
)

func NewMenu() *Menu {
	return &Menu{
		ResourceMeta: *apps.NewResourceMeta(),
		Pages:        []*Page{},
	}
}

type Menu struct {
	// 基础数据
	apps.ResourceMeta
	// 菜单定义
	CreateMenuRequest
	// 用户是否有权限访问该菜单, 只有在策略模块查询时，才会计算出该字段
	HasPermission *bool `json:"has_permission,omitempty" gorm:"column:has_permission;type:tinyint(1)" optional:"true" description:"用户是否有权限访问该菜单"`
	// 菜单关联的页面
	Pages []*Page `json:"pages,omitempty" gorm:"-" description:"菜单关联的页面"`
}

func (m *Menu) SetHasPermission(v bool) *Menu {
	m.HasPermission = &v
	return m
}

func (m *Menu) TableName() string {
	return "menus"
}

func NewCreateMenuRequest() *CreateMenuRequest {
	return &CreateMenuRequest{
		Extras: map[string]string{},
	}
}

type CreateMenuRequest struct {
	// 服务
	Service string `json:"service" gorm:"column:service;type:varchar(100);index" bson:"service" description:"服务名称"`
	// 父Menu Id
	ParentId uint64 `json:"parent_id" bson:"parent_id" gorm:"column:parent_id;type:uint;index" description:"父Menu Id" optional:"true"`
	// 菜单路径
	Path string `json:"path" bson:"path" gorm:"column:path" description:"菜单路径" unique:"true"`
	// 菜单名称
	Name string `json:"name" bson:"name" gorm:"column:name" description:"菜单名称"`
	// 图标
	Icon string `json:"icon" bson:"icon" gorm:"column:icon" description:"图标" optional:"true"`
	// 标签
	Label string `json:"label" gorm:"column:label;type:varchar(200);index" description:"标签" optional:"true"`
	// 其他扩展信息
	Extras map[string]string `json:"extras" gorm:"column:extras;serializer:json;type:json" description:"其他扩展信息" optional:"true"`
}

func (r *CreateMenuRequest) Validate() error {
	return validator.Validate(r)
}
