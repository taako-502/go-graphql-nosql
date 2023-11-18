package ddbmanager

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// 引数はオプショナルにすること
func New(endpoint string) *dynamo.DB {
	sess := session.Must(session.NewSession())

	// DynamoDBサービスクライアントを作成
	return dynamo.New(sess)
}
