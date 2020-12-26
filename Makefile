#####################
#### Utils
#####################

create-dev-db-file-if-not-exist:
	touch -a ./mmbt.db

create-bin-db-file-if-not-exist:
	touch -a ./bin/mmbt.db

#####################
#### Single Application Commands
#####################

# Go

run-go:
	go run main.go

build-go:
	go build -o bin/mmbt main.go

test-go:
	go test

# JS

run-js:
	yarn --cwd js run start

build-js:
	yarn --cwd js run build

#####################
#### Primary Commands
#####################

dev:
	make create-db-file-if-not-exist && make run-go & make run-js

build: build-go build-js

test: test-go

start:
	(cd bin/ && ./mmbt)

bs: create-bin-db-if-not-exist build start
