package main

import (
	"fmt"
	"net"

	pb "work/fifth-week/grpc/grpc/proto"

	cc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedGreetsServer
}

func (s *server) SayHello(request *pb.HelloRequest, se pb.Greets_SayHelloServer) error {
	fmt.Println(request)
	var err error
	for i := 0; i < 10; i++ {
		if i == 0 {
			err = se.Send(&pb.HelloReply{Name: request.GetName(), Message: request.GetMessage()})
		} else {
			err = se.Send(&pb.HelloReply{Name: request.GetName()})
		}
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
func main() {
	//协议类型以及ip，port
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		fmt.Println(err)
		return
	}
	//定义一个rpc的server
	ss := cc.NewServer()
	//注册服务，相当与注册SayHello接口
	pb.RegisterGreetsServer(ss, &server{})
	//进行映射绑定
	reflection.Register(ss)

	//启动服务
	err = ss.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
