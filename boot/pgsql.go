package boot

import (
	"fmt"

	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func PgSqlInit(cfg model.RDBMS, log *zap.Logger) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Dbname,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Error(
			"failed opening db connection",
			zap.String("dbname", cfg.Dbname),
			zap.Error(err),
		)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	db.SetMaxOpenConns(DBMaxPool)
	db.SetMaxIdleConns(DBMaxIdle)
	db.SetConnMaxLifetime(DBMaxLifeTime)

	log.Info("database connected", zap.String("dbname", cfg.Dbname))

	return db, nil
}
