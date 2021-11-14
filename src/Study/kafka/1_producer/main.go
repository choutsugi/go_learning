package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {

	// 1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // ACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 分区
	config.Producer.Return.Successes = true                   // 交付成功消息

	// 2.连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer func(client sarama.SyncProducer) {
		err := client.Close()
		if err != nil {
			fmt.Println("client close failed, err:", err)
		}
	}(client)

	// 3.封装消息
	msg := &sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("第三次"),
	}

	// 4.发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
