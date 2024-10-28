# Makefile

DB_USER=user
DB_PASSWORD=1234
DB_NAME=db_test
DB_HOST=localhost
DB_PORT=5432
MIGRATIONS_DIR=./migrations

infra-up:
	docker-compose up -d

run:
	go run cmd/main.go

test:
	go test ./...

install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-up:
	migrate -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path $(MIGRATIONS_DIR) up

migrate-down:
	migrate -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path $(MIGRATIONS_DIR) down 1

create-migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name