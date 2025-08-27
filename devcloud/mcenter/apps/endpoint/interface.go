package endpoint

import (
	"context"

	"slices"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/modules/iam/apps"
)

const (
	APP_NAME = "endpoint"
)

func GetService() Service {
	return ioc.Controller().Get(APP_NAME).(Service)
}

type Service interface {
	// 注册API接口
	RegistryEndpoint(context.Context, *RegistryEndpointRequest) (*types.Set[*Endpoint], error)
	// 查询API接口列表
	QueryEndpoint(context.Context, *QueryEndpointRequest) (*types.Set[*Endpoint], error)
	// 查询API接口详情
	DescribeEndpoint(context.Context, *DescribeEndpointRequest) (*Endpoint, error)
}

func NewQueryEndpointRequest() *QueryEndpointRequest {
	return &QueryEndpointRequest{}
}

type QueryEndpointRequest struct {
	Services []string `form:"services" json:"serivces"`
}

func (r *QueryEndpointRequest) WithService(services ...string) *QueryEndpointRequest {
	for _, service := range services {
		if !slices.Contains(r.Services, service) {
			r.Services = append(r.Services, services...)
		}
	}
	return r
}

func (r *QueryEndpointRequest) IsMatchAllService() bool {
	return slices.Contains(r.Services, "*")
}

func NewDescribeEndpointRequest() *DescribeEndpointRequest {
	return &DescribeEndpointRequest{}
}

type DescribeEndpointRequest struct {
	apps.GetRequest
}

func NewRegistryEndpointRequest() *RegistryEndpointRequest {
	return &RegistryEndpointRequest{
		Items: []*RouteEntry{},
	}
}

type RegistryEndpointRequest struct {
	Items []*RouteEntry `json:"items"`
}

func (r *RegistryEndpointRequest) AddItem(items ...*RouteEntry) *RegistryEndpointRequest {
	r.Items = append(r.Items, items...)
	return r
}

func (r *RegistryEndpointRequest) Validate() error {
	return validator.Validate(r)
}
