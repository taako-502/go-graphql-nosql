package server

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/taako-502/go-graphql-nosql/handler/dynamodb_manager"
	"github.com/taako-502/go-graphql-nosql/handler/server/internal/middleware"
)

// Handler is the main function called by AWS Lambda.
func (s *server) LambdaHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	mux := http.NewServeMux()

	DB, err := dynamodb_manager.New(ctx, s.awsConfig.region)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to create DynamoDB client",
		}, err
	}
	mux.Handle("POST /graphql", middleware.GraphqlHandler(DB, s.awsConfig.region))

	handler := middleware.CORS(mux, s.corsAllowedOrigins)
	adapter := httpadapter.New(handler)

	return adapter.ProxyWithContext(ctx, event)
}
