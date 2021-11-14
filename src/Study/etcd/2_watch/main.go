package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	defer client.Close()

	// watch
	watchCh := client.Watch(context.Background(), "group")
	for rsp := range watchCh {
		for _, event := range rsp.Events {
			fmt.Printf("type:%s key:%s value:%s\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}
}
