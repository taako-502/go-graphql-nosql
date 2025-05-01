package middleware

import (
	"net/http"

	"github.com/taako-502/go-graphql-nosql/handler/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/guregu/dynamo/v2"
)

func GraphqlHandler(DB *dynamo.DB, region string) http.HandlerFunc {
	// FIXME: これは古い書き方なので、新しい書き方Newに変更する
	c := graph.Config{Resolvers: &graph.Resolver{
		DB: DB,
	}}
	srv := handler.New(graph.NewExecutableSchema(c))
	srv.AddTransport(transport.POST{})
	return func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	}
}
