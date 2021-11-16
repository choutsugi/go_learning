package etcd

import (
	"LogAgent/common"
	"LogAgent/logger"
	"LogAgent/system"
	"LogAgent/tailfile"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	client *clientv3.Client
)

func Init(address []string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		logger.Z.Errorf("etcd: init failed, err:%v", err)
		return
	}
	return
}

// GetConf 获取配置项
func GetConf(key string) (collectEntryList []common.CollectEntry, err error) {
	// 获取IP生成Key
	ip, err := system.GetLocalIPByDial()
	if err != nil {
		return
	}
	key = fmt.Sprintf(key, ip)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		logger.Z.Errorf("etcd: get conf failed by key %s, err:%v", key, err)
		return
	}
	logger.Z.Infof("etcd: get conf:%s success", key)

	if len(resp.Kvs) == 0 {
		logger.Z.Errorf("etcd: conf of key:%s is not exist", key)
		err = errors.New("conf not exist")
		return
	}

	ret := resp.Kvs[0]
	// 对从etcd获取的Json格式的配置数据进行解析
	err = json.Unmarshal(ret.Value, &collectEntryList)
	if err != nil {
		logger.Z.Errorf("etcd: conf of key:%s unmarshal failed, err:%v", key, err)
		return
	}
	return
}

// WatchConf 监视etcd配置变化
func WatchConf(key string) {
	watchChan := client.Watch(context.Background(), key)
	var newConf []common.CollectEntry
	for resp := range watchChan {
		for _, event := range resp.Events {
			newConf = []common.CollectEntry{}
			logger.Z.Info("etcd: configuration has been updated.")
			fmt.Printf("type:%s, key:%s, value:%s", event.Type, event.Kv.Key, event.Kv.Value)
			err := json.Unmarshal(event.Kv.Value, &newConf)
			if err != nil {
				logger.Z.Errorf("etcd: conf of key:%s unmarshal failed, err:%v", event.Kv.Key, err)
				continue
			}
			// 如果配置更新则通知tailfile刷新任务
			tailfile.UpdateConf(newConf)
		}
	}
}
