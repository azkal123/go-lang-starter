.PHONY: run build test migrate-up migrate-down clean

run:
	go run cmd/main.go

build:
	go build -o bin/server cmd/main.go

test:
	go test -v ./...

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

clean:
	rm -rf bin/
	go clean
