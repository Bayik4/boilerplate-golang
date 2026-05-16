package service

import (
	"context"

	"github.com/bayik4/boilerplate-golang/internal/repository/health"
)

type HealthService struct {
	repo health.HealthRepository
}

func newHealthService(repo health.HealthRepository) *HealthService {
	return &HealthService{
		repo: repo,
	}
}

func (s *HealthService) GetHealth(ctx context.Context) (string, error) {
	return s.repo.Get(ctx)
}