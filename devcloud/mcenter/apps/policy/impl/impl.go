package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/namespace"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/policy"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/role"
)

func init() {
	ioc.Controller().Registry(&PolicyServiceImpl{})
}

var _ policy.Service = (*PolicyServiceImpl)(nil)

type PolicyServiceImpl struct {
	ioc.ObjectImpl

	namespace namespace.Service
	role      role.Service
}

func (i *PolicyServiceImpl) Init() error {
	if datasource.Get().AutoMigrate {
		err := datasource.DB().AutoMigrate(&policy.Policy{})
		if err != nil {
			return err
		}
	}
	i.namespace = namespace.GetService()
	i.role = role.GetService()
	return nil
}

func (i *PolicyServiceImpl) Name() string {
	return policy.AppName
}
