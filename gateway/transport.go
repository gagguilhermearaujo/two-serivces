package gateway

import (
	"log"

	"github.com/gagguilhermearaujo/two-services/hashing"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGatewayServer() *GatewayServer {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to hashing service: %v", err)
	}

	return &GatewayServer{
		fiberApp:       fiber.New(),
		hashingService: hashing.NewHashingClient(conn),
	}
}

type GatewayServer struct {
	fiberApp       *fiber.App
	hashingService hashing.HashingClient
}

func (s GatewayServer) Serve() {
	s.fiberApp.Listen(":3000")
}

func (s *GatewayServer) MakeEndpoints() {
	makeCreateHashEndpoint(s)
	makeCheckHashEndpoint(s)
	makeGetHashEndpoint(s)
}
