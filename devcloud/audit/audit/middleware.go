package audit

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	ioc_kafka "github.com/infraboard/mcube/v2/ioc/config/kafka"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/permission"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
)

func init() {
	ioc.Config().Registry(&EventSender{
		Topic: "audit_go18",
	})
}

// EventSender 审计中间件
type EventSender struct {
	ioc.ObjectImpl
	log *zerolog.Logger

	// 当前这个消费者 配置的topic
	Topic string `toml:"topic" json:"topic" yaml:"topic"  env:"TOPIC"`
	//
	wirter *kafka.Writer
}

// Name 中间件对象名称
func (c *EventSender) Name() string {
	return "audit_middleware"
}

func (c *EventSender) Priority() int {
	return permission.GetCheckerPriority() - 1
}

func (c *EventSender) Init() error {
	c.log = log.Sub(c.Name())
	c.wirter = ioc_kafka.Producer(c.Topic)

	// 注册认证中间件
	gorestful.RootRouter().Filter(c.SendEvent())
	return nil
}
