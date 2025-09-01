package policy

import (
	"context"

	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/modules/iam/apps"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/endpoint"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/namespace"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/view"
)

const (
	AppName = "policy"
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	// 策略管理
	PolicyService
	// 权限查询, 整合用户多个角色的权限合集
	PermissionService
}

type PolicyService interface {
	// 创建策略
	CreatePolicy(context.Context, *CreatePolicyRequest) (*Policy, error)
	// 查询策略列表
	QueryPolicy(context.Context, *QueryPolicyRequest) (*types.Set[*Policy], error)
	// 查询详情
	DescribePolicy(context.Context, *DescribePolicyRequest) (*Policy, error)
	// 更新策略
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*Policy, error)
	// 删除策略
	DeletePolicy(context.Context, *DeletePolicyRequest) (*Policy, error)
}

func NewQueryPolicyRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{
		PageRequest: request.NewDefaultPageRequest(),
	}
}

type QueryPolicyRequest struct {
	*request.PageRequest
	// 忽略分页
	SkipPage bool `json:"skip_page"`
	// 关联用户Id
	UserId *uint64 `json:"user_id"`
	// 关联空间
	NamespaceId *uint64 `json:"namespace_id"`
	// 没有过期
	Expired *bool `json:"expired"`
	// 有没有启动
	Enabled *bool `json:"active"`
	// 关联查询出空间对象
	WithNamespace bool `json:"with_namespace"`
	// 关联查询出用户对象
	WithUser bool `json:"with_user"`
	// 关联查询角色对象
	WithRole bool `json:"with_role"`
}

func (r *QueryPolicyRequest) SetNamespaceId(nsId uint64) *QueryPolicyRequest {
	r.NamespaceId = &nsId
	return r
}

func (r *QueryPolicyRequest) SetUserId(uid uint64) *QueryPolicyRequest {
	r.UserId = &uid
	return r
}

func (r *QueryPolicyRequest) SetExpired(v bool) *QueryPolicyRequest {
	r.Expired = &v
	return r
}

func (r *QueryPolicyRequest) SetEnabled(v bool) *QueryPolicyRequest {
	r.Enabled = &v
	return r
}

func (r *QueryPolicyRequest) SetSkipPage(v bool) *QueryPolicyRequest {
	r.SkipPage = v
	return r
}

func (r *QueryPolicyRequest) SetWithRole(v bool) *QueryPolicyRequest {
	r.WithRole = v
	return r
}
func (r *QueryPolicyRequest) SetWithUsers(v bool) *QueryPolicyRequest {
	r.WithUser = v
	return r
}
func (r *QueryPolicyRequest) SetWithUser(v bool) *QueryPolicyRequest {
	r.WithNamespace = v
	return r
}

func NewDescribePolicyRequest() *DescribePolicyRequest {
	return &DescribePolicyRequest{}
}

type DescribePolicyRequest struct {
	apps.GetRequest
}

type UpdatePolicyRequest struct {
	apps.GetRequest
	CreatePolicyRequest
}

func NewDeletePolicyRequest() *DeletePolicyRequest {
	return &DeletePolicyRequest{}
}

type DeletePolicyRequest struct {
	apps.GetRequest
}

type PermissionService interface {
	// 查询用户可以访问的空间
	QueryNamespace(context.Context, *QueryNamespaceRequest) (*types.Set[*namespace.Namespace], error)
	// 查询用户可以访问的菜单
	QueryMenu(context.Context, *QueryMenuRequest) (*types.Set[*view.Menu], error)
	// 查询用户可以访问的Api接口
	QueryEndpoint(context.Context, *QueryEndpointRequest) (*types.Set[*endpoint.Endpoint], error)
	// 校验页面权限
	ValidatePagePermission(context.Context, *ValidatePagePermissionRequest) (*ValidatePagePermissionResponse, error)
	// 校验接口权限
	ValidateEndpointPermission(context.Context, *ValidateEndpointPermissionRequest) (*ValidateEndpointPermissionResponse, error)
}

type ValidatePagePermissionRequest struct {
	UserId      uint64 `json:"user_id" form:"user_id"`
	NamespaceId uint64 `json:"namespace_id" form:"namespace_id"`
	Path        string `json:"path" form:"path"`
}

func NewValidatePagePermissionResponse(req ValidatePagePermissionRequest) *ValidatePagePermissionResponse {
	return &ValidatePagePermissionResponse{
		ValidatePagePermissionRequest: req,
	}
}

type ValidatePagePermissionResponse struct {
	ValidatePagePermissionRequest
	HasPermission bool       `json:"has_permission"`
	Page          *view.Page `json:"page"`
}

func NewValidateEndpointPermissionRequest() *ValidateEndpointPermissionRequest {
	return &ValidateEndpointPermissionRequest{}
}

type ValidateEndpointPermissionRequest struct {
	UserId      uint64 `json:"user_id" form:"user_id"`
	NamespaceId uint64 `json:"namespace_id" form:"namespace_id"`
	Service     string `json:"service" form:"service"`
	Path        string `json:"path" form:"path"`
	Method      string `json:"method" form:"method"`
}

func NewValidateEndpointPermissionResponse(req ValidateEndpointPermissionRequest) *ValidateEndpointPermissionResponse {
	return &ValidateEndpointPermissionResponse{
		ValidateEndpointPermissionRequest: req,
	}
}

type ValidateEndpointPermissionResponse struct {
	ValidateEndpointPermissionRequest
	HasPermission bool               `json:"has_permission"`
	Endpoint      *endpoint.Endpoint `json:"endpoint"`
}

func (r *ValidateEndpointPermissionResponse) String() string {
	return pretty.ToJSON(r)
}

func NewQueryNamespaceRequest() *QueryNamespaceRequest {
	return &QueryNamespaceRequest{}
}

type QueryNamespaceRequest struct {
	UserId      uint64 `json:"user_id"`
	NamespaceId uint64 `json:"namespace_id"`
}

func (r *QueryNamespaceRequest) SetUserId(v uint64) *QueryNamespaceRequest {
	r.UserId = v
	return r
}

func (r *QueryNamespaceRequest) SetNamespaceId(v uint64) *QueryNamespaceRequest {
	r.NamespaceId = v
	return r
}

func NewQueryMenuRequest() *QueryMenuRequest {
	return &QueryMenuRequest{}
}

type QueryMenuRequest struct {
	UserId      uint64 `json:"user_id"`
	NamespaceId uint64 `json:"namespace_id"`
}

func NewQueryEndpointRequest() *QueryEndpointRequest {
	return &QueryEndpointRequest{}
}

type QueryEndpointRequest struct {
	UserId      uint64 `json:"user_id"`
	NamespaceId uint64 `json:"namespace_id"`
}

func (r *QueryEndpointRequest) SetUserId(v uint64) *QueryEndpointRequest {
	r.UserId = v
	return r
}

func (r *QueryEndpointRequest) SetNamespaceId(v uint64) *QueryEndpointRequest {
	r.NamespaceId = v
	return r
}
