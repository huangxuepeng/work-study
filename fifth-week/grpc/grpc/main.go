package main

import (
	"context"
	"fmt"
	"log"
	"time"
	pb "work/fifth-week/grpc/grpc/proto"

	cc "google.golang.org/grpc"
)

func main() {
	//创建一个grpc连接
	conn, err := cc.Dial("localhost:8002", cc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//创建RPC客户端
	client := pb.NewGreetsClient(conn)
	//设置超时时间
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用方法
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "小超", Message: "回来吃饭吗"})
	if err != nil {
		log.Fatalf("couldn not greet: %v", err)
	}
	for {
		rec, err := reply.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(rec)
	}
}
