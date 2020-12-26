#####################
#### Primary Commands
#####################

init: fetch-dependencies-go fetch-dependencies-js init-dev-db init-bin-db

dev:
	make run-go & make run-js

build: build-go build-js

test: test-go

start:
	(cd bin/ && ./mmbt)

bs:  build start

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
	go test

# JS

fetch-dependencies-js:
	yarn --cwd js

run-js:
	yarn --cwd js run start

build-js:
	yarn --cwd js run build

#####################
#### Utils
#####################

init-dev-db:
	touch -a ./mmbt.db

init-bin-db:
	touch -a ./bin/mmbt.db
