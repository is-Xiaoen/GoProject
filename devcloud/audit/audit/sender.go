package audit

import (
	"context"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/is-Xiaoen/GoProject/devcloud/audit/apps/event"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/endpoint"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
)

// SendEvent 审计日志的发送逻辑
func (a *EventSender) SendEvent() restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, fc *restful.FilterChain) {
		sr := req.SelectedRoute()
		md := NewMetaData(sr.Metadata())

		// 开关打开 , 则开启审计
		if md.GetBool(META_AUDIT_KEY) {
			// 获取当前是否需要审计
			e := event.NewEvent()

			// 用户信息
			tk := token.GetTokenFromCtx(req.Request.Context())
			if tk != nil {
				e.Who = tk.UserName
				e.Namespace = tk.NamespaceName
			}

			// ioc 里面获取当前应用的名称
			e.Service = application.Get().AppName
			e.ResourceType = md.GetString(endpoint.META_RESOURCE_KEY)
			e.Action = md.GetString(endpoint.META_ACTION_KEY)

			// {id} /:id
			e.ResourceId = req.PathParameter("id")
			e.UserAgent = req.Request.UserAgent()
			e.Extras["method"] = sr.Method()
			e.Extras["path"] = sr.Path()
			e.Extras["operation"] = sr.Operation()

			// 补充处理后的数据
			e.StatusCode = resp.StatusCode()
			// 发送topic, 使用这个中间件使用者,需要配置kafka
			err := a.wirter.WriteMessages(context.Background(), e.ToKafkaMessage())
			if err != nil {
				a.log.Error().Msgf("send message error, %s", err)
			} else {
				a.log.Debug().Msgf("send audit event ok, who: %s, resource: %s, action: %s", e.Who, e.ResourceType, e.Action)
			}
		}

		// 路由给后续逻辑
		fc.ProcessFilter(req, resp)
	}
}

func NewMetaData(data map[string]any) *MetaData {
	return &MetaData{
		data: data,
	}
}

type MetaData struct {
	data map[string]any
}

func (m *MetaData) GetString(key string) string {
	if v, ok := m.data[key]; ok {
		return v.(string)
	}
	return ""
}

func (m *MetaData) GetBool(key string) bool {
	if v, ok := m.data[key]; ok {
		return v.(bool)
	}
	return false
}
