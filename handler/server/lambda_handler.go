package server

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/taako-502/go-graphql-nosql/handler/dynamodb_manager"
	"github.com/taako-502/go-graphql-nosql/handler/server/internal/middleware"
)

// Handler is the main function called by AWS Lambda.
func (s *server) LambdaHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received request: %s %s", event.HTTPMethod, event.Path)

	DB, err := dynamodb_manager.New(ctx, s.awsConfig.region)
	if err != nil {
		log.Printf("Failed to create DynamoDB client: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to create DynamoDB client",
		}, err
	}

	log.Println("Successfully created DynamoDB client")

	mux := http.NewServeMux()
	// AWS LambdaとAPI Gateway側で/graphqlに接続している
	mux.Handle("POST /", middleware.GraphqlHandler(DB, s.awsConfig.region))

	// NOTE: CORSはAPI Gatewayで設定
	adapter := httpadapter.New(mux)

	log.Println("Invoking HTTP adapter")
	return adapter.ProxyWithContext(ctx, event)
}
