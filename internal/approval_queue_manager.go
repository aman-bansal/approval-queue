package internal

import (
	"context"
	"github.com/aman-bansal/approval-queue/config"
	"github.com/aman-bansal/approval-queue/internal/impl"
	protoimpl "github.com/aman-bansal/approval-queue/pkg/grpc"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
)

func Init(ctx context.Context, appEnv string) error {
	conf, err := config.Get("test");
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcmiddleware.ChainUnaryServer(
				//todo add interceptor for authorization and authentication
			),
		),
	)

	controller := impl.NewApprovalQueueManager(conf)
	protoimpl.RegisterApprovalQueueManagerServer(grpcServer, controller)
	lis, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		return err
	}

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
