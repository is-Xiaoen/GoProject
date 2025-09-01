package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/namespace"
	"gorm.io/gorm"
)

// 创建空间
func (i *NameSpaceServiceImpl) CreateNamespace(ctx context.Context, in *namespace.CreateNamespaceRequest) (*namespace.Namespace, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	ins := namespace.NewNamespace()
	ins.CreateNamespaceRequest = *in

	if err := datasource.DBFromCtx(ctx).
		Create(ins).
		Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// 查询空间
func (i *NameSpaceServiceImpl) QueryNamespace(ctx context.Context, in *namespace.QueryNamespaceRequest) (*types.Set[*namespace.Namespace], error) {
	set := types.New[*namespace.Namespace]()

	query := datasource.DBFromCtx(ctx).Model(&namespace.Namespace{})
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

// 查询空间详情
func (i *NameSpaceServiceImpl) DescribeNamespace(ctx context.Context, in *namespace.DescribeNamespaceRequest) (*namespace.Namespace, error) {
	query := datasource.DBFromCtx(ctx)

	ins := &namespace.Namespace{}
	if err := query.Where("id = ?", in.Id).First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("namespace %d not found", in.Id)
		}
		return nil, err
	}

	return ins, nil
}

// 更新空间
func (i *NameSpaceServiceImpl) UpdateNamespace(ctx context.Context, in *namespace.UpdateNamespaceRequest) (*namespace.Namespace, error) {
	descReq := namespace.NewDescribeNamespaceRequest()
	descReq.SetId(in.Id)
	ins, err := i.DescribeNamespace(ctx, descReq)
	if err != nil {
		return nil, err
	}

	ins.CreateNamespaceRequest = in.CreateNamespaceRequest
	return ins, datasource.DBFromCtx(ctx).Where("id = ?", in.Id).Updates(ins).Error
}

// 删除空间
func (i *NameSpaceServiceImpl) DeleteNamespace(ctx context.Context, in *namespace.DeleteNamespaceRequest) (*namespace.Namespace, error) {
	descReq := namespace.NewDescribeNamespaceRequest()
	descReq.SetId(in.Id)
	ins, err := i.DescribeNamespace(ctx, descReq)
	if err != nil {
		return nil, err
	}

	return ins, datasource.DBFromCtx(ctx).
		Where("id = ?", in.Id).
		Delete(&namespace.Namespace{}).
		Error
}
