# Boilerplate Golang

Boilerplate API server menggunakan Gin, sqlx (Postgres) dan Redis. Struktur proyek modular untuk memudahkan pengembangan layanan REST.

## Fitur utama
- Struktur layered: repository -> service -> handler -> router
- Inisialisasi konfigurasi dari file [config.yaml](config.yaml)
- Postgres connection via [`boot.PgSqlInit`](boot/pgsql.go)
- Redis connection via [`boot.RedisInit`](boot/redis.go)
- Logging dengan Uber Zap via [`boot.LoggerInit`](boot/logger.go)

## Ringkasan eksekusi
Entrypoint: [main.go](main.go) yang memanggil [`cmd.Execute`](cmd/root.go). Perintah utama untuk menjalankan server adalah subcommand `serve` yang didefinisikan di [`cmd/rest.go`](cmd/rest.go) (`cmd.restCommand`).

Alur inisialisasi singkat:
- Logger: [`boot.LoggerInit`](boot/logger.go)  
- Load konfigurasi: [`boot.LoadConfig`](boot/config.go) (membaca [config.yaml](config.yaml))  
- Database: [`boot.PgSqlInit`](boot/pgsql.go) dan konfigurasi di [`boot/config_database.go`](boot/config_database.go)  
- Redis: [`boot.RedisInit`](boot/redis.go) dan konfigurasi di [`boot/config_redis.go`](boot/config_redis.go)  
- Repository: [`repository.New`](internal/repository/initial.go)  
- Service: [`service.New`](internal/service/initial.go)  
- Handler: [`handler.New`](internal/handler/initial.go)  
- Router: [`router.NewRouter`](internal/router/initial.go)

## Konfigurasi
Sesuaikan nilai di [config.yaml](config.yaml):
- app.port untuk port server
- database.pgsql.* untuk koneksi Postgres
- redis.* untuk koneksi Redis
- jwt.secret untuk secret JWT (jika diperlukan)

Loader konfigurasi dan inisialisasi terkait ada di:
- [`boot/config.go`](boot/config.go)
- [`boot/config_app.go`](boot/config_app.go)
- [`boot/config_database.go`](boot/config_database.go)
- [`boot/config_redis.go`](boot/config_redis.go)

## Struktur penting (ringkasan)
- [main.go](main.go) — entry.
- [cmd/root.go](cmd/root.go), [cmd/rest.go](cmd/rest.go) — CLI & server start.
- [boot/logger.go](boot/logger.go), [boot/pgsql.go](boot/pgsql.go), [boot/redis.go](boot/redis.go) — init infra.
- [internal/model/config.go](internal/model/config.go), [config.yaml](config.yaml) — model dan konfigurasi.
- [internal/repository/initial.go](internal/repository/initial.go), [internal/repository/util/db_exec.go](internal/repository/util/db_exec.go) — repositori & helper DB.
- [internal/repository/health/database/impl.go](internal/repository/health/database/impl.go) — contoh implementasi repository.
- [internal/service/initial.go](internal/service/initial.go), [internal/service/health.go](internal/service/health.go) — service layer.
- [internal/handler/initial.go](internal/handler/initial.go), [internal/handler/health.go](internal/handler/health.go) — HTTP handler.
- [internal/router/initial.go](internal/router/initial.go), [internal/router/api_health.go](internal/router/api_health.go) — routing.
- [internal/utility/response.go](internal/utility/response.go) — response helper (contoh: `utility.ResponseWithoutData`).

## Endpoint contoh
- GET /api/v1/health/ — health check  
  Handler: [`handler.HandlerHealth.CheckHealth`](internal/handler/health.go)  
  Service: [`service.HealthService.GetHealth`](internal/service/health.go)  
  Response helper: [`utility.ResponseWithoutData`](internal/utility/response.go)

## Menjalankan secara lokal
1. Pastikan Go terinstal (lihat [go.mod](go.mod)).
2. Sesuaikan [config.yaml](config.yaml).
3. Jalankan:
```sh
go run main.go serve
```