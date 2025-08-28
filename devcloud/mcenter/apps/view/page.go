package view

import (
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/modules/iam/apps"
)

func NewPage() *Page {
	return &Page{
		ResourceMeta: *apps.NewResourceMeta(),
	}
}

type Page struct {
	// 基础数据
	apps.ResourceMeta
	// 菜单定义
	CreatePageRequest
	// 用户是否有权限访问该页面, 只有在策略模块查询时，才会计算出该字段
	HasPermission *bool `json:"has_permission,omitempty" gorm:"column:has_permission;type:tinyint(1)" optional:"true" description:"用户是否有权限访问该页面"`
}

func (p *Page) TableName() string {
	return "pages"
}

func NewCreatePageRequest() *CreatePageRequest {
	return &CreatePageRequest{
		Extras: map[string]string{},
	}
}

type CreatePageRequest struct {
	// 菜单Id
	MenuId uint64 `json:"menu_id" bson:"menu_id" gorm:"column:menu_id;type:uint;index" description:"菜单Id"`
	// 页面路径
	Path string `json:"path" bson:"path" gorm:"column:path" description:"页面路径" unique:"true"`
	// 页面名称
	Name string `json:"name" bson:"name" gorm:"column:name" description:"页面名称"`
	// 标签
	Label string `json:"label" gorm:"column:label;type:varchar(200);index" description:"标签" optional:"true"`
	// 页面组件，比如按钮
	Components []Component `json:"components" gorm:"column:components;type:json;serializer:json" description:"组件" optional:"true"`
	// 其他扩展信息
	Extras map[string]string `json:"extras" gorm:"column:extras;serializer:json;type:json" description:"其他扩展信息" optional:"true"`
}

func (r *CreatePageRequest) Validate() error {
	return validator.Validate(r)
}

// 组件
type Component struct {
	// 组件名称
	Name string `json:"name" bson:"name" description:"组件名称"`
	// 组件说明
	Description string `json:"description" optional:"true" description:"组件说明"`
	// 组件使用文档链接
	RefDocURL string `json:"ref_doc_url" optional:"true" description:"组件使用文档链接"`
	// 关联的Api接口
	RefEndpointId []uint64 `json:"ref_endpoints" description:"该页面管理的Api接口关联的接口" optional:"true"`
	// 其他扩展信息
	Extras map[string]string `json:"extras" description:"其他扩展信息" optional:"true"`
}
