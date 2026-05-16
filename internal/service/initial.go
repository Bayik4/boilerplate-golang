package service

import "github.com/bayik4/boilerplate-golang/internal/repository"

type Service struct {
	HealthService *HealthService
}

func New(repo *repository.Repository) *Service {
	return &Service{
		HealthService: newHealthService(repo.HealthRepository),
	}
}
