package main

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	go main()

	tests := []struct {
		name         string
		path         string
		bodyPayload  map[string]any
		wantResponse map[string]any
	}{
		{
			name: "Integration test for CreateHash handler happy flow",
			path: "/CreateHash",
			bodyPayload: map[string]any{
				"payload": "kuririn",
			},
			wantResponse: map[string]any{
				"hash": "32ab96a3f54e7bd29187a448f6bbc87e8760194d55f66a73587687f0e8ac875b",
			},
		},
		{
			name: "Integration test for CheckHash handler true response",
			path: "/CheckHash",
			bodyPayload: map[string]any{
				"payload": "kuririn",
			},
			wantResponse: map[string]any{
				"hash_exists": true,
			},
		},
		{
			name: "Integration test for CheckHash handler false response",
			path: "/CheckHash",
			bodyPayload: map[string]any{
				"payload": "Vegeta",
			},
			wantResponse: map[string]any{
				"hash_exists": false,
			},
		},
		{
			name: "Integration test for GetHash handler happy flow",
			path: "/GetHash",
			bodyPayload: map[string]any{
				"payload": "kuririn",
			},
			wantResponse: map[string]any{
				"hash": "32ab96a3f54e7bd29187a448f6bbc87e8760194d55f66a73587687f0e8ac875b",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawResponse := testEndpoint(tt.path, tt.bodyPayload)
			if !reflect.DeepEqual(rawResponse, tt.wantResponse) {
				t.Errorf("%v failed, got rawResponse = %v, wantResponse %v", tt.name, rawResponse, tt.wantResponse)
			}
		})
	}
}

func testEndpoint(path string, bodyPayload map[string]any) map[string]any {
	client := &http.Client{}
	bytesPayload, err := json.Marshal(bodyPayload)
	if err != nil {
		panic(err)
	}
	var data = strings.NewReader(string(bytesPayload))
	req, err := http.NewRequest("POST", "http://localhost:3000"+path, data)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response map[string]any
	if err := json.Unmarshal(bodyText, &response); err != nil {
		panic(err)
	}

	return response
}
