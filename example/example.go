package example

import (
	"fmt"

	"github.com/guregu/dynamo"
)

type User struct {
	UserID string `dynamo:"UserID,hash"`
	Name   string `dynamo:"Name,range"`
	Age    int    `dynamo:"Age"`
	Text   string `dynamo:"Text"`
}

func Example(db *dynamo.DB) {
	// テーブル作成をする為に、一度テーブルを削除します
	db.Table("UserTable").DeleteTable().Run()

	// テーブル作成
	if err := db.CreateTable("UserTable", User{}).Run(); err != nil {
		panic(err)
	}
	// テーブルの指定
	table := db.Table("UserTable")

	// User構造体をuser変数に定義
	var user User

	// DBにPutします
	if err := table.Put(&User{UserID: "1234", Name: "太郎", Age: 20}).Run(); err != nil {
		panic(err)
	}

	// DBからGetします
	if err := table.Get("UserID", "1234").Range("Name", dynamo.Equal, "太郎").One(&user); err != nil {
		panic(err)
	}
	fmt.Printf("GetDB%+v\n", user)

	// DBのデータをUpdateします
	text := "新しいtextです"
	if err := table.Update("UserID", "1234").Range("Name", "太郎").Set("Text", text).Value(&user); err != nil {
		panic(err)
	}
	fmt.Printf("UpdateDB%+v\n", user)

	// DBのデータをDeleteします
	if err := table.Delete("UserID", "1234").Range("Name", "太郎").Run(); err != nil {
		panic(err)
	}

	// Delete出来ているか確認
	if err := table.Get("UserID", "1234").Range("Name", dynamo.Equal, "太郎").One(&user); err != nil {
		// Delete出来ていれば、dynamo: no item found のエラーとなる
		fmt.Println("getError:", err)
	}
}
