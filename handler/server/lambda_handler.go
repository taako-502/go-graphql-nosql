package server

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	ddbmanager "github.com/taako-502/go-graphql-nosql/handler/dynamodb"
	"github.com/taako-502/go-graphql-nosql/handler/server/internal/middleware"
)

// Handler is the main function called by AWS Lambda.
func (s *server) LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	mux := http.NewServeMux()

	DB, err := ddbmanager.New(ctx, s.awsConfig.region)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	mux.Handle("POST /query", middleware.GraphqlHandler(DB, s.awsConfig.region))

	handler := middleware.CORS(mux, s.corsAllowedOrigins)
	adapter := httpadapter.New(handler)

	return adapter.ProxyWithContext(ctx, req)
}
