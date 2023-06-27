.PHONY: goose-up-docker
goose-up-docker:
	/app/goose/bin/goose -dir ${MIGRATION_DIR} ${DB_TYPE} "host=${POSTGRES_HOST} port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" up