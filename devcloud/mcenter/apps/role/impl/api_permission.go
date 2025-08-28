package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/endpoint"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/role"
	"gorm.io/gorm"
)

// 添加角色关联API
func (i *RoleServiceImpl) AddApiPermission(ctx context.Context, in *role.AddApiPermissionRequest) ([]*role.ApiPermission, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate add api permission error, %s", err)
	}

	perms := []*role.ApiPermission{}
	if err := datasource.DBFromCtx(ctx).Transaction(func(tx *gorm.DB) error {
		for i := range in.Items {
			item := in.Items[i]
			perm := role.NewApiPermission(in.RoleId, item)
			if err := tx.Save(perm).Error; err != nil {
				return err
			}
			perms = append(perms, perm)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return perms, nil
}

// 查询角色关联的权限条目
func (i *RoleServiceImpl) QueryApiPermission(ctx context.Context, in *role.QueryApiPermissionRequest) ([]*role.ApiPermission, error) {
	query := datasource.DBFromCtx(ctx).Model(&role.ApiPermission{})
	if len(in.RoleIds) > 0 {
		query = query.Where("role_id IN ?", in.RoleIds)
	}
	if len(in.ApiPermissionIds) > 0 {
		query = query.Where("id IN ?", in.ApiPermissionIds)
	}

	perms := []*role.ApiPermission{}
	if err := query.
		Order("created_at desc").
		Find(&perms).Error; err != nil {
		return nil, err
	}
	return perms, nil
}

// 移除角色关联API
func (i *RoleServiceImpl) RemoveApiPermission(ctx context.Context, in *role.RemoveApiPermissionRequest) ([]*role.ApiPermission, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	perms, err := i.QueryApiPermission(ctx, role.NewQueryApiPermissionRequest().AddRoleId(in.RoleId).AddPermissionId(in.ApiPermissionIds...))
	if err != nil {
		return nil, err
	}

	if err := datasource.DBFromCtx(ctx).
		Where("role_id = ?", in.RoleId).
		Where("id IN ?", in.ApiPermissionIds).
		Delete(&role.ApiPermission{}).
		Error; err != nil {
		return nil, err
	}

	return perms, nil
}

// 查询匹配到的Api接口列表
func (i *RoleServiceImpl) QueryMatchedEndpoint(ctx context.Context, in *role.QueryMatchedEndpointRequest) (*types.Set[*endpoint.Endpoint], error) {
	set := types.New[*endpoint.Endpoint]()

	// 查询角色的权限
	perms, err := i.QueryApiPermission(ctx, role.NewQueryApiPermissionRequest().AddRoleId(in.RoleIds...))
	if err != nil {
		return nil, err
	}

	// 查询服务的Endpoint列表
	endpointReq := endpoint.NewQueryEndpointRequest()
	for _, perm := range perms {
		endpointReq.WithService(perm.Service)
	}
	endpoints, err := endpoint.GetService().QueryEndpoint(ctx, endpointReq)
	if err != nil {
		return nil, err
	}

	// 找出能匹配的API
	endpoints.ForEach(func(t *endpoint.Endpoint) {
		for _, perm := range perms {
			if perm.IsMatch(t) {
				if !endpoint.IsEndpointExist(set, t) {
					set.Add(t)
				}
			}
		}
	})

	return set, nil
}
