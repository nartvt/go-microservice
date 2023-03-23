package grpc

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"product-service/app/transport/proto-gen/rpc"
	"sync"
	"time"
)

var (
	doOne sync.Once
	conn  *grpc.ClientConn

	productGrpcClient *ProductClient

	grpcReadTimeout  time.Duration
	grpcWriteTimeout time.Duration
)

type ConfigClient struct {
	Host         string
	Port         int
	ReadTimeOut  int
	WriteTimeOut int
}
type ProductClient struct {
	productClient rpc.ProductServiceClient
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
	productGrpcClient = &ProductClient{
		productClient: rpc.NewProductServiceClient(conn),
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

func GetProductGrpcReadTimeout() time.Duration {
	return grpcReadTimeout
}

func GetProductGrpcWriteTimeout() time.Duration {
	return grpcWriteTimeout
}

func GetGrpcClient() *ProductClient {
	if productGrpcClient == nil {
		panic("cannot initial auth grpc client")
	}
	return productGrpcClient
}

func (u ProductClient) ProductService() rpc.ProductServiceClient {
	return u.productClient
}

func CloseGrpcClient() {
	if conn != nil {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}
}
