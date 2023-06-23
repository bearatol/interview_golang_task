include .env_example
export $(shell sed 's/=.*//' .env_example)

.PHONY: run
run:
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


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo