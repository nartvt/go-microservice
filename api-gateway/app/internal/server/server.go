package main

import (
	"api-gateway/app/config"
	"api-gateway/app/infra/db"
	"api-gateway/app/middleware"
	"api-gateway/app/router"
	"auth-service/app/infra/grpc"
	client "auth-service/app/transport/grpc"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	setupInfra()
	app := fiber.New(middleware.Config())
	defer closeInfra(app)
	initRouter(app)
	run(app)
}

func run(app *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	timeOut := time.Duration(config.Get().Server.ShutdownTimeOut) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	go func() {
		address := fmt.Sprintf("%s:%d", config.Get().Server.Host, config.Get().Server.Port)
		err := app.Listen(address)
		if err != nil {
			log.Println("Application startup interrupt", err.Error())
			cancel()
		}
	}()

	select {
	case sig := <-c:
		fmt.Printf("Received signal: %v\n", sig)
	case <-ctx.Done():
		fmt.Println("Server timed out, shutting down...")
	}

	time.Sleep(time.Second)
	log.Println("Shutdown complete.")
}

func setupInfra() {
	setupDatabase()
	setupAuthGrpcClient()
}

func setupAuthGrpcClient() {
	// Every client connect to grpc server
	// must build their own configuration themselves.
	conf := config.Get()
	grpcConf := grpc.GrpcConfigClient{
		Host:         conf.AuthGrpcConfig.Host,
		Port:         conf.AuthGrpcConfig.Port,
		ReadTimwOut:  conf.AuthGrpcConfig.ReadTimeout,
		WriteTimeOut: conf.AuthGrpcConfig.WriteTimeout,
	}
	grpc.InitGrpcClient(grpcConf)
	pingGrpc()
}

func pingGrpc() {
	if _, err := client.Hello.Ping(); err != nil {
		log.Println(err.Error())
		panic(err)
	}
}
func closeInfra(app *fiber.App) {
	db.ClosePostgres()
	fmt.Println("Shutting down server...")
	if err := app.ShutdownWithTimeout(time.Duration(config.Get().Server.ShutdownTimeOut)); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}

}

func setupDatabase() {
	db.InitPostgres()
}
func initRouter(app *fiber.App) {
	router.SetupRoutes(app)
}
