package ddbmanager

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// 引数はオプショナルにすること
func New(endpoint string) *dynamo.DB {
	// セッションを初期化
	awsConfig := getAwsConfig(endpoint)
	sess := session.Must(session.NewSession(awsConfig))

	// DynamoDBサービスクライアントを作成
	return dynamo.New(sess)
}

func getAwsConfig(endpoint string) *aws.Config {
	awsRegion := os.Getenv("AWS_REGION")
	dynamoEndpoint := endpoint
	if dynamoEndpoint == "" {
		dynamoEndpoint = os.Getenv("DYNAMO_ENDPOINT")
	}
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	token := os.Getenv("AWS_SESSION_TOKEN")
	return &aws.Config{
		Region:   aws.String(awsRegion),
		Endpoint: aws.String(dynamoEndpoint),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			token,
		),
	}
}
