package util

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DBExecutor struct {
	Log *zap.Logger
}

func NewDBExecutor(log *zap.Logger) *DBExecutor {
	return &DBExecutor{
		Log: log,
	}
}

func (d *DBExecutor) Exec(
	ctx context.Context,
	db sqlx.ExtContext,
	query string,
	args ...any,
) error {

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		d.Log.Error(
			"failed execute query",
			zap.String("query", query),
			zap.Error(err),
		)
		return err
	}

	return nil
}

func (d *DBExecutor) Get(
	ctx context.Context,
	db sqlx.QueryerContext,
	dest any,
	query string,
	args ...any,
) error {

	err := sqlx.GetContext(
		ctx,
		db,
		dest,
		query,
		args...,
	)

	if err != nil {
		d.Log.Error(
			"failed query get",
			zap.String("query", query),
			zap.Error(err),
		)
		return err
	}

	return nil
}

func (d *DBExecutor) Select(
	ctx context.Context,
	db sqlx.QueryerContext,
	dest any,
	query string,
	args ...any,
) error {

	err := sqlx.SelectContext(
		ctx,
		db,
		dest,
		query,
		args...,
	)

	if err != nil {
		d.Log.Error(
			"failed query select",
			zap.String("query", query),
			zap.Error(err),
		)
		return err
	}

	return nil
}