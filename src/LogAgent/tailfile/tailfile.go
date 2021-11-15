package tailfile

import (
	"LogAgent/common"
	"LogAgent/kafka"
	"LogAgent/logger"
	"strings"
	"time"

	"github.com/Shopify/sarama"

	"github.com/hpcloud/tail"
)

// 根据etcd配置中的每一组path和topic创建task对象
type tailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

// tailTask初始化
func (t *tailTask) init() (err error) {
	config := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}

	t.instance, err = tail.TailFile(t.path, config)
	return
}

// 开启tailTask任务
func (t *tailTask) run() {
	logger.Z.Infof("collect for path:%s is running...", t.path)
	for {
		line, ok := <-t.instance.Lines
		if !ok {
			logger.Z.Warnf("tail file close reopen, filename:%s\n", t.path)
			time.Sleep(1 * time.Second)
			continue
		}
		// 过滤空行
		if len(strings.Trim(line.Text, "\r")) == 0 {
			continue
		}

		// 使用channel实现异步发送
		msg := &sarama.ProducerMessage{
			Topic: t.topic,
			Value: sarama.StringEncoder(line.Text),
		}
		kafka.Write(msg)
		logger.Z.Infof("log:%s", line.Text)
	}
}

func newTailTask(conf common.CollectEntry) (task *tailTask) {
	task = &tailTask{
		path:  conf.Path,
		topic: conf.Topic,
	}
	return task
}
