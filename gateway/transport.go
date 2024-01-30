package hashing

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHandler(s Service) {
	createHashHandler := httptransport.NewServer(
		makeCreateHashEndpoint(s),
		decodeCreateHashRequest,
		encodeResponse,
	)
	http.Handle("/CreateHash", createHashHandler)
	http.ListenAndServe(":8080", nil)
}

func decodeCreateHashRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createHashRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
