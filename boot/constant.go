package boot

import "time"

var (
	DBMaxPool     int           = 25
	DBMaxIdle     int           = 10
	DBMaxLifeTime time.Duration = time.Minute * 5
)
