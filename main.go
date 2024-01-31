package main

import (
	"log"
	"net"

	"github.com/gagguilhermearaujo/two-services/gateway"
	"github.com/gagguilhermearaujo/two-services/hashing"
	"google.golang.org/grpc"
)

func main() {
	hashingListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen hashing service: %v", err)
	}

	hashingService := hashing.NewService()
	hashingEndpoints := hashing.MakeEndpoints(hashingService)
	grpcServer := hashing.NewGrpcServer(hashingEndpoints)

	baseServer := grpc.NewServer()
	hashing.RegisterHashingServer(baseServer, grpcServer)
	go baseServer.Serve(hashingListener)

	gatewayServer := gateway.NewGatewayServer()
	gatewayServer.MakeEndpoints()
	gatewayServer.Serve()

}
