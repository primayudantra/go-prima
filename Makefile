include .env

DATABASE_URL=postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

pre-commit:
	go mod tidy
	go mod vendor

# Run Up Migration
migration-up:
	migrate -path db/migration -database "${DATABASE_URL}" -verbose up

# Run Down Migration
migration-down:
	migrate -path db/migration -database "${DATABASE_URL}" -verbose down

# Create Migration
migration-create:
	migrate create -ext sql -dir db/migration -seq $(filename)

# Create Database
create-db:
	docker exec -it simple-projects-postgres createdb --username=${DB_USERNAME} --owner=${DB_USERNAME} ${DB_NAME}

# Drop Database
drop-db:
	docker exec -it simple-projects-postgres dropdb --username=${DB_USERNAME} --owner=${DB_USERNAME} ${DB_NAME}

run-docker-compose:
	docker-compose up -d

sqlc:
	sqlc generate

.PHONY: migration-up migration-down migration-create create-db drop-db