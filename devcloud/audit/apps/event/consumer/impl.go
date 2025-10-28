package consumer

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	ioc_kafka "github.com/infraboard/mcube/v2/ioc/config/kafka"
	kafka "github.com/segmentio/kafka-go"
)

func init() {
	ioc.Controller().Registry(&consumer{
		GroupId: "audit",
		Topics:  []string{"audit_go18"},
		ctx:     context.Background(),
	})
}

// 业务具体实现
type consumer struct {
	// 继承模版
	ioc.ObjectImpl

	// 模块子Logger
	log *zerolog.Logger

	// Kafka消费者
	reader *kafka.Reader
	// 允许时上下文
	ctx context.Context

	// 消费组Id
	GroupId string `toml:"group_id" json:"group_id" yaml:"group_id"  env:"GROUP_ID"`
	// 当前这个消费者 配置的topic
	Topics []string `toml:"topic" json:"topic" yaml:"topic"  env:"TOPIC"`
}

// 对象名称
func (i *consumer) Name() string {
	return "maudit_consumer"
}

// 初始化
func (i *consumer) Init() error {
	// 对象
	i.log = log.Sub(i.Name())
	i.reader = ioc_kafka.ConsumerGroup(i.GroupId, i.Topics)

	go i.Run(i.ctx)
	return nil
}

func (i *consumer) Close(ctx context.Context) error {
	i.ctx.Done()
	return nil
}
