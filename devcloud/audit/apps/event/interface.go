package event

import (
	"context"

	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
)

var (
	AppName = "event"
)

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

type Service interface {
	// SaveEvent 存储
	SaveEvent(context.Context, *types.Set[*Event]) error
	// QueryEvent 查询
	QueryEvent(context.Context, *QueryEventRequest) (*types.Set[*Event], error)
}

func NewQueryEventRequest() *QueryEventRequest {
	return &QueryEventRequest{
		PageRequest: request.NewDefaultPageRequest(),
	}
}

type QueryEventRequest struct {
	// 分页请求参数
	*request.PageRequest
}
