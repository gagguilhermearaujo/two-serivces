package hashing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateHash endpoint.Endpoint
	CheckHash  endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateHash: makeCreateHashEndpoint(s),
		CheckHash:  makeCheckHashEndpoint(s),
	}
}

func makeCreateHashEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(endpointCreateHashRequest)
		hash, err := s.CreateHash(req.Payload)
		return endpointCreateHashResponse{Hash: hash}, err
	}
}

type endpointCreateHashRequest struct {
	Payload string `json:"payload"`
}
type endpointCreateHashResponse struct {
	Hash string `json:"hash,omitempty"`
}

func makeCheckHashEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(endpointCheckHashRequest)
		hashExists, err := s.CheckHash(req.Payload)
		return endpointCheckHashResponse{HashExists: hashExists}, err
	}
}

type endpointCheckHashRequest struct {
	Payload string `json:"payload"`
}
type endpointCheckHashResponse struct {
	HashExists bool `json:"hash_exists,omitempty"`
}
