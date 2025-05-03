build:
	go mod tidy
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler cmd/graphql/main.go

gqlgen:
	go run github.com/99designs/gqlgen generate

migrate:
	go run cmd/migrate/main.go
