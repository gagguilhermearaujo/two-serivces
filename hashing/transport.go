package hashing

import (
	"context"
	"errors"

	"github.com/gagguilhermearaujo/two-services/hashing/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	createHash grpctransport.Handler
	checkHash  grpctransport.Handler
	getHash    grpctransport.Handler
	pb.UnimplementedHashingServer
}

func (s gRPCServer) CreateHash(ctx context.Context, req *pb.CreateHashRequest) (*pb.CreateHashResponse, error) {
	_, resp, err := s.createHash.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.CreateHashResponse), nil
}

func (s gRPCServer) CheckHash(ctx context.Context, req *pb.CheckHashRequest) (*pb.CheckHashResponse, error) {
	_, resp, err := s.checkHash.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.CheckHashResponse), nil
}

func (s gRPCServer) GetHash(ctx context.Context, req *pb.GetHashRequest) (*pb.GetHashResponse, error) {
	_, resp, err := s.getHash.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GetHashResponse), nil
}

func NewGrpcServer(endpoints Endpoints) pb.HashingServer {
	return &gRPCServer{
		createHash: grpctransport.NewServer(
			endpoints.CreateHash,
			decodeCreateHashRequest,
			encodeCreateHashResponse,
		),
		checkHash: grpctransport.NewServer(
			endpoints.CheckHash,
			decodeCheckHashRequest,
			encodeCheckHashResponse,
		),
		getHash: grpctransport.NewServer(
			endpoints.GetHash,
			decodeGetHashRequest,
			encodeGetHashResponse,
		),
	}
}

func decodeCreateHashRequest(ctx context.Context, request any) (any, error) {
	req, ok := request.(*pb.CreateHashRequest)
	if !ok {
		return nil, errors.New("invalid request body")
	}

	return endpointCreateHashRequest{Payload: req.Payload}, nil
}

func encodeCreateHashResponse(ctx context.Context, response any) (any, error) {
	res, ok := response.(endpointCreateHashResponse)
	if !ok {
		return nil, errors.New("invalid response body")
	}

	return &pb.CreateHashResponse{Hash: res.Hash}, nil
}

func decodeCheckHashRequest(ctx context.Context, request any) (any, error) {
	req, ok := request.(*pb.CheckHashRequest)
	if !ok {
		return nil, errors.New("invalid request body")
	}

	return endpointCheckHashRequest{Payload: req.Payload}, nil
}

func encodeCheckHashResponse(ctx context.Context, response any) (any, error) {
	res, ok := response.(endpointCheckHashResponse)
	if !ok {
		return nil, errors.New("invalid response body")
	}

	return &pb.CheckHashResponse{HashExists: res.HashExists}, nil
}

func decodeGetHashRequest(ctx context.Context, request any) (any, error) {
	req, ok := request.(*pb.GetHashRequest)
	if !ok {
		return nil, errors.New("invalid request body")
	}

	return endpointGetHashRequest{Payload: req.Payload}, nil
}

func encodeGetHashResponse(ctx context.Context, response any) (any, error) {
	res, ok := response.(endpointGetHashResponse)
	if !ok {
		return nil, errors.New("invalid response body")
	}

	return &pb.GetHashResponse{Hash: res.Hash}, nil
}
