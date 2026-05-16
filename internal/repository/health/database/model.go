package database

import "database/sql"

type HealthDB struct {
	Status sql.NullString `db:"status"`
}
