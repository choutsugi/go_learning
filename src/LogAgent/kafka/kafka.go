package kafka

import (
	"LogAgent/logger"

	"github.com/Shopify/sarama"
)

var (
	client  sarama.SyncProducer
	msgChan chan *sarama.ProducerMessage
)

// Write 向kafka写消息
func Write(msg *sarama.ProducerMessage) {
	msgChan <- msg
}

func Init(address []string, chanSize int64) (err error) {
	// 1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 2.连接kafka
	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		logger.Z.Errorf("kafka: producer closed, err:%v", err)
		return
	}

	// 3.初始化MsgChan
	msgChan = make(chan *sarama.ProducerMessage, chanSize)

	// 4.启动后台goroutine用于发送
	go sendMsg()

	return
}

func sendMsg() {
	for {
		select {
		case msg := <-msgChan:
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				logger.Z.Warningf("send msg failed, err:%v", err)
				return
			}
			logger.Z.Infof("send msg to kafka success, pid:%v offset:%v", pid, offset)
		}
	}
}
