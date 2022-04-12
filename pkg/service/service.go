package service

import "github.com/Ascemme/microservice.users.status/pkg/repository"

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}
