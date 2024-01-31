package hashing

import (
	"testing"
)

func Test_service_CreateHash(t *testing.T) {
	type fields struct {
		hashes map[string]string
	}
	type args struct {
		payload string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantHashedString string
		wantErr          bool
	}{
		{
			name:             "Happy flow",
			fields:           fields{hashes: map[string]string{}},
			args:             args{payload: "Vegeta prince of saiyans"},
			wantHashedString: "123182aa9fc7f51d3e3c13ff9f85b7df0b107adce3c58813b976c2a6005833eb",
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				hashes: tt.fields.hashes,
			}
			gotHashedString, err := s.CreateHash(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHashedString != tt.wantHashedString {
				t.Errorf("service.CreateHash() = %v, want %v", gotHashedString, tt.wantHashedString)
			}
		})
	}
}

func Test_service_GetHash(t *testing.T) {
	type fields struct {
		hashes map[string]string
	}
	type args struct {
		payload string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantHashedString string
		wantErr          bool
	}{
		{
			name: "Happy flow",
			fields: fields{hashes: map[string]string{
				"Vegeta prince of saiyans": "123182aa9fc7f51d3e3c13ff9f85b7df0b107adce3c58813b976c2a6005833eb",
			}},
			args:             args{payload: "Vegeta prince of saiyans"},
			wantHashedString: "123182aa9fc7f51d3e3c13ff9f85b7df0b107adce3c58813b976c2a6005833eb",
			wantErr:          false,
		},
		{
			name:             "Nonexistent hash flow",
			fields:           fields{hashes: map[string]string{}},
			args:             args{payload: "Vegeta prince of saiyans"},
			wantHashedString: "",
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				hashes: tt.fields.hashes,
			}
			gotHashedString, err := s.GetHash(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHashedString != tt.wantHashedString {
				t.Errorf("service.GetHash() = %v, want %v", gotHashedString, tt.wantHashedString)
			}
		})
	}
}

func Test_service_CheckHash(t *testing.T) {
	type fields struct {
		hashes map[string]string
	}
	type args struct {
		payload string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantHashExists bool
		wantErr        bool
	}{
		{
			name: "Happy flow true",
			fields: fields{hashes: map[string]string{
				"Vegeta prince of saiyans": "123182aa9fc7f51d3e3c13ff9f85b7df0b107adce3c58813b976c2a6005833eb",
			}},
			args:           args{payload: "Vegeta prince of saiyans"},
			wantHashExists: true,
			wantErr:        false,
		},
		{
			name:           "Happy flow false",
			fields:         fields{hashes: map[string]string{}},
			args:           args{payload: "Vegeta prince of saiyans"},
			wantHashExists: false,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				hashes: tt.fields.hashes,
			}
			gotHashExists, err := s.CheckHash(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CheckHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHashExists != tt.wantHashExists {
				t.Errorf("service.CheckHash() = %v, want %v", gotHashExists, tt.wantHashExists)
			}
		})
	}
}
