package service

import "github.com/bitstored/file-service/pkg/repo"

type Service struct {
	Repo repo.Repository
}

func NewService() *Service {
	return &Service{repo.FileRepository()}
}
