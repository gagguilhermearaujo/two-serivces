package main

import (
	"log"
	"net"

	"github.com/gagguilhermearaujo/two-services/hashing"
	"google.golang.org/grpc"
)

func main() {
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	hashingService := hashing.NewService()
	hashingEndpoints := hashing.MakeEndpoints(hashingService)
	grpcServer := hashing.NewGrpcServer(hashingEndpoints)

	baseServer := grpc.NewServer()
	hashing.RegisterHashingServer(baseServer, grpcServer)
	baseServer.Serve(grpcListener)

}
