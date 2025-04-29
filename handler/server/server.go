package main

import (
	"context"
	"flag"
	"fmt"
	ddbmanager "go-graphql-nosql/handler/dynamodb"
	"go-graphql-nosql/handler/graph"
	"go-graphql-nosql/handler/server/internal/config"
	"go-graphql-nosql/handler/server/internal/localserver"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
)

var ginLambda *ginadapter.GinLambda

// Defining the Graphql handler
func graphqlHandler(db *dynamo.DB) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Handler is the main function called by AWS Lambda.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO: EchoAPIに書き換える
	// If no name is provided in the HTTP request body, throw an error
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		r := gin.Default()
		hosts := []string{
			os.Getenv("FRONTEND_HOST_1"),
			os.Getenv("FRONTEND_HOST_2"),
			os.Getenv("FRONTEND_HOST_3"),
		}
		r.Use(config.SettingCors(hosts))

		// Setting up Gin
		db := ddbmanager.New("")
		r.POST("/query", graphqlHandler(db))

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// 環境変数読み込み
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// コマンドライン引数をパース
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// ローカル環境で打鍵するときに使う
	// go run handler/server.go -migrate
	if *migrate {
		endpoint := os.Getenv("MIGRATION_ENDPOINT")
		db := ddbmanager.New(endpoint)
		manager := ddbmanager.DDBMnager{DB: db}

		// マイグレーション実行
		fmt.Println("Running migrations...")
		if err := manager.Migration(); err != nil {
			log.Fatalf("マイグレーションに失敗しました: %v", err)
			os.Exit(1)
		}
		fmt.Println("マイグレーションが完了しました。")
		return
	}

	if os.Getenv("ENVIRONMENT") == "local" {
		// ローカル環境
		endpoint := os.Getenv("DYNAMO_ENDPOINT")
		db := ddbmanager.New(endpoint)
		localserver.StartLocalServer(db)
	} else {
		// AWS Lambda
		lambda.Start(Handler)
	}
}
