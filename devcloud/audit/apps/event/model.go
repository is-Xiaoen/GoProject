package event

import (
	"encoding/json"
	"time"

	"github.com/rs/xid"
	"github.com/segmentio/kafka-go"
)

func NewEvent() *Event {
	return &Event{
		Id:     xid.New().String(),
		Label:  map[string]string{},
		Extras: map[string]string{},
		Time:   time.Now().Unix(),
	}
}

// Event 用户操作事件
// 如何映射成 MongoDB BSON
type Event struct {
	// 事件Id
	// _id 在mongodb 表示的是对象Id
	Id string `json:"id" bson:"_id"`
	// 谁
	Who string `json:"who" bson:"who"`
	// 在什么时间
	Time int64 `json:"time" bson:"time"`
	// 操作人Ip
	Ip string `json:"ip" bson:"ip"`
	// User Agent
	UserAgent string `json:"user_agent" bson:"user_agent"`

	// 做了什么操作, 服务:资源:动作
	// 服务 <cmdb, mcenter, ....>
	Service string `json:"service" bson:"service"`
	// 资源 <secret, user, namespace, ...>
	ResourceType string `json:"resource_type" bson:"resource_type"`
	// 动作 <list, get, update, create, delete, ....>
	Action string `json:"action" bson:"action"`

	// 详情信息
	ResourceId string `json:"resource_id" bson:"resource_id"`
	// 状态码 404
	StatusCode int `json:"status_code" bson:"status_code"`
	// 具体信息
	ErrorMessage string `json:"error_message" bson:"error_message"`

	// 标签
	Label map[string]string `json:"label" bson:"label"`
	// 扩展信息
	Extras map[string]string `json:"extras" bson:"extras"`
}

func (e *Event) Load(data []byte) error {
	return json.Unmarshal(data, e)
}

func (e *Event) ToKafkaMessage() kafka.Message {
	data, _ := json.Marshal(e)
	return kafka.Message{
		Value: data,
	}
}
