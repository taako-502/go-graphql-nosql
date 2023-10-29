package main

import (
	"fmt"
	"go-graphql-nosql/example"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
)

func main() {
	// 環境変数読み込み
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// 環境変数から設定値を取得
	awsRegion := os.Getenv("AWS_REGION")
	dynamoEndpoint := os.Getenv("DYNAMO_ENDPOINT")

	// クライアントの設定
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Endpoint:    aws.String(dynamoEndpoint),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}

	db := dynamo.New(sess)
	example.Example(db)
}
