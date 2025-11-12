package application

import (
	"context"

	"github.com/infraboard/mcube/v2/types"
)

type Service interface {
	// CreateApplication 创建应用
	CreateApplication(context.Context, CreateApplicationRequest) (*Application, error)
	// QueryApplication 查询应用
	QueryApplication(context.Context, QueryApplicationRequest) (*types.Set[*Application], error)
	// UpdateApplication 更新应用
	UpdateApplication(context.Context, UpdateApplicationRequest) (*Application, error)
	// DeleteApplication 删除应用
	DeleteApplication(context.Context, DeleteApplicationRequest) (*Application, error)
	// DescribeApplication 获取应用
	DescribeApplication(context.Context, DescribeApplicationRequest) (*Application, error)
}

// QueryApplicationRequest 查询 请求
type QueryApplicationRequest struct {
	// 应用ID
	Id string `json:"id" bson:"_id"`
	// 应用名称
	Name string `json:"name" bson:"name"`
	// 应用状态
	Status string `json:"status" bson:"status"`
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
	// 应用ID
	Id string `json:"id" bson:"_id"`
}
