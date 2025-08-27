package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/endpoint"
)

func init() {
	ioc.Controller().Registry(&EndpointServiceImpl{})
}

var _ endpoint.Service = (*EndpointServiceImpl)(nil)

// 他是user service 服务的控制器
type EndpointServiceImpl struct {
	ioc.ObjectImpl
}

func (i *EndpointServiceImpl) Init() error {
	// 自动创建表
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&endpoint.Endpoint{})
		if err != nil {
			return err
		}
	}
	return nil
}

// 定义托管到Ioc里面的名称
func (i *EndpointServiceImpl) Name() string {
	return endpoint.APP_NAME
}
