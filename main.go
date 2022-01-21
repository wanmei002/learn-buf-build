package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"gt/pb"
	"log"
	"net"
	"runtime/debug"
)

var (
	logger = logrus.New()
)

type ser struct {
	pb.UnimplementedFarmServiceServer
}

func (s *ser) AskPrice(ctx context.Context, in *pb.Pig) (*pb.RetPrice, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok != true {
		logrus.Debug("metadata get data failed;\n")
	}
	logrus.Printf("metadata get data: %+v\n", md)

	tagMap := grpc_ctxtags.Extract(ctx).Values()
	logrus.Printf("tag map is :%+v\n", tagMap)

	logger.Debugf("in val:%+v\n", *in)
	return &pb.RetPrice{
		TotalPrice: "1",
		UnitPrice:  "2",
	}, nil
}

func main() {
	ls, err := net.Listen("tcp", ":9989")
	if err != nil {
		log.Println("listen failed; err:", err)
		return
	}

	gs := grpc.NewServer(grpcSerOpt()...)
	pb.RegisterFarmServiceServer(gs, &ser{})
	reflection.Register(gs)
	log.Println("start grpc service addr: 9989")
	if err = gs.Serve(ls); err != nil {
		log.Println("service start failed; err:", err)
		return
	}
}

func grpcSerOpt() []grpc.ServerOption {

	logger.SetFormatter(new(logrus.TextFormatter))
	logger.SetLevel(logrus.DebugLevel)

	logrusEntry := logrus.NewEntry(logger)
	return []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			grpc_recovery.UnaryServerInterceptor(
				grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p interface{}) (err error) {
					debug.PrintStack()
					return status.Errorf(codes.Internal, "%v", p)
				})),
		),
	}
}
