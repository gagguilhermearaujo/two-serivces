package hashing

import (
	"crypto/sha256"
	"fmt"
)

func NewService() Service {
	return &service{
		hashes: map[string]string{},
	}
}

type Service interface {
	CheckHash(payload string) (hashExists bool)

	GetHash(payload string) (hashedString string, err error)

	CreateHash(payload string) (hashedString string, err error)
}

type service struct {
	hashes map[string]string
}

func (s *service) CheckHash(payload string) (hashExists bool) {
	_, hashExists = s.hashes[payload]
	return
}

func (s *service) GetHash(payload string) (hashedString string, err error) {
	hashedString, _ = s.hashes[payload]
	return
}

func (s *service) CreateHash(payload string) (hashedString string, err error) {
	hash := sha256.New()
	hash.Write([]byte(payload))
	hashedString = fmt.Sprintf("%x", hash.Sum(nil))
	s.hashes[payload] = hashedString
	return
}
