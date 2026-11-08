.PHONY: migrate-up migrate-down migrate-version migrate-force migrate-create

migrate-up:
	go run . migrate up

migrate-down:
	go run . migrate down

migrate-version:
	go run . migrate version

migrate-force:
ifndef version
	$(error version is required. Example: make migrate-force version=1)
endif
	go run . migrate force $(version)

migrate-create:
ifndef name
	$(error name is required. Example: make migrate-create name=create_users_table)
endif
	go run . migrate create $(name)