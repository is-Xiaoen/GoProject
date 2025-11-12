package application

import (
	"time"

	"github.com/infraboard/mcube/v2/tools/pretty"
)

// Application 应用
type Application struct {
	// 对象Id
	Id string `json:"id" bson:"_id"`
	// 更新时间
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
	// 更新人
	UpdateBy string `json:"update_by" bson:"update_by"`
	// 创建请求
	CreateApplicationRequest
}

func (a *Application) String() string {
	return pretty.ToJSON(a)
}

type CreateApplicationRequest struct {
	// 创建人
	CreateBy string `json:"create_by" bson:"create_by" description:"创建人"`
	// 创建时间
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	// 应用所属空间名称
	Namespace string `json:"namespace" bson:"namespace" description:"应用所属空间名称"`
	// 应用创建参数
	CreateApplicationSpec
}

type CreateApplicationSpec struct {
	// 该应用是否已经准备就绪
	Ready bool `json:"ready" bson:"ready" description:"该应用是否已经准备就绪"`
	// 应用名称
	Name string `json:"name" bson:"name" description:"应用名称"`
	// 应用描述
	Description string `json:"description" bson:"description" description:"应用描述"`
	// 应用图标
	Icon string `json:"icon" bson:"icon" description:"应用图标"`
	// 应用类型
	Type TYPE `json:"type" bson:"type" description:"应用类型, SOURCE_CODE, CONTAINER_IMAGE, OTHER"`
	// 应用代码仓库信息
	CodeRepository CodeRepository `json:"code_repository" bson:"code_repository" description:"应用代码仓库信息"`
	// 应用镜像仓库信息
	ImageRepository ImageRepository `json:"image_repository" bson:"image_repository" description:"应用镜像仓库信息"`
	// 应用所有者
	Owner string `json:"owner" bson:"owner" description:"应用所有者"`
	// 应用等级, 评估这个应用的重要程度
	Level uint32 `json:"level" bson:"level" description:"应用等级, 评估这个应用的重要程度"`
	// 应用优先级, 应用启动的先后顺序
	Priority uint32 `json:"priority" bson:"priority" description:"应用优先级, 应用启动的先后顺序"`
	// 应用标签
	Labels map[string]string `json:"labels" bson:"labels" description:"应用标签"`
}

// CodeRepository 服务代码仓库信息
type CodeRepository struct {
	// 仓库提供商
	Provider SCM_PROVIDER `json:"provider" bson:"provider"`
	// token 操作仓库, 比如设置回调
	Token string `json:"token" bson:"token"`
	// 仓库对应的项目Id
	ProjectId string `json:"project_id" bson:"project_id"`
	// 仓库对应空间
	Namespace string `json:"namespace" bson:"namespace"`
	// 仓库web url地址
	WebUrl string `json:"web_url" bson:"web_url"`
	// 仓库ssh url地址
	SshUrl string `json:"ssh_url" bson:"ssh_url"`
	// 仓库http url地址
	HttpUrl string `json:"http_url" bson:"http_url"`
	// 源代码使用的编程语言, 构建时, 不同语言有不同的构建环境
	Language *LANGUAGE `json:"language" bson:"language"`
	// 开启Hook设置
	EnableHook bool `json:"enable_hook" bson:"enable_hook"`
	// Hook设置
	HookConfig string `json:"hook_config" bson:"hook_config"`
	// scm设置Hook后返回的id, 用于删除应用时，取消hook使用
	HookId string `json:"hook_id" bson:"hook_id"`
	// 仓库的创建时间
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

// ImageRepository 镜像仓库
type ImageRepository struct {
	// 服务镜像地址
	Address string `json:"address" bson:"address"`
}
