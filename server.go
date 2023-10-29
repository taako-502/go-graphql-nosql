package main

import (
	"fmt"
	"go-graphql-nosql/example"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 環境変数読み込み
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// 環境変数から設定値を取得
	awsRegion := os.Getenv("AWS_REGION")
	dynamoEndpoint := os.Getenv("DYNAMO_ENDPOINT")

	// クライアントの設定
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Endpoint:    aws.String(dynamoEndpoint),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}

	// DynamoDB
	db := dynamo.New(sess)

	// Echo API（https://echo.labstack.com/）
	e := echo.New()
	// CORSの設定（https://echo.labstack.com/docs/middleware/cors）
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3333"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	e.GET("/example", func(c echo.Context) error {
		// 動作確認用のサンプルプログラム
		if err := example.Example(db); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "okey")
	})
	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
