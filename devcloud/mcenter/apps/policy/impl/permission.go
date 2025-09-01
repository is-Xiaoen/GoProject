package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/endpoint"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/namespace"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/policy"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/role"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/view"
)

// 查询用户可以访问的空间
func (i *PolicyServiceImpl) QueryNamespace(ctx context.Context, in *policy.QueryNamespaceRequest) (*types.Set[*namespace.Namespace], error) {
	nsReq := namespace.NewQueryNamespaceRequest()

	policies, err := i.QueryPolicy(ctx,
		policy.NewQueryPolicyRequest().
			SetSkipPage(true).
			SetUserId(in.UserId).
			SetExpired(false).
			SetEnabled(true))
	if err != nil {
		return nil, err
	}

	policies.ForEach(func(t *policy.Policy) {
		if t.NamespaceId != nil {
			nsReq.AddNamespaceIds(*t.NamespaceId)
		}
	})

	return i.namespace.QueryNamespace(ctx, nsReq)
}

// 查询用户可以访问的Api接口
// 找到用户可以访问的角色列表，然后在找出角色对应的Api访问权限
func (i *PolicyServiceImpl) QueryEndpoint(ctx context.Context, in *policy.QueryEndpointRequest) (*types.Set[*endpoint.Endpoint], error) {
	set := types.New[*endpoint.Endpoint]()
	policies, err := i.QueryPolicy(ctx,
		policy.NewQueryPolicyRequest().
			SetSkipPage(true).
			SetNamespaceId(in.NamespaceId).
			SetUserId(in.UserId).
			SetExpired(false).
			SetEnabled(true))
	if err != nil {
		return nil, err
	}

	roleReq := role.NewQueryMatchedEndpointRequest()
	policies.ForEach(func(t *policy.Policy) {
		roleReq.Add(t.RoleId)
	})

	if policies.Len() > 0 {
		set, err = role.GetService().QueryMatchedEndpoint(ctx, roleReq)
		if err != nil {
			return nil, err
		}
	}

	return set, nil
}

// 校验Api接口权限
func (i *PolicyServiceImpl) ValidateEndpointPermission(ctx context.Context, in *policy.ValidateEndpointPermissionRequest) (*policy.ValidateEndpointPermissionResponse, error) {
	resp := policy.NewValidateEndpointPermissionResponse(*in)

	// 空间Owner有所有权限
	ns, err := namespace.GetService().DescribeNamespace(ctx, namespace.NewDescribeNamespaceRequest().SetNamespaceId(in.NamespaceId))
	if err != nil {
		return nil, err
	}
	if ns.IsOwner(in.UserId) {
		resp.HasPermission = true
		return resp, nil
	}

	// 非空间管理员需要独立鉴权, 查询用户可以访问的API列表
	endpointReq := policy.NewQueryEndpointRequest()
	endpointReq.UserId = in.UserId
	endpointReq.NamespaceId = in.NamespaceId
	endpointSet, err := i.QueryEndpoint(ctx, endpointReq)
	if err != nil {
		return nil, err
	}
	for _, item := range endpointSet.Items {
		if item.IsMatched(in.Service, in.Method, in.Path) {
			resp.HasPermission = true
			resp.Endpoint = item
			break
		}
	}
	return resp, nil
}

// 查询用户可以访问的菜单
func (i *PolicyServiceImpl) QueryMenu(ctx context.Context, in *policy.QueryMenuRequest) (*types.Set[*view.Menu], error) {
	return nil, nil
}

// 校验Menu视图权限
func (i *PolicyServiceImpl) ValidatePagePermission(ctx context.Context, in *policy.ValidatePagePermissionRequest) (*policy.ValidatePagePermissionResponse, error) {
	return nil, nil
}
