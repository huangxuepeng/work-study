package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:23279", "localhost:33279"}, //etcd集群三个实例的端口
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")

	defer cli.Close()

	for true {
		rch := cli.Watch(context.Background(), "/logagent/conf/", clientv3.WithPrefix()) //阻塞在这里，如果没有key里没有变化，就一直停留在这里
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
}
