package hashing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type createHashRequest struct {
	Payload string `json:"payload"`
}

type createHashRespose struct {
	Hash string `json:"hash,omitempty"`
	Err  error  `json:"error,omitempty"`
}

func makeCreateHashEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(createHashRequest)
		hash, err := s.CreateHash(req.Payload)
		return createHashRespose{Hash: hash, Err: err}, nil
	}
}
