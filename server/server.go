package main

import (
	"context"
	"fmt"
	pb "work/fifth-week/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

// 实现在proto 定义的SayHello 接口, 其实就是重写了一个方法
func (s *Server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloReply, error) {
	// 获取参数
	fmt.Println(in.GetName(), in.GetMessage())
	return &pb.SayHelloReply{Message: "hello" + in.GetMessage(), Name: "world" + in.GetName()}, nil
}
