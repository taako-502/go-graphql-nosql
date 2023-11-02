package main

import (
	"flag"
	"fmt"
	ddbmanager "go-graphql-nosql/dynamodb"
	"go-graphql-nosql/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	// 環境変数読み込み
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// コマンドライン引数をパース
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// go run main.go -migrate
	if *migrate {
		// DynamoDBの初期化
		endpoint := os.Getenv("MIGRATION_ENDPOINT")
		db := ddbmanager.New(endpoint)
		manager := ddbmanager.DDBMnager{DB: db}

		// マイグレーション実行
		fmt.Println("Running migrations...")
		err := manager.Migration()
		if err != nil {
			log.Fatalf("マイグレーションに失敗しました: %v", err)
			os.Exit(1)
		}
		fmt.Println("マイグレーションが完了しました。")
		return
	}

	db := ddbmanager.New(os.Getenv(""))
	// DynamoDB
	// サンプルプログラム（一時コメントアウト）
	// if err := example.Example(db); err != nil {
	// 	panic(err)
	// }

	// 環境変数から設定値を取得
	graphqlServerPort := os.Getenv("GRAPHQL_SERVER_PORT")
	if graphqlServerPort == "" {
		graphqlServerPort = defaultPort
	}

	// GraphQLサーバーの設定
	graphqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphqlServer)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", graphqlServerPort)
	log.Fatal(http.ListenAndServe(":"+graphqlServerPort, nil))
}
