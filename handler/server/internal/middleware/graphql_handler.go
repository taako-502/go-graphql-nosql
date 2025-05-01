package middleware

import (
	"context"
	"net/http"

	"github.com/taako-502/go-graphql-nosql/handler/graph"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/guregu/dynamo/v2"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GraphqlHandler(db *dynamo.DB, region string) http.HandlerFunc {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		gqlerr := graphql.DefaultErrorPresenter(ctx, e)
		return gqlerr
	})

	return func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	}
}
