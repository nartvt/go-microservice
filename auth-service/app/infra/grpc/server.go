package grpc

import (
	"auth-service/app/config"
	"auth-service/app/domain/usercases/user/repo"
	"auth-service/app/proto-gen/rpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitGrpcServer() {
	conf := config.Get()
	address := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	var (
		userGrpcServer = repo.User
		roleGrpcServer = repo.RoleRepo
	)

	ops := []grpc.ServerOption{
		grpc.ConnectionTimeout(time.Duration(conf.Server.ConnectTimeOut) * time.Second), // set connection timeout
		grpc.UnaryInterceptor(interceptor),                                              // set unary interceptor
	}

	grpcServer := grpc.NewServer(ops...)
	rpc.RegisterUserServiceServer(grpcServer, userGrpcServer)
	rpc.RegisterRoleServiceServer(grpcServer, roleGrpcServer)
	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Printf("Listening on %v", listen.Addr())
	<-shutdownChan // wait for the shutdown signal
	log.Println("Shutting down gRPC server...")

	grpcServer.GracefulStop()
	log.Println("gRPC server gracefully stopped.")
}
func interceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	log.Printf("received request: %v", req)
	log.Println(info)
	resp, err := handler(ctx, req)
	log.Printf("sent response: %v", resp)
	return resp, err
}
