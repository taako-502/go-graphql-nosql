package localserver

import (
	ddbmanager "go-graphql-nosql/handler/dynamodb"
	"go-graphql-nosql/handler/graph"
	"go-graphql-nosql/handler/internal/config"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

func StartLocalServer() {
	db := ddbmanager.New()

	// Routerの設定
	router := chi.NewRouter()
	router.Use(config.SettingCrosForLocalServer().Handler)
	graphqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))
	graphqlServer.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return r.Host == os.Getenv("DOMAIN")
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", graphqlServer)

	graphqlServerPort := os.Getenv("GRAPHQL_SERVER_PORT")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", graphqlServerPort)
	if err := http.ListenAndServe(":"+graphqlServerPort, router); err != nil {
		panic(err)
	}
}
