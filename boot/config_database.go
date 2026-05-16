package boot

import (
	"github.com/bayik4/boilerplate-golang/internal/model"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func LoadDatabaseConfig(vpr *viper.Viper, log *zap.Logger) model.DatabaseConfig {
	init := model.RDBMS{
		Host:     vpr.GetString("database.pgsql.db_name.host"),
		Port:     vpr.GetString("database.pgsql.db_name.port"),
		Username: vpr.GetString("database.pgsql.db_name.username"),
		Password: vpr.GetString("database.pgsql.db_name.password"),
		Dbname:   vpr.GetString("database.pgsql.db_name.dbname"),
	}

	db, err := PgSqlInit(init, log)
	if err != nil {
		log.Fatal("failed connect database", zap.Error(err))
	}

	init.Conn = db

	return model.DatabaseConfig{
		Pgsql: model.RDBMSItems{
			DBInit: init,
		},
	}
}