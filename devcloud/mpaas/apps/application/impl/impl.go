package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/is-Xiaoen/GoProject/devcloud/mpaas/apps/application"
)

func init() {
	ioc.Controller().Registry(&ApplicationServiceImpl{})
}

var _ application.Service = (*ApplicationServiceImpl)(nil)

type ApplicationServiceImpl struct {
	ioc.ObjectImpl
}

func (i *ApplicationServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&application.Application{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *ApplicationServiceImpl) Name() string {
	return application.APP_NAME
}
