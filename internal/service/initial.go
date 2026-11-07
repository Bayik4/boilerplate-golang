package service

import (
	"github.com/bayik4/boilerplate-golang/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	HealthService *HealthService
	logger        *zap.Logger
}

func New(repo *repository.Repository, logger *zap.Logger) *Service {
	return &Service{
		HealthService: newHealthService(repo.HealthRepository),
		logger:        logger,
	}
}
