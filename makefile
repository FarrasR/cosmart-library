.PHONY: test build dep build run all

dep:
	go mod download
	go mod verify

test: 
	go test ./...

build: 
	go build -o cosmart-library main.go

run: build
	./cosmart-library
	 
migrate:
	go build -o cosmart-library-migrate ./database/migration/main.go
	./cosmart-library-migrate

seed:
	go build -o cosmart-library-seed ./database/seed/main.go
	./cosmart-library-seed

all:
	dep
	migrate
	seed
	run

