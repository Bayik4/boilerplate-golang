package repository

import (
	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/bayik4/boilerplate-golang/internal/repository/health"
	"github.com/bayik4/boilerplate-golang/internal/repository/health/database"
	"github.com/bayik4/boilerplate-golang/internal/repository/util"
)

type Repository struct {
	HealthRepository health.HealthRepository
}

func New(db model.RDBMS, dbexec util.DBExecutor) *Repository {
	return &Repository{
		HealthRepository: database.NewHealthRepository(db, dbexec),
	}
}