# 消息队列

[课件](https://gitee.com/infraboard/go-course/blob/master/extra/mq/README.md)

## 安装

安装: https://www.mcube.top/docs/framework/component/kafka/
    https://www.mcube.top/guide/component/kafka.html
```yaml
version: '2'
services:
  zoo1:
    image: wurstmeister/zookeeper
    restart: unless-stopped
    hostname: zoo1
    ports:
      - "2181:2181"
    container_name: zookeeper

  kafka1:
    image: wurstmeister/kafka
    ports:
      - "9092:9092"
    environment:
      KAFKA_ADVERTISED_HOST_NAME: localhost
      KAFKA_ZOOKEEPER_CONNECT: "zoo1:2181"
      KAFKA_BROKER_ID: 1
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_CREATE_TOPICS: "stream-in:1:1,stream-out:1:1"
    depends_on:
      - zoo1
    container_name: kafka
```

```sh
docker-compose up -d
```