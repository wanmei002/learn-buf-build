package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gt/greateway/pb"
	"log"
	"net"
    "net/http"
)

type ser struct {
	pb.UnsafePhoneServiceServer
}

func (s *ser) DialRequest(ctx context.Context, in *pb.Dial) (*pb.Ask, error) {
	return &pb.Ask{
		Code: 0,
		Msg:  "success",
		Data: &pb.AskContext{
			Hostname: in.Hostname,
			Port:     in.Port,
			Name:     in.Name,
			CanCall:  true,
		},
	}, nil
}

func main() {
	ls, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Println("listen 8087 failed; err:", err)
		return
	}

	gs := grpc.NewServer()
	pb.RegisterPhoneServiceServer(gs, &ser{})
	// 在这里把 微服务启动起来
	go func() {
		log.Println("start server : 8087")
		log.Fatal(gs.Serve(ls))
	}()
	// grpc客户端, 跟grpc服务通信
	dl, err := grpc.Dial(
		":8087",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("grpc dial failed; err:", err)
	}
	// 把在 proto中声明的路由注册到多路路由中
	mux := runtime.NewServeMux()
	err = pb.RegisterPhoneServiceHandler(context.Background(), mux, dl)
    if err != nil {
        log.Fatalln("mux failed; err: ", err)
        return
    }
	// 启动 gateway 服务
    httpService := http.Server{
        Addr:              ":8091",
        Handler:           mux,
    }
	log.Println("http service 8091")
    log.Fatalln(httpService.ListenAndServe())
}
