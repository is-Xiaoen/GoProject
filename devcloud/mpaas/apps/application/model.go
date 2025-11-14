package application

import (
	"bytes"
	"time"

	"github.com/google/uuid"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/policy"
)

func NewApplication(req CreateApplicationRequest) (*Application, error) {
	app := &Application{
		CreateApplicationRequest: req,
	}

	if err := app.Validate(); err != nil {
		return nil, err
	}

	// 动态计算评审状态
	if len(req.Audits) == 0 {
		app.SetReady(true)
	} else {
		app.SetReady(false)
	}

	app.BuildId()
	return app, nil
}

// Application 应用
type Application struct {
	// 对象Id
	Id string `json:"id" bson:"_id" gorm:"column:id;primary_key"`
	// 更新时间
	UpdateAt *time.Time `json:"update_at" bson:"update_at" gorm:"column:update_at"`
	// 更新人
	UpdateBy string `json:"update_by" bson:"update_by" gorm:"column:update_by"`
	// 创建请求
	CreateApplicationRequest
	// 应用状态
	ApplicationStatus
}

func (a *Application) TableName() string {
	return "applications"
}

func (a *Application) String() string {
	return pretty.ToJSON(a)
}

func (a *Application) SetReady(v bool) *Application {
	a.Ready = &v
	return a
}

func (a *Application) BuildId() {
	bf := bytes.NewBuffer([]byte{})

	switch a.Type {
	case TYPE_SOURCE_CODE:
		bf.WriteString(a.CodeRepository.SshUrl)
	case TYPE_CONTAINER_IMAGE:
		bf.WriteString(a.GetImageRepositoryPrimaryAddress())
	}

	a.Id = uuid.NewSHA1(uuid.Nil, bf.Bytes()).String()
}

func NewCreateApplicationRequest() *CreateApplicationRequest {
	return &CreateApplicationRequest{
		CreateAt: time.Now(),
		CreateApplicationSpec: CreateApplicationSpec{
			Extras:          map[string]string{},
			ImageRepository: []ImageRepository{},
		},
	}
}

type CreateApplicationRequest struct {
	// 创建人
	CreateBy string `json:"create_by" bson:"create_by" gorm:"column:create_by" description:"创建人"`
	// 创建时间
	CreateAt time.Time `json:"create_at" bson:"create_at" gorm:"column:create_at" description:"创建时间"`
	// 资源范围, Namespace是继承的, Scope是API添加的
	policy.ResourceLabel
	// 应用创建参数
	CreateApplicationSpec
}

func (a *CreateApplicationRequest) Validate() error {
	return validator.Validate(a)
}

func (a *CreateApplicationRequest) GetImageRepositoryPrimaryAddress() string {
	for _, repo := range a.ImageRepository {
		if repo.IsPrimary {
			return repo.Address
		}
	}

	return ""
}

type CreateApplicationSpec struct {
	// 应用名称
	Name string `json:"name" bson:"name" gorm:"column:name" description:"应用名称"`
	// 应用描述
	Description string `json:"description" bson:"description" gorm:"column:description" description:"应用描述"`
	// 应用图标
	Icon string `json:"icon" bson:"icon" gorm:"column:icons" description:"应用图标"`
	// 应用类型
	Type TYPE `json:"type" bson:"type" gorm:"column:type" description:"应用类型, SOURCE_CODE, CONTAINER_IMAGE, OTHER"`
	// 应用代码仓库信息
	CodeRepository CodeRepository `json:"code_repository" bson:",inline" gorm:"embedded" description:"应用代码仓库信息"`
	// 应用镜像仓库信息
	ImageRepository []ImageRepository `json:"image_repository" gorm:"column:image_repository;type:json;serializer:json;" bson:"image_repository"  description:"应用镜像仓库信息"`
	// 应用所有者
	Owner string `json:"owner" bson:"owner" gorm:"column:owner" description:"应用所有者"`
	// 应用等级, 评估这个应用的重要程度
	Level *uint32 `json:"level" bson:"level" gorm:"column:level" description:"应用等级, 评估这个应用的重要程度"`
	// 应用优先级, 应用启动的先后顺序
	Priority *uint32 `json:"priority" bson:"priority" gorm:"column:priority" description:"应用优先级, 应用启动的先后顺序"`
	// 额外的其他属性
	Extras map[string]string `json:"extras" form:"extras" bson:"extras" gorm:"column:extras;type:json;serializer:json;"`

	// 指定应用的评审方
	Audits []ApplicationReadyAudit `json:"audits" bson:"audits" gorm:"column:audits;type:json;serializer:json" description:"参与应用准备就绪的评审方"`
}

// CodeRepository 服务代码仓库信息
type CodeRepository struct {
	// 仓库提供商
	Provider SCM_PROVIDER `json:"provider" bson:"provider" gorm:"column:provider"`
	// token 操作仓库, 比如设置回调
	Token string `json:"token" bson:"token" gorm:"column:token"`
	// 仓库对应的项目Id
	ProjectId string `json:"project_id" bson:"project_id" gorm:"column:project_id"`
	// 仓库对应空间
	Namespace string `json:"namespace" bson:"namespace" gorm:"column:namespace"`
	// 仓库web url地址
	WebUrl string `json:"web_url" bson:"web_url" gorm:"column:web_url"`
	// 仓库ssh url地址
	SshUrl string `json:"ssh_url" bson:"ssh_url" gorm:"column:ssh_url"`
	// 仓库http url地址
	HttpUrl string `json:"http_url" bson:"http_url" gorm:"column:http_url"`
	// 源代码使用的编程语言, 构建时, 不同语言有不同的构建环境
	Language *LANGUAGE `json:"language" bson:"language" gorm:"column:language"`
	// 开启Hook设置
	EnableHook bool `json:"enable_hook" bson:"enable_hook" gorm:"column:enable_hook"`
	// Hook设置
	HookConfig string `json:"hook_config" bson:"hook_config" gorm:"column:hook_config"`
	// scm设置Hook后返回的id, 用于删除应用时，取消hook使用
	HookId string `json:"hook_id" bson:"hook_id" gorm:"column:hook_id"`
	// 仓库的创建时间
	CreatedAt *time.Time `json:"created_at" bson:"created_at" gorm:"column:created_at"`
}

// ImageRepository 镜像仓库
type ImageRepository struct {
	// 服务镜像地址, 比如 gcr.lank8s.cn/kaniko-project/executor
	Address string `json:"address" bson:"address"`
	// 是不是主仓库
	IsPrimary bool `json:"is_primary" bson:"is_primary"`
}

type ApplicationStatus struct {
	// 该应用是否已经准备就绪，多方确认的一个过程后计算出来的
	Ready *bool `json:"ready" bson:"ready" gorm:"column:ready" description:"该应用是否已经准备就绪"`
	// 就绪状态修改时间
	UpdateAt *time.Time `json:"ready_update_at" bson:"ready_update_at" gorm:"column:ready_update_at" description:"就绪状态修改时间"`
}

// ApplicationReadyAudit 参与应用准备就绪的评审方
type ApplicationReadyAudit struct {
	// 评审角色, Dev, Test, Ops
	RoleName string `json:"role_name"`
	// 评审人
	AuditBy string `json:"audit_by"`
	// 评审时间
	AuditAt time.Time `json:"audit_at"`
	// 是否就绪
	Ready bool `json:"ready"`
	// 评审建议
	Message string `json:"message"`
}
