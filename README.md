# Boilerplate Golang

Ringkasan  
Boilerplate ini menyediakan starting point untuk membuat REST API di Go menggunakan:
- Gin (HTTP router)
- Cobra (CLI)
- Viper (konfigurasi)
- Zap (logging)
- PostgreSQL (sqlx) sebagai RDBMS
- Redis (go-redis) sebagai cache/kv
- Struktur modular untuk router, handler, repository, dan utilitas

Prasyarat
- Go (lihat go.mod)
- PostgreSQL (opsional, jika ingin koneksi DB)
- Redis (opsional)
- Make/CLI dasar (opsional)

Instalasi & menjalankan
1. Clone repo:
```sh
git clone <repo-url>
cd boilerplate-golang
```

2. Jalankan tanpa build:
```sh
go run main.go serve
```

3. Build lalu jalankan:
```sh
go build -o app .
./app serve
```

Perintah `serve` didefinisikan di `cmd/rest.go` dan didaftarkan pada root command di `cmd/root.go`.

Konfigurasi
File konfigurasi utama: `config.yaml` (dibaca dengan Viper). Fungsi pembaca config ada di folder `boot/`:
- `boot.LoadConfig`
- `boot.LoadAppConfig`
- `boot.LoadDatabaseConfig`
- `boot.LoadRedisConfig`

Model konfigurasi: `internal/model/config.go` (struct ConfigModel, RDBMS, dsb).

Catatan: Periksa konsistensi key di file config dan loader. Ada potensi mismatch antara key `database.pgsql.db_name.*` vs `database.pgsql.db_init.*`.

Struktur proyek (ringkas)
- main.go — entry point aplikasi
- cmd/
  - root.go — root Cobra command
  - rest.go — command untuk serve
- boot/ — inisialisasi konfigurasi, logger, DB, Redis
- internal/
  - router/ — pembuatan router Gin dan pendaftaran route
  - handler/ — handler HTTP (mis. health)
  - middleware/ — middleware (mis. logging)
  - repository/ — util DB dan repositori
  - utility/ — helper response, dsb
  - model/ — model konfigurasi & response

Endpoint contoh
- Health check
  - Method: GET
  - URL: /api/v1/health/
  - Contoh:
  ```sh
  curl -v http://localhost:8080/api/v1/health/
  ```

Logging & observabilitas
- Logger: inisialisasi di `boot/logger.go` (zap).
- Middleware logging: `internal/middleware/logger.go`.
- Periksa output logger saat gagal load config atau gagal koneksi DB/Redis.

Database & Redis
- Inisialisasi PostgreSQL: `boot/PgSqlInit` (mengembalikan *sqlx.DB).
- Inisialisasi Redis: `boot/RedisInit` (mengembalikan *redis.Client).
- DB helper: `internal/repository/util/db_exec.go` (Exec/Get/Select wrappers).

Testing
- Tambahkan unit test di package terkait (`*_test.go`).  
- Jalankan:
```sh
go test ./...
```

Debugging cepat
- Periksa log (stdout) untuk pesan fatal dari boot atau koneksi.
- Pastikan `config.yaml` sesuai environment (path key, credential).
- Jika menggunakan Docker / container, pastikan network DB/Redis reachable.

File penting
- main.go
- cmd/root.go, cmd/rest.go
- boot/*.go
- internal/router/*, internal/handler/*, internal/repository/*, internal/utility/*, internal/model/*
- config.yaml
- go.mod

Kontribusi
- Ikuti gaya kode Go (gofmt/golint).
- Buat branch per fitur/bugfix dan ajukan PR.

Lisensi
- Tambahkan file LICENSE sesuai kebutuhan proyek.

Jika ingin, saya bisa:
- Buat README versi Bahasa Inggris,
- Perbarui config.yaml contoh, atau
- Buat badge/CI minimal (GitHub Actions) untuk build & test.
```// filepath: /home/enigma/projects/boilerplate-golang/readme.md
// ...existing code...
# Boilerplate Golang

Ringkasan  
Boilerplate ini menyediakan starting point untuk membuat REST API di Go menggunakan:
- Gin (HTTP router)
- Cobra (CLI)
- Viper (konfigurasi)
- Zap (logging)
- PostgreSQL (sqlx) sebagai RDBMS
- Redis (go-redis) sebagai cache/kv
- Struktur modular untuk router, handler, repository, dan utilitas

Prasyarat
- Go (lihat go.mod)
- PostgreSQL (opsional, jika ingin koneksi DB)
- Redis (opsional)
- Make/CLI dasar (opsional)

Instalasi & menjalankan
1. Clone repo:
```sh
git clone <repo-url>
cd boilerplate-golang
```

2. Jalankan tanpa build:
```sh
go run main.go serve
```

3. Build lalu jalankan:
```sh
go build -o app .
./app serve
```

Perintah `serve` didefinisikan di `cmd/rest.go` dan didaftarkan pada root command di `cmd/root.go`.

Konfigurasi
File konfigurasi utama: `config.yaml` (dibaca dengan Viper). Fungsi pembaca config ada di folder `boot/`:
- `boot.LoadConfig`
- `boot.LoadAppConfig`
- `boot.LoadDatabaseConfig`
- `boot.LoadRedisConfig`

Model konfigurasi: `internal/model/config.go` (struct ConfigModel, RDBMS, dsb).

Catatan: Periksa konsistensi key di file config dan loader. Ada potensi mismatch antara key `database.pgsql.db_name.*` vs `database.pgsql.db_init.*`.

Struktur proyek (ringkas)
- main.go — entry point aplikasi
- cmd/
  - root.go — root Cobra command
  - rest.go — command untuk serve
- boot/ — inisialisasi konfigurasi, logger, DB, Redis
- internal/
  - router/ — pembuatan router Gin dan pendaftaran route
  - handler/ — handler HTTP (mis. health)
  - middleware/ — middleware (mis. logging)
  - repository/ — util DB dan repositori
  - utility/ — helper response, dsb
  - model/ — model konfigurasi & response

Endpoint contoh
- Health check
  - Method: GET
  - URL: /api/v1/health/
  - Contoh:
  ```sh
  curl -v http://localhost:8080/api/v1/health/
  ```

Logging & observabilitas
- Logger: inisialisasi di `boot/logger.go` (zap).
- Middleware logging: `internal/middleware/logger.go`.
- Periksa output logger saat gagal load config atau gagal koneksi DB/Redis.

Database & Redis
- Inisialisasi PostgreSQL: `boot/PgSqlInit` (mengembalikan *sqlx.DB).
- Inisialisasi Redis: `boot/RedisInit` (mengembalikan *redis.Client).
- DB helper: `internal/repository/util/db_exec.go` (Exec/Get/Select wrappers).

Testing
- Tambahkan unit test di package terkait (`*_test.go`).  
- Jalankan:
```sh
go test ./...
```

Debugging cepat
- Periksa log (stdout) untuk pesan fatal dari boot atau koneksi.
- Pastikan `config.yaml` sesuai environment (path key, credential).
- Jika menggunakan Docker / container, pastikan network DB/Redis reachable.

File penting
- main.go
- cmd/root.go, cmd/rest.go
- boot/*.go
- internal/router/*, internal/handler/*, internal/repository/*, internal/utility/*, internal/model/*
- config.yaml
- go.mod

Kontribusi
- Ikuti gaya kode Go (gofmt/golint).
- Buat branch per fitur/bugfix dan ajukan PR.