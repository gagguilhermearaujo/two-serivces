package hashing

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

func NewService() Service {
	return &service{
		hashes: map[string]string{},
	}
}

type Service interface {
	CheckHash(payload string) (hashExists bool, err error)
	GetHash(payload string) (hashedString string, err error)
	CreateHash(payload string) (hashedString string, err error)
}

type service struct {
	hashes map[string]string
}

func (s *service) CheckHash(payload string) (hashExists bool, err error) {
	_, hashExists = s.hashes[payload]
	return
}

func (s *service) GetHash(payload string) (hashedString string, err error) {
	hashedString, hashExists := s.hashes[payload]
	if !hashExists {
		err = errors.New("Hash does not exists")
	}
	return
}

func (s *service) CreateHash(payload string) (hashedString string, err error) {
	hash := sha256.New()
	hash.Write([]byte(payload))
	hashedString = fmt.Sprintf("%x", hash.Sum(nil))
	s.hashes[payload] = hashedString
	return
}
