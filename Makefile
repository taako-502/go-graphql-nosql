build:
	go mod tidy
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler handler/server/server.go

gqlgen:
	go run github.com/99designs/gqlgen generate
