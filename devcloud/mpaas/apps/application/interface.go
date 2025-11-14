package application

import (
	"context"

	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/policy"
)

const (
	APP_NAME = "application"
)

func GetService() Service {
	return ioc.Controller().Get(APP_NAME).(Service)
}

type Service interface {
	// CreateApplication 创建应用
	CreateApplication(context.Context, *CreateApplicationRequest) (*Application, error)
	// QueryApplication 查询应用
	QueryApplication(context.Context, *QueryApplicationRequest) (*types.Set[*Application], error)
	// UpdateApplication 更新应用
	UpdateApplication(context.Context, *UpdateApplicationRequest) (*Application, error)
	// DeleteApplication 删除应用
	DeleteApplication(context.Context, *DeleteApplicationRequest) (*Application, error)
	// DescribeApplication 获取应用
	DescribeApplication(context.Context, *DescribeApplicationRequest) (*Application, error)
}

// QueryApplicationRequest 查询 请求
type QueryApplicationRequest struct {
	policy.ResourceScope
	QueryApplicationRequestSpec
}

type QueryApplicationRequestSpec struct {
	*request.PageRequest
	// 应用ID
	Id string `json:"id" bson:"_id"`
	// 应用名称
	Name string `json:"name" bson:"name"`
	// 应用是否就绪
	Ready *bool `json:"ready" bson:"ready"`
}

// UpdateApplicationRequest 更新 请求
type UpdateApplicationRequest struct {
	// 更新人
	UpdateBy string `json:"update_by" bson:"update_by"`
	DescribeApplicationRequest
	CreateApplicationSpec
}

// DeleteApplicationRequest 删除 请求
type DeleteApplicationRequest struct {
	DescribeApplicationRequest
}

type DescribeApplicationRequest struct {
	policy.ResourceScope
	// 应用ID
	Id string `json:"id" bson:"_id"`
}
