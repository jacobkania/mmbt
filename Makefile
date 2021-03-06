include .env
include .env.local

#####################
#### Primary Commands
#####################

init: fetch-dependencies-go fetch-dependencies-js migrate-up

dev:
	make run-db-local && make run-go & make run-js

build: build-go build-js

test: run-db-local migrate-test test-go

start:
	(cd bin/ && ./mmbt)

bs:  run-db-local build start

kill:
	kill $(lsof -t -i :8080); kill $(lsof -t -i :3000)

#####################
#### Single Application Commands
#####################

# Go

fetch-dependencies-go:
	go get

run-go:
	go run main.go

build-go:
	go build -o bin/mmbt main.go

test-go:
	@TEST_DB_URL=$(TEST_DB_URL) go test ./...

# JS

fetch-dependencies-js:
	yarn --cwd js

run-js:
	yarn --cwd js run start

build-js:
	yarn --cwd js run build

run-js-reload:
	yarn --cwd js run start-reload-theme

# Database

new-migration:
	migrate create -ext sql -dir db/migration -format "20060102030405" $(NAME)

migrate-up:
	migrate -path db/migration -database $(DB_URL) -verbose up

migrate-last-down:
	migrate -path db/migration -database $(DB_URL) -verbose down 1

migrate-drop:
	migrate -path db/migration -database $(DB_URL) -verbose drop

migrate-test:
	migrate -path db/migration -database $(TEST_DB_URL) -verbose drop -f
	migrate -path db/migration -database $(TEST_DB_URL) -verbose up

run-db-local:
	brew services start postgres

stop-db-local:
	brew services stop postgres

#####################
#### Utils
#####################
