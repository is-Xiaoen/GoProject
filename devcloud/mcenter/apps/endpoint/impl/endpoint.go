package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/endpoint"
	"gorm.io/gorm"
)

// 注册API接口
// 这是一个批量接口, 一次添加多条记录
// 需要保证事务: 同时成功，或者同时失败, MySQL事务
func (i *EndpointServiceImpl) RegistryEndpoint(ctx context.Context, in *endpoint.RegistryEndpointRequest) (*types.Set[*endpoint.Endpoint], error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	set := types.New[*endpoint.Endpoint]()
	err := datasource.DBFromCtx(ctx).Transaction(func(tx *gorm.DB) error {
		for i := range in.Items {
			item := in.Items[i].BuildUUID()
			ins := endpoint.NewEndpoint().SetRouteEntry(*item)

			oldEnpoint := endpoint.NewEndpoint()
			if err := tx.Where("uuid = ?", item.UUID).Take(oldEnpoint).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					return err
				}

				// 需要创建
				if err := tx.Save(ins).Error; err != nil {
					return err
				}
			} else {
				// 需要更新
				ins.Id = oldEnpoint.Id
				if err := tx.Where("uuid = ?", item.UUID).Updates(ins).Error; err != nil {
					return err
				}

			}
			set.Add(ins)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return set, nil
}

// 查询API接口列表
func (i *EndpointServiceImpl) QueryEndpoint(ctx context.Context, in *endpoint.QueryEndpointRequest) (*types.Set[*endpoint.Endpoint], error) {
	set := types.New[*endpoint.Endpoint]()

	query := datasource.DBFromCtx(ctx).Model(&endpoint.Endpoint{})
	if len(in.Services) > 0 && !in.IsMatchAllService() {
		query = query.Where("service IN ?", in.Services)
	}

	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	err = query.
		Order("created_at desc").
		Find(&set.Items).
		Error
	if err != nil {
		return nil, err
	}
	return set, nil
}

// 查询API接口详情
func (i *EndpointServiceImpl) DescribeEndpoint(ctx context.Context, in *endpoint.DescribeEndpointRequest) (*endpoint.Endpoint, error) {
	query := datasource.DBFromCtx(ctx)

	ins := &endpoint.Endpoint{}
	if err := query.First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("endpoint %d not found", in.Id)
		}
		return nil, err
	}

	return ins, nil
}
