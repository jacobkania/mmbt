run-go:
	go run src/main.go

build-go:
	go build -o bin/main src/main.go

test-go:
	go test ./src

run-js:
	yarn --cwd js run start

build-js:
	yarn --cwd js run build

dev: run-go run-js

build: build-go build-js

test: test-go
