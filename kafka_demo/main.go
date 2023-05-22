package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Fetch.Max = 100
	config.Consumer.Fetch.Default = 2 * int32(time.Second) // 设置拉取间隔为 2 秒

	// 创建消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	// 指定要消费的主题和分区
	topic := "my-topic"
	partition := int32(0)

	// 创建分区消费者
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}
	defer partitionConsumer.Close()

	// 消费消息
	for message := range partitionConsumer.Messages() {
		fmt.Printf("Received message: %s\n", string(message.Value))
	}

	log.Println("Consumer stopped")
}
