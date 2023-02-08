package main

import (
	"context"
	service "day4/Grpc/hello-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	connect, _ := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer connect.Close()

	client := service.NewGetUserByIdClient(connect)

	resp, _ := client.GetUserById(context.Background(), &service.UserRequest{Username: "方糖", Address: "清华大学", Age: 22})
	log.Println(resp.GetData())

}
