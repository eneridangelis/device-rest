# Makefile

# Variáveis para conexão com o banco de dados
# [ene] colocar no arquivo .env depois
DB_USER=user
DB_PASSWORD=1234
DB_NAME=db_test
DB_HOST=localhost
DB_PORT=5432
MIGRATIONS_DIR=./migrations

# Instalação do golang-migrate
install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Comando para rodar as migrações
migrate-up:
	migrate -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path $(MIGRATIONS_DIR) up

# Comando para reverter a última migração
migrate-down:
	migrate -database postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path $(MIGRATIONS_DIR) down 1

# Comando para criar uma nova migração
create-migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name