package kafka_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"testing"

	"github.com/segmentio/kafka-go"
)

// TestListTopics tests listing Kafka topics using the kafka-go library.
// maudit_new
// stream-out
// maudit
func TestListTopics(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}

// TestCreateTopic tests the creation of a Kafka topic using the kafka-go library.
func TestCreateTopic(t *testing.T) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	err = controllerConn.CreateTopics(kafka.TopicConfig{Topic: "audit_go18", NumPartitions: 3, ReplicationFactor: 1})

	if err != nil {
		t.Fatal(err)
	}
}

// TestWriteMessage tests writing messages to a Kafka topic using the kafka-go library.
// kafka.Writer
func TestWriteMessage(t *testing.T) {
	// make a writer that produces to topic-A, using the least-bytes distribution
	publisher := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),
		// NOTE: When Topic is not defined here, each Message must define it instead.
		Topic:    "audit_go18",
		Balancer: &kafka.LeastBytes{},
		// The topic will be created if it is missing.
		AllowAutoTopicCreation: true,
		// 支持消息压缩
		// Compression: kafka.Snappy,
		// 支持TLS
		// Transport: &kafka.Transport{
		//     TLS: &tls.Config{},
		// }
	}
	defer publisher.Close()

	err := publisher.WriteMessages(context.Background(),
		kafka.Message{
			// 支持 Writing to multiple topics
			//  NOTE: Each Message has Topic defined, otherwise an error is returned.
			// Topic: "topic-A",
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

}

// TestReadMessage tests reading messages from a Kafka topic using the kafka-go library.
func TestReadMessage(t *testing.T) {
	// make a new reader that consumes from topic-A
	subscriber := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		// Consumer Groups, 不指定就是普通的一个Consumer
		GroupID: "devcloud-go18-audit",
		// 可以指定Partition消费消息
		// Partition: 0,
		Topic:    "audit_go18",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer subscriber.Close()

	for {
		// subscriber.FetchMessage(context.Background())
		m, err := subscriber.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		// 处理完消息后需要提交该消息已经消费完成, 消费者挂掉后保存消息消费的状态
		// if err := r.CommitMessages(ctx, m); err != nil {
		//     log.Fatal("failed to commit messages:", err)
		// }
	}
}
