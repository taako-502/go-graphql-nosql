package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	ddbmanager "github.com/taako-502/go-graphql-nosql/handler/dynamodb"
)

func main() {
	// 環境変数読み込み
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// ローカル環境で打鍵するときに使う
	ctx := context.Background()
	db, err := ddbmanager.NewForLocal(ctx, os.Getenv("MIGRATION_ENDPOINT"))
	if err != nil {
		log.Fatalf("DynamoDBの接続に失敗しました: %v", err)
	}
	manager := ddbmanager.DDBMnager{DB: db}

	// マイグレーション実行
	fmt.Println("Running migrations...")
	if err := manager.Migration(ctx); err != nil {
		log.Fatalf("マイグレーションに失敗しました: %v", err)
		os.Exit(1)
	}
	fmt.Println("マイグレーションが完了しました。")
}
