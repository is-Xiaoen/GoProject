package namespace

import (
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/modules/iam/apps"
)

func NewNamespace() *Namespace {
	return &Namespace{
		ResourceMeta: *apps.NewResourceMeta(),
	}
}

type Namespace struct {
	// 基础数据
	apps.ResourceMeta
	// 空间属性
	CreateNamespaceRequest
}

func (n *Namespace) IsOwner(ownerUserId uint64) bool {
	return n.OwnerUserId == ownerUserId
}

func (n *Namespace) TableName() string {
	return "namespaces"
}

func (n *Namespace) String() string {
	return pretty.ToJSON(n)
}

func NewCreateNamespaceRequest() *CreateNamespaceRequest {
	return &CreateNamespaceRequest{
		Extras:  map[string]string{},
		Enabled: true,
	}
}

type CreateNamespaceRequest struct {
	// 父Namespace Id
	ParentId uint64 `json:"parent_id" bson:"parent_id" gorm:"column:parent_id;type:uint;index" description:"父Namespace Id"`
	// 全局唯一
	Name string `json:"name" bson:"name" validate:"required" gorm:"column:name;type:varchar(200);not null;uniqueIndex" description:"空间名称" unique:"true"`
	// 空间负责人
	OwnerUserId uint64 `json:"owner_user_id" bson:"owner_user_id" gorm:"column:owner_user_id;type:uint;index;not null" description:" 空间负责人Id"`
	// 禁用项目, 该项目所有人暂时都无法访问
	Enabled bool `json:"enabled" bson:"enabled" gorm:"column:enabled;type:tinyint(1)" description:"是否启用"`
	// 空间描述图片
	Icon string `json:"icon" bson:"icon" gorm:"column:icon;type:varchar(200)" description:"空间图标"`
	// 空间描述
	Description string `json:"description" bson:"description" gorm:"column:description;type:text" description:"空间描述"`
	// 标签
	Label string `json:"label" gorm:"column:label;type:varchar(200);index" description:"标签"`
	// 扩展信息
	Extras map[string]string `json:"extras" bson:"extras" gorm:"column:extras;serializer:json;type:json" description:"扩展信息"`
}

func (r *CreateNamespaceRequest) Validate() error {
	return validator.Validate(r)
}
