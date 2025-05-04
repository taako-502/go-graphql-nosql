package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/taako-502/go-graphql-nosql/pkg/dynamodb_manager"
	"github.com/taako-502/go-graphql-nosql/pkg/handler/server/internal/middleware"
)

func (s *server) LocalServer() error {
	ctx := context.Background()
	DB, err := dynamodb_manager.NewForLocal(ctx, s.dynamoEndpoint)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("GET /", playground.Handler("GraphQL playground", "/graphql"))
	mux.Handle("POST /graphql", middleware.GraphqlHandler(DB, s.awsConfig.region))
	handlerWithCORS := middleware.CORS(mux, s.corsAllowedOrigins)

	graphqlServerPort := s.graphqlServerPort
	if graphqlServerPort == "" {
		graphqlServerPort = "8080"
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", graphqlServerPort)
	if err := http.ListenAndServe(":"+graphqlServerPort, handlerWithCORS); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
