package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/role"
	"gorm.io/gorm"
)

// 创建角色
func (i *RoleServiceImpl) CreateRole(ctx context.Context, in *role.CreateRoleRequest) (*role.Role, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	ins := role.NewRole()
	ins.CreateRoleRequest = *in

	if err := datasource.DBFromCtx(ctx).
		Create(ins).
		Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// 列表查询
func (i *RoleServiceImpl) QueryRole(ctx context.Context, in *role.QueryRoleRequest) (*types.Set[*role.Role], error) {
	set := types.New[*role.Role]()

	query := datasource.DBFromCtx(ctx).Model(&role.Role{})
	if len(in.RoleIds) > 0 {
		query = query.Where("id IN ?", in.RoleIds)
		in.PageSize = uint64(len(in.RoleIds))
	}
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.
		Order("created_at desc").
		Offset(int(in.ComputeOffset())).
		Limit(int(in.PageSize)).
		Find(&set.Items).
		Error
	if err != nil {
		return nil, err
	}
	return set, nil
}

// 详情查询
func (i *RoleServiceImpl) DescribeRole(ctx context.Context, in *role.DescribeRoleRequest) (*role.Role, error) {
	query := datasource.DBFromCtx(ctx)

	ins := &role.Role{}
	if err := query.Where("id = ?", in.Id).First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("role %d not found", in.Id)
		}
		return nil, err
	}

	pm, err := i.QueryApiPermission(ctx, role.NewQueryApiPermissionRequest().AddRoleId(in.Id))
	if err != nil {
		return nil, err
	}
	ins.ApiPermissions = pm

	return ins, nil
}

// 更新角色
func (i *RoleServiceImpl) UpdateRole(ctx context.Context, in *role.UpdateRoleRequest) (*role.Role, error) {
	descReq := role.NewDescribeRoleRequest()
	descReq.SetId(in.Id)
	ins, err := i.DescribeRole(ctx, descReq)
	if err != nil {
		return nil, err
	}

	ins.CreateRoleRequest = in.CreateRoleRequest
	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Updates(ins).Error
}

// 删除角色
func (i *RoleServiceImpl) DeleteRole(ctx context.Context, in *role.DeleteRoleRequest) (*role.Role, error) {
	descReq := role.NewDescribeRoleRequest()
	descReq.SetId(in.Id)
	ins, err := i.DescribeRole(ctx, descReq)
	if err != nil {
		return nil, err
	}

	return ins, datasource.DBFromCtx(ctx).
		Where("id = ?", in.Id).
		Delete(&role.Role{}).
		Error
}
