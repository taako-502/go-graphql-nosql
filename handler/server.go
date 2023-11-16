package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	ddbmanager "go-graphql-nosql/handler/dynamodb"
	"go-graphql-nosql/handler/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "8080"

func main() {
	// 環境変数読み込み
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// コマンドライン引数をパース
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	// go run server.go -migrate
	if *migrate {
		// DynamoDBの初期化
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

	db := ddbmanager.New(os.Getenv(""))
	// DynamoDB
	// サンプルプログラム（一時コメントアウト）
	// if err := example.Example(db); err != nil {
	// 	panic(err)
	// }

	// 環境変数から設定値を取得
	graphqlServerPort := os.Getenv("GRAPHQL_SERVER_PORT")
	if graphqlServerPort == "" {
		graphqlServerPort = defaultPort
	}

	// GraphQLサーバーの設定
	router := chi.NewRouter()
	// frontendHost := os.Getenv("FRONTEND_HOST")
	// graphqlServerHost := os.Getenv("GRAPHQL_SERVER_HOST")
	router.Use(cors.New(cors.Options{
		// AllowedOrigins: []string{
		// 	frontendHost,
		// 	// graphqlServerHost,
		// },
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodPut,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	graphqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))
	domain := os.Getenv("DOMAIN")
	graphqlServer.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return r.Host == domain
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	graphqlServer.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		var extension map[string]interface{}
		if errors.Is(err, graph.ErrUserNotFound) {
			extension = map[string]interface{}{
				"code":   "USER_NOT_FOUND",
				"status": http.StatusBadRequest,
			}
		} else if errors.Is(err, graph.ErrCodeLoginFailed) {
			extension = map[string]interface{}{
				"code":   "LOGIN_FAILED",
				"status": http.StatusBadRequest,
			}
		}
		return &gqlerror.Error{
			Message:    err.Error(),
			Path:       graphql.GetPath(ctx),
			Locations:  nil,
			Extensions: extension,
			Rule:       "",
		}
	})
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", graphqlServer)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", graphqlServerPort)
	if err := http.ListenAndServe(":"+graphqlServerPort, router); err != nil {
		panic(err)
	}
}
