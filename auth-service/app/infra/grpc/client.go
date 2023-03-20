package grpc

import (
	"auth-service/app/proto-gen/rpc"
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

type GrpcConfigClient struct {
	Host         string
	Port         int
	ReadTimwOut  int
	WriteTimeOut int
}
type AuthClient struct {
	example rpc.ExampleServiceClient
}

func InitGrpcClient(conf GrpcConfigClient) {
	doOne.Do(func() {
		err := initConn(conf)
		if err != nil {
			panic(err)
		}
		initClient(conn)
		loadGrpcConf(conf)
	})
}

func initConn(conf GrpcConfigClient) error {
	var err error
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", conf.Host, conf.Port),
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

func loadGrpcConf(conf GrpcConfigClient) {
	if conf.ReadTimwOut < 1000 || conf.WriteTimeOut < 1000 {
		panic(errors.New("This config of gRPC timeout is detrimental to working system"))
	}

	grpcReadTimeout = time.Duration(conf.ReadTimwOut) * time.Millisecond
	grpcWriteTimeout = time.Duration(conf.WriteTimeOut) * time.Millisecond

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
