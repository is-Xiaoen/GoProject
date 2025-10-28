package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/is-Xiaoen/GoProject/devcloud/audit/apps/event"
	"github.com/rs/zerolog"

	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
	"go.mongodb.org/mongo-driver/mongo"

	// 引入消费者
	_ "github.com/is-Xiaoen/GoProject/devcloud/audit/apps/event/consumer"
)

func init() {
	ioc.Controller().Registry(&EventServiceImpl{})
}

var _ event.Service = (*EventServiceImpl)(nil)

// EventServiceImpl 业务具体实现
type EventServiceImpl struct {
	// 继承模板
	ioc.ObjectImpl

	// 模块子Logger
	log *zerolog.Logger

	// MongoDB 集合
	col *mongo.Collection
}

// Name 对象名称
func (i *EventServiceImpl) Name() string {
	return event.AppName
}

func (i *EventServiceImpl) Priority() int {
	return event.PRIORITY
}

// Init 初始化
func (i *EventServiceImpl) Init() error {
	// 对象
	i.log = log.Sub(i.Name())

	i.log.Debug().Msgf("database: %s", ioc_mongo.Get().Database)
	// 需要一个集合Collection
	i.col = ioc_mongo.DB().Collection("events")
	return nil
}
