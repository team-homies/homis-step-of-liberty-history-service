package app

import (
	"main/app/api"
	g "main/app/grpc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/grpc"
)

// fiber 인스턴스 생성
func InitFiber() *fiber.App {
	app := fiber.New()
	app.Use(cors.New()) // cors미들웨어를 쓴다는 의미
	api.InitRoutes(app)

	return app
}

// go함수
func InitGrpc() *grpc.Server {
	grpcServer := grpc.NewServer()
	g.RegisterServices(grpcServer)
	return grpcServer
}
