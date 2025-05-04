package middleware

import (
	"net/http"
	"os"

	"github.com/taako-502/go-graphql-nosql/pkg/handler/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/guregu/dynamo/v2"
)

func GraphqlHandler(DB *dynamo.DB, region string) http.HandlerFunc {
	c := graph.Config{Resolvers: graph.NewResolver(DB,
		graph.DBNames{
			User: os.Getenv("USER_TABLE_NAME"),
			Todo: os.Getenv("TODO_TABLE_NAME"),
		},
	)}
	srv := handler.New(graph.NewExecutableSchema(c))
	srv.AddTransport(transport.POST{})
	return func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	}
}
