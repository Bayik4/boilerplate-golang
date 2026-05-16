package database

import (
	"context"

	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/bayik4/boilerplate-golang/internal/repository/util"
)

type HealthRepository struct {
	db     model.RDBMS
	dbexec util.DBExecutor
}

func NewHealthRepository(db model.RDBMS, dbexec util.DBExecutor) *HealthRepository {
	return &HealthRepository{
		db: db,
		dbexec: dbexec,
	}
}

func (r *HealthRepository) Get(ctx context.Context) (string, error) {
	_ = r.dbexec.Exec(ctx, r.db.Conn, QueryHealth)
	return "ok", nil
}
