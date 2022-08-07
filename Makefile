BIN_NAME=advisree-be
#!make
include .env
export $(shell sed 's/=.*//' .env)

run: generate-docs
	@go run main.go

generate-docs:
	@swag init

migrate-up:
	@migrate -path db/migrations -database "mysql://${DB_USER}:${DB_PASS}@(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose up

migrate-create:
	@migrate create -ext sql -dir db/migrations $(migration)