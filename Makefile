run-go:
	go run main.go

build-go:
	go build -o bin/main main.go

test-go:
	go test

run-js:
	yarn --cwd js run start

build-js:
	yarn --cwd js run build





dev:
	make run-go & make run-js

build: build-go build-js

test: test-go
