package middleware

import (
	"net/http"

	"github.com/taako-502/go-graphql-nosql/handler/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/guregu/dynamo/v2"
)

func GraphqlHandler(DB *dynamo.DB, region string) http.HandlerFunc {
	// FIXME: これは古い書き方なので、新しい書き方Newに変更する
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			DB: DB,
		},
	}))
	return func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	}
}
