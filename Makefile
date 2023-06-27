-include .env_example
export $(shell sed 's/=.*//' .env_example)

.PHONY: run-core
run-core: go-tidy
	go run services/core/cmd/core/main.go -l

.PHONY: run
run: go-tidy
	docker compose -f docker-compose.yml up -d --build

COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 DOCKER_DEFAULT_PLATFORM=linux/amd64 docker compose -f docker-compose.test.yml up -d --build

.PHONY: run-local
run-local:
	COMPOSE_DOCKER_CLI_BUILD=1 \
	DOCKER_BUILDKIT=1 \
	DOCKER_DEFAULT_PLATFORM=linux/amd64 \
	docker compose -f docker-compose.yml up -d --build price_generator auth_generator interview_db && \
	go run services/core/main.go --local

.PHONY: go-tidy
go-tidy:
	cd ./pkg && go mod tidy -e && cd .. && \
	cd ./proto && go mod tidy -e && cd .. && \
	cd ./services/core && go mod tidy -e && cd ../.. && \
	cd ./services/price_generator && go mod tidy -e && cd ../..

.PHONY: go-load
go-load:
	go mod download

.PHONY: protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./proto/*/*.proto

.PHONY: swag
swag:
	swag init -d ./services/core -o ./services/core/api/swagger  -g internal/handler/handler.go

## goose-install: install goose project
.PHONY: goose-install
goose-install:
	go install github.com/pressly/goose/v3/cmd/goose@v3.5.0

## goose-create: create a new migration file
.PHONY: goose-create
goose-create:
	@goose -dir ${MIGRATION_DIR} create init sql

## goose-status: check status of the migration
.PHONY: goose-status
goose-status:
	@goose -dir ${MIGRATION_DIR} ${DB_TYPE} "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" status

## goose-up: up a new migration
.PHONY: goose-up
goose-up:
	@goose -dir ${MIGRATION_DIR} ${DB_TYPE} "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" up

## goose-down: down a last migration
.PHONY: goose-down
goose-down:
	@goose -dir ${MIGRATION_DIR} ${DB_TYPE} "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" down

## goose-redo: reup a last migration
.PHONY: goose-redo
goose-redo:
	@goose -dir ${MIGRATION_DIR} ${DB_TYPE} "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" redo

## goose-reset: reup all migration
.PHONY: goose-reset
goose-reset:
	@goose -dir ${MIGRATION_DIR} ${DB_TYPE} "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" reset


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo