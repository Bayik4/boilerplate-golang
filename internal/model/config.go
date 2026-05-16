package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type ConfigModel struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
}

type AppConfig struct {
	Name string
	Host string
	Port string
}

type DatabaseConfig struct {
	Pgsql RDBMSItems
}

type RDBMSItems struct {
	DBInit RDBMS
}

type RDBMS struct {
	Host      string
	Port      string
	Username  string
	Password  string
	Dbname    string
	Conn *sqlx.DB
}

type RedisConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	DbIndex  int
	Conn     *redis.Client
}

type JWTConfig struct {
	SecretKey string
}
