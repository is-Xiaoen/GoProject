package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/mpaas/apps/application"
)

// CreateApplication 实现 application.Service.
func (i *ApplicationServiceImpl) CreateApplication(ctx context.Context, in *application.CreateApplicationRequest) (*application.Application, error) {
	ins, err := application.NewApplication(*in)
	if err != nil {
		return nil, err
	}

	if err := datasource.DBFromCtx(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// QueryApplication 实现 application.Service
func (i *ApplicationServiceImpl) QueryApplication(ctx context.Context, in *application.QueryApplicationRequest) (*types.Set[*application.Application], error) {
	set := types.New[*application.Application]()

	query := datasource.DBFromCtx(ctx).Model(&application.Application{})
	if in.Id != "" {
		query = query.Where("id = ?", in.Id)
	}

	if in.Name != "" {
		query = query.Where("name = ?", in.Name)
	}

	if in.Ready != nil {
		query = query.Where("ready = ?", *in.Ready)
	}

	in.GormResourceFilter(query)

	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.Order("created_at desc").
		Offset(int(in.ComputeOffset())).Limit(int(in.PageSize)).
		Find(&set.Items).Error
	if err != nil {
		return nil, err
	}
	return set, nil
}

// DescribeApplication 实现 application.Service.
func (i *ApplicationServiceImpl) DescribeApplication(context.Context, *application.DescribeApplicationRequest) (*application.Application, error) {
	panic("unimplemented")
}

// UpdateApplication 实现 application.Service.
func (i *ApplicationServiceImpl) UpdateApplication(context.Context, *application.UpdateApplicationRequest) (*application.Application, error) {
	panic("unimplemented")
}

// DeleteApplication 实现 application.Service.
func (i *ApplicationServiceImpl) DeleteApplication(context.Context, *application.DeleteApplicationRequest) (*application.Application, error) {
	panic("unimplemented")
}
