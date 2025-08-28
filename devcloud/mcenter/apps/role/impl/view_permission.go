package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/role"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/view"
	"gorm.io/gorm"
)

// 添加角色关联菜单
func (i *RoleServiceImpl) AddViewPermission(ctx context.Context, in *role.AddViewPermissionRequest) ([]*role.ViewPermission, error) {
	if err := in.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate add view permission error, %s", err)
	}

	perms := []*role.ViewPermission{}
	if err := datasource.DBFromCtx(ctx).Transaction(func(tx *gorm.DB) error {
		for i := range in.Items {
			item := in.Items[i]
			perm := role.NewViewPermission(in.RoleId, item)
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

// 查询角色关联的视图权限
func (i *RoleServiceImpl) QueryViewPermission(ctx context.Context, in *role.QueryViewPermissionRequest) ([]*role.ViewPermission, error) {
	query := datasource.DBFromCtx(ctx).Model(&role.ViewPermission{})
	if len(in.RoleIds) > 0 {
		query = query.Where("role_id IN ?", in.RoleIds)
	}
	if len(in.ViewPermissionIds) > 0 {
		query = query.Where("in IN ?", in.ViewPermissionIds)
	}

	perms := []*role.ViewPermission{}
	if err := query.Order("created_at desc").
		Where("id IN ?", in.RoleIds).
		Find(&perms).Error; err != nil {
		return nil, err
	}
	return perms, nil
}

// 移除角色关联菜单
func (i *RoleServiceImpl) RemoveViewPermission(ctx context.Context, in *role.RemoveViewPermissionRequest) ([]*role.ViewPermission, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	perms, err := i.QueryViewPermission(ctx, role.NewQueryViewPermissionRequest().AddRoleId(in.RoleId).AddPermissionId(in.ViewPermissionIds...))
	if err != nil {
		return nil, err
	}

	if err := datasource.DBFromCtx(ctx).
		Where("role_id = ?", in.RoleId).
		Where("id IN ?", in.ViewPermissionIds).
		Delete(&role.ViewPermission{}).
		Error; err != nil {
		return nil, err
	}

	return perms, nil
}

// 查询能匹配到视图菜单
func (i *RoleServiceImpl) QueryMatchedPage(ctx context.Context, in *role.QueryMatchedPageRequest) (*types.Set[*view.Menu], error) {
	return nil, nil
}
