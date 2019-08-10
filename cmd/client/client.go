package main

import (
	"context"
	"demo/conf"
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"

	pb "demo/pb/user"
)

func init() {
	// flags
	flag.StringVar(&conf.GRPCAddr, "grpc-addr", conf.GetEnv("GRPCAddr", "0.0.0.0:5000"), "grpc服务地址")
}

func main() {
	conn, err := grpc.Dial(conf.GRPCAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("创建grpc连接失败! Error: %s", err)
	}
	defer conn.Close()

	Create(conn)
	Delete(conn)
}

func Create(conn *grpc.ClientConn) {
	in := &pb.CreateReq{
		Name: "wss",
		Age:  19,
	}

	out := &pb.CreateResp{}
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := conn.Invoke(ctx, "/pb.User/Create", in, out)

	fmt.Println(err)
	fmt.Println(out)
}
func Delete(conn *grpc.ClientConn) {
	in := &pb.DeleteReq{
		Name: "wss",
		Id:   19,
	}

	out := &pb.DeleteResp{}
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := conn.Invoke(ctx, "/pb.User/Delete", in, out)
	fmt.Println(err)
	fmt.Println(out)
}
