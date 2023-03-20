package client

import (
	"auth-component/app/proto-gen/rpc"
	conf "auth-service/app/config"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

var (
	doOne sync.Once
	conn  *grpc.ClientConn

	authGrpcClient *AuthClient

	grpcReadTimeout  time.Duration
	grpcWriteTimeout time.Duration
)

type AuthClient struct {
	example rpc.ExampleServiceClient
}

func InitGrpcClient() {
	doOne.Do(func() {
		err := initConn()
		if err != nil {
			panic(err)
		}
		initClient(conn)
		loadGrpcConf()
	})
}

func initConn() error {
	var err error
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", conf.Config.AuthGrpcConfig.Host, conf.Config.AuthGrpcConfig.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	return nil
}

func initClient(conn *grpc.ClientConn) {
	authGrpcClient = &AuthClient{
		example: rpc.NewExampleServiceClient(conn),
	}
}

func loadGrpcConf() {
	authConf := conf.Config.AuthGrpcConfig
	if authConf.ReadTimeout < 1000 || authConf.WriteTimeout < 1000 {
		panic(errors.New("This config of gRPC timeout is detrimental to working system"))
	}

	grpcReadTimeout = time.Duration(authConf.ReadTimeout) * time.Millisecond
	grpcWriteTimeout = time.Duration(authConf.WriteTimeout) * time.Millisecond

	fmt.Println("Grpc read timeout: ", grpcReadTimeout)
	fmt.Println("Grpc write timeout: ", grpcWriteTimeout)
}

func GetAuthGrpcReadTimeout() time.Duration {
	return grpcReadTimeout
}

func GetAuthGrpcWriteTimeout() time.Duration {
	return grpcWriteTimeout
}

func GrpcClient() *AuthClient {
	if authGrpcClient == nil {
		panic("cannot initial auth grpc client")
	}
	return authGrpcClient
}

func (u AuthClient) Hello() rpc.ExampleServiceClient {
	return u.example
}

func CloseGrpcClient() {
	if conn != nil {
		conn.Close()
	}
}
