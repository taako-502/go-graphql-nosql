package ddbmanager

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// 引数はオプショナルにすること
func New(endpoint string) *dynamo.DB {
	// セッションを初期化
	dynamoEndpoint := endpoint
	if dynamoEndpoint == "" {
		dynamoEndpoint = os.Getenv("DYNAMO_ENDPOINT")
	}
	awsRegion := os.Getenv("AWS_SS_REGION")
	accessKeyID := os.Getenv("AWS_SS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SS_SECRET_ACCESS_KEY")
	token := os.Getenv("AWS_SS_SESSION_TOKEN")
	awsConfig := getAwsConfig(dynamoEndpoint, awsRegion, accessKeyID, secretAccessKey, token)
	sess := session.Must(session.NewSession(awsConfig))

	// DynamoDBサービスクライアントを作成
	return dynamo.New(sess)
}

func getAwsConfig(endpoint string, awsRegion string, accessKeyID string, secretAccessKey string, token string) *aws.Config {
	return &aws.Config{
		Region:   aws.String(awsRegion),
		Endpoint: aws.String(strings.TrimSpace(endpoint)), // 前後のタブ文字や空白を削除
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			token,
		),
	}
}
