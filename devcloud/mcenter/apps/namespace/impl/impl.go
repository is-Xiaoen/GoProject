package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/namespace"
)

func init() {
	ioc.Controller().Registry(&NameSpaceServiceImpl{})
}

var _ namespace.Service = (*NameSpaceServiceImpl)(nil)

type NameSpaceServiceImpl struct {
	ioc.ObjectImpl
}

func (i *NameSpaceServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&namespace.Namespace{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *NameSpaceServiceImpl) Name() string {
	return namespace.AppName
}
