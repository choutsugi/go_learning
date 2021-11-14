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

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	_, err = client.Put(ctx, "group", "LAB")
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	cancel()

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*1)
	rsp, err := client.Get(ctx, "group")
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}

	for _, kv := range rsp.Kvs {
		fmt.Printf("key:%s value:%s\n", kv.Key, kv.Value)
	}
	cancel()
}
