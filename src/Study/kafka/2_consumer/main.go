package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	// 创建消费者
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}

	//获取指定topic下的分区列表
	partitions, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitions)

	var wg sync.WaitGroup

	for partition := range partitions {
		// 为每个分区创建消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer for partition %d, err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费消息
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition: %d Offset:%d Key:%v Value:%v\n",
					msg.Partition,
					msg.Offset,
					msg.Key,
					msg.Value,
				)
			}
		}(pc)
	}
	wg.Wait()
}
