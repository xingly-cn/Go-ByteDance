package main

import (
	"context"
	service "day4/Grpc/hello-server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type server struct {
	service.UnimplementedGetUserByIdServer
}

func (s *server) GetUserById(c context.Context, req *service.UserRequest) (*service.UserResponse, error) {
	return &service.UserResponse{Code: http.StatusOK, Msg: "成功", Data: "Rpc远程调用成功"}, nil
}

func main() {
	// 监听端口
	listen, _ := net.Listen("tcp", ":8888")

	// 创建grpc
	rpc := grpc.NewServer()

	// 注册服务
	service.RegisterGetUserByIdServer(rpc, &server{})

	// 启动服务
	err := rpc.Serve(listen)
	if err != nil {
		log.Println("启动服务失败")
	}

}
