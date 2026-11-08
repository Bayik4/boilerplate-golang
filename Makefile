MIGRATE=migrate
MIGRATIONS_PATH=./migrations
DATABASE_URL?=postgres://user_initial:123@localhost:5432/db_initial?sslmode=disable

.PHONY: migrate-create migrate-up migrate-down migrate-version migrate-force

migrate-create:
	@$(MIGRATE) create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

migrate-up:
	@$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

migrate-down:
	@$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

migrate-version:
	@$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version

migrate-force:
	@$(MIGRATE) -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force $(version)