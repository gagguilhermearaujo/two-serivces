package hashing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateHash endpoint.Endpoint
	CheckHash  endpoint.Endpoint
	GetHash    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateHash: makeCreateHashEndpoint(s),
		CheckHash:  makeCheckHashEndpoint(s),
		GetHash:    makeGetHashEndpoint(s),
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
		req := request.(endpointCreateHashRequest)
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

func makeGetHashEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(endpointGetHashRequest)
		hash, err := s.GetHash(req.Payload)
		return endpointGetHashResponse{Hash: hash}, err
	}
}

type endpointGetHashRequest struct {
	Payload string `json:"payload"`
}
type endpointGetHashResponse struct {
	Hash string `json:"hash,omitempty"`
}
