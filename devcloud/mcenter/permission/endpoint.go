package permission

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/endpoint"
	"github.com/rs/zerolog"
)

func init() {
	ioc.Api().Registry(&ApiRegister{})
}

func GetApiRegister() *ApiRegister {
	return ioc.Api().Get("api_register").(*ApiRegister)
}

// 接口注册模块: 扫描当前GoResuful Container下所有路径，并完成注册
type ApiRegister struct {
	ioc.ObjectImpl

	log *zerolog.Logger
}

func (c *ApiRegister) Name() string {
	return "api_register"
}

// 这个Init一定要放到所有的路由都添加完成后进行
func (i *ApiRegister) Priority() int {
	return -100
}

func (a *ApiRegister) Init() error {
	a.log = log.Sub(a.Name())
	// 注册认证中间件
	entries := endpoint.NewEntryFromRestfulContainer(gorestful.RootRouter())
	req := endpoint.NewRegistryEndpointRequest()
	req.AddItem(entries...)
	set, err := endpoint.GetService().RegistryEndpoint(context.Background(), req)
	if err != nil {
		return err
	}
	a.log.Info().Msgf("registry endpoinst: %s", set.Items)
	return nil
}
