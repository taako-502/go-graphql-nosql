package main

import (
	"fmt"
	"go-graphql-nosql/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	// 環境変数読み込み
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// 環境変数から設定値を取得
	awsRegion := os.Getenv("AWS_REGION")
	dynamoEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	graphqlServerPort := os.Getenv("GRAPHQL_SERVER_PORT")
	if graphqlServerPort == "" {
		graphqlServerPort = defaultPort
	}

	// クライアントの設定
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Endpoint:    aws.String(dynamoEndpoint),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}

	// DynamoDB
	dynamo.New(sess)
	// db := dynamo.New(sess)
	// サンプルプログラム（一時コメントアウト）
	// if err := example.Example(db); err != nil {
	// 	panic(err)
	// }

	// GraphQLサーバーの設定
	graphqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphqlServer)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", graphqlServerPort)
	log.Fatal(http.ListenAndServe(":"+graphqlServerPort, nil))
}
