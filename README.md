# Boilerplate Golang

A modular Go REST API starter built with Gin, Cobra, Viper, sqlx, PostgreSQL, Redis, Zap, and OpenTelemetry.

This project is designed as a clean foundation for backend services with a layered architecture and configuration-driven startup.

## Features

- Layered architecture: repository -> service -> handler -> router
- CLI entrypoint with Cobra
- Viper-based configuration loading from `config.yaml`
- PostgreSQL integration using `sqlx`
- Database migrations using `golang-migrate`
- Redis integration using `go-redis`
- Structured logging with `zap`
- Optional OpenTelemetry tracing support
- Example health-check endpoint

## Tech stack

- Go
- Gin HTTP framework
- Cobra CLI
- Viper configuration
- sqlx + PostgreSQL
- Redis
- OpenTelemetry
- Zap logger

## Project structure

```text
boilerplate-golang/
├── boot/                      # application bootstrap and infrastructure setup
│   ├── config.go
│   ├── config_app.go
│   ├── config_database.go
│   ├── config_logger.go
│   ├── config_redis.go
│   ├── config_telemetry.go
│   ├── logger.go
│   ├── pgsql.go
│   └── redis.go
├── cmd/                       # CLI commands
│   ├── rest.go
│   └── root.go
├── migrations/                # versioned SQL migration pairs
├── Makefile                   # database migration commands
├── internal/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   ├── router/
│   ├── service/
│   ├── telemetry/
│   └── utility/
├── config.yaml                # runtime configuration
├── go.mod
├── main.go
└── README.md
```

## Requirements

- Go version declared in `go.mod`
- PostgreSQL instance
- Redis instance
- `migrate` CLI from [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations
- Optional OTLP collector if OpenTelemetry is enabled

## Configuration

The application reads runtime settings from `config.yaml`.

Example configuration:

```yaml
app:
  name: boilerplate
  host: localhost
  port: 8080
  env: development

log:
  service_name: boilerplate
  level: debug
  format: json

otel:
  enabled: true
  exporter_otlp_endpoint: localhost:4317
  exporter_otlp_insecure: true
  service_name: boilerplate
  sampling_ratio: 1.0

database:
  pgsql:
    db_name:
      host: localhost
      port: 5432
      username: user_initial
      password: 123
      dbname: db_initial

redis:
  host: localhost
  port: 6379
  username: ""
  password: ""
  dbname: ""
  dbindex: 0

jwt:
  secret: ""
```

The configuration loaders live in the `boot` package, including:

- `boot/config.go`
- `boot/config_app.go`
- `boot/config_database.go`
- `boot/config_redis.go`
- `boot/config_logger.go`
- `boot/config_telemetry.go`

## Running the app

1. Update the values in `config.yaml` for your local environment.
2. Ensure PostgreSQL and Redis are available.
3. Start the server:

```bash
go run main.go serve
```

This invokes the `serve` command defined in `cmd/rest.go`.

## Database migrations

Database schema changes are managed with `golang-migrate`. Each migration is a
pair of numbered `.up.sql` and `.down.sql` files in `migrations/`. The
repository includes an initial migration that creates `initial_table`.

Install the `migrate` CLI before running the migration targets:

```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.20.1
```

The Makefile defaults to the local PostgreSQL database. Override
`DATABASE_URL` to target another database:

```bash
export DATABASE_URL='postgres://DB_USER:DB_PASSWORD@DB_HOST:5432/DB_NAME?sslmode=disable'
```

Create a new sequential migration pair:

```bash
make migrate-create name=add_example_table
```

Apply all pending migrations:

```bash
make migrate-up
```

Roll back the most recent migration (one version):

```bash
make migrate-down
```

Check the current migration version:

```bash
make migrate-version
```

Force-set the recorded migration version:

```bash
make migrate-force version=1
```

`migrate-force` changes the migration version without applying or reverting
SQL, so use it only to recover a migration state intentionally. Migrations are
run explicitly with these Make targets; starting the REST server does not
automatically apply them.

## Application startup flow

The startup sequence is roughly:

1. Initialize logger
2. Load application configuration
3. Initialize OpenTelemetry if enabled
4. Initialize PostgreSQL connection
5. Initialize Redis connection
6. Build repository, service, and handler layers
7. Register router and start HTTP server

## API

The app includes a health check endpoint.

- GET `/api/v1/health/`

Example:

```bash
curl http://localhost:8080/api/v1/health/
```

## Notes

- Entry point: `main.go`
- CLI root command: `cmd/root.go`
- Server command: `cmd/rest.go`
- Router registration: `internal/router`
- Health implementation: `internal/handler/health.go`, `internal/service/health.go`

This repository is meant to be a starting point for building a production-ready Go API with common infrastructure already in place.