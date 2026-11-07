package repository

import (
	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/bayik4/boilerplate-golang/internal/repository/health"
	"github.com/bayik4/boilerplate-golang/internal/repository/health/database"
	"github.com/bayik4/boilerplate-golang/internal/repository/util"
	"go.uber.org/zap"
)

type Repository struct {
	HealthRepository health.HealthRepository
	logger           *zap.Logger
}

func New(db model.RDBMS, dbexec util.DBExecutor, logger *zap.Logger) *Repository {
	return &Repository{
		HealthRepository: database.NewHealthRepository(db, dbexec),
		logger:           logger,
	}
}
