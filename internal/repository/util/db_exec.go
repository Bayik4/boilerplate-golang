package util

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DBExecutor struct {
	DB  *sqlx.DB
	Log *zap.Logger
}

func NewDBExecutor(db *sqlx.DB, log *zap.Logger) *DBExecutor {
	return &DBExecutor{
		DB:  db,
		Log: log,
	}
}

func (d *DBExecutor) Exec(ctx context.Context, query string, args ...any) error {

	_, err := d.DB.ExecContext(ctx, query, args...)
	if err != nil {
		d.Log.Error("failed execute query",
			zap.String("query", query),
			zap.Error(err),
		)
		return err
	}

	return nil
}

func (d *DBExecutor) Get(ctx context.Context, dest any, query string, args ...any) error {

	err := d.DB.GetContext(ctx, dest, query, args...)
	if err != nil {
		d.Log.Error("failed query get",
			zap.String("query", query),
			zap.Error(err),
		)
		return err
	}

	return nil
}

func (d *DBExecutor) Select(ctx context.Context, dest any, query string, args ...any) error {

	err := d.DB.SelectContext(ctx, dest, query, args...)
	if err != nil {
		d.Log.Error("failed query select",
			zap.String("query", query),
			zap.Error(err),
		)
		return err
	}

	return nil
}