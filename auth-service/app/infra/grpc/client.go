package grpc

import (
	"auth-service/app/proto-gen/rpc"
	"errors"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	doOne sync.Once
	conn  *grpc.ClientConn

	authGrpcClient *AuthClient

	grpcReadTimeout  time.Duration
	grpcWriteTimeout time.Duration
)

type ConfigClient struct {
	Host         string
	Port         int
	ReadTimeOut  int
	WriteTimeOut int
}
type AuthClient struct {
	userClient rpc.UserServiceClient
	roleClient rpc.RoleServiceClient
}

func InitGrpcClient(conf ConfigClient) {
	doOne.Do(func() {
		err := initConn(conf)
		if err != nil {
			panic(err)
		}
		initClient(conn)
		loadGrpcConf(conf)
	})
}

func initConn(conf ConfigClient) error {
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
		userClient: rpc.NewUserServiceClient(conn),
		roleClient: rpc.NewRoleServiceClient(conn),
	}
}

func loadGrpcConf(conf ConfigClient) {
	if conf.ReadTimeOut < 1000 || conf.WriteTimeOut < 1000 {
		panic(errors.New("this config of gRPC timeout is detrimental to working system"))
	}

	grpcReadTimeout = time.Duration(conf.ReadTimeOut) * time.Millisecond
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

func GetGrpcClient() *AuthClient {
	if authGrpcClient == nil {
		panic("cannot initial auth grpc client")
	}
	return authGrpcClient
}

func (u AuthClient) UserClient() rpc.UserServiceClient {
	return u.userClient
}

func (u AuthClient) RoleClient() rpc.RoleServiceClient {
	return u.roleClient
}

func CloseGrpcClient() {
	if conn != nil {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}
}
