package ddbmanager

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/guregu/dynamo/v2"
)

// 引数はオプショナルにすること
func New(endpoint string) *dynamo.DB {
	// NOTE: 一時的にコメントアウト
	return nil

	// var sess *session.Session
	// if os.Getenv("ENVIRONMENT") == "local" {
	// 	sess = session.Must(session.NewSession(getAwsConfig(endpoint)))
	// } else {
	// 	sess = session.Must(session.NewSession())
	// }

	// // DynamoDBサービスクライアントを作成
	// return dynamo.New(sess)
}

func getAwsConfig(endpoint string) *aws.Config {
	// NOTE: 一時的にコメントアウト
	return nil

	// return &aws.Config{
	// 	Region:   aws.String("ap-northeast-1"),
	// 	Endpoint: aws.String(endpoint),
	// 	Credentials: credentials.NewStaticCredentials(
	// 		"dammy",
	// 		"dammy",
	// 		"dammy",
	// 	),
	// }
}
