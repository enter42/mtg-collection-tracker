.PHONY: build run clean test db-create

build:
	go build -o bin/server cmd/server/main.go

run:
	go run cmd/server/main.go

clean:
	rm -rf bin/

test:
	go test -v ./...

db-create:
	mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS mtg_collection;"

tidy:
	go mod tidy

install:
	go mod download
