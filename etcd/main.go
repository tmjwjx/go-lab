package main

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	// 1. 连接 etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("连接失败：", err)
	}
	defer cli.Close()
	log.Println("成功连接 etcd！")

	// 2. 写入键值
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = cli.Put(ctx, "windows_key", "windows_value")
	cancel()
	if err != nil {
		log.Fatal("写入失败：", err)
	}
	log.Println("写入键值成功！")

	// 3. 读取键值
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Get(ctx, "windows_key")
	cancel()
	if err != nil {
		log.Fatal("读取失败：", err)
	}
	for _, kv := range resp.Kvs {
		log.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
	}
}
