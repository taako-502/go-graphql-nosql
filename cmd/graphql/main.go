package main

import (
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/taako-502/go-graphql-nosql/handler/server"
)

func main() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" || env == "local" {
		// 環境変数ENVIRONMENTが空白の場合はローカル環境とみなす
		if err := godotenv.Load(".env"); err != nil {
			log.Printf("環境変数の読込に失敗しました: %v\r\n", err)
		}
	}

	server := server.NewServer(
		os.Getenv("DYNAMO_REGION"),
		os.Getenv("AWS_S3_BUCKET"),
		strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		os.Getenv("DYNAMO_ENDPOINT"),
		os.Getenv("GRAPHQL_SERVER_PORT"),
		os.Getenv("AUTH_SERVER_URL"),
	)

	// Lambda環境
	if env == "dev" || env == "prod" {
		lambda.Start(server.LambdaHandler)
	}

	// ローカル環境
	panic(server.LocalServer())
}
