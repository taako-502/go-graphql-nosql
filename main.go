package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
)

type User struct {
	UserID string `dynamo:"UserID,hash"`
	Name   string `dynamo:"Name,range"`
	Age    int    `dynamo:"Age"`
	Text   string `dynamo:"Text"`
}

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

	db := dynamo.New(sess)

	// テーブル作成をする為に、一度テーブルを削除します
	db.Table("UserTable").DeleteTable().Run()

	// テーブル作成
	err = db.CreateTable("UserTable", User{}).Run()
	if err != nil {
		panic(err)
	}
	// テーブルの指定
	table := db.Table("UserTable")

	// User構造体をuser変数に定義
	var user User

	// DBにPutします
	err = table.Put(&User{UserID: "1234", Name: "太郎", Age: 20}).Run()
	if err != nil {
		panic(err)
	}

	// DBからGetします
	err = table.Get("UserID", "1234").Range("Name", dynamo.Equal, "太郎").One(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetDB%+v\n", user)

	// DBのデータをUpdateします
	text := "新しいtextです"
	err = table.Update("UserID", "1234").Range("Name", "太郎").Set("Text", text).Value(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("UpdateDB%+v\n", user)

	// DBのデータをDeleteします
	err = table.Delete("UserID", "1234").Range("Name", "太郎").Run()
	if err != nil {
		panic(err)
	}

	// Delete出来ているか確認
	err = table.Get("UserID", "1234").Range("Name", dynamo.Equal, "太郎").One(&user)
	if err != nil {
		// Delete出来ていれば、dynamo: no item found のエラーとなる
		fmt.Println("getError:", err)
	}
}
