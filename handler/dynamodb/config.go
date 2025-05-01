package ddbmanager

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/guregu/dynamo/v2"
)

var TEST_REGION = "ap-northeast-1"

// 引数はオプショナルにすること
func New(ctx context.Context, region string) (*dynamo.DB, error) {
	// NOTE: 一時的にコメントアウト
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	// DynamoDBサービスクライアントを作成
	return dynamo.New(cfg), nil
}

// amazon/dynamodb-local に接続する
func NewByLocalOrCI(ctx context.Context, endpoint string) (*dynamo.DB, error) {
	conf, err := getAwsConfigForDummy(ctx, endpoint)
	if err != nil {
		return nil, fmt.Errorf("GetAwsConfigForDummy(%s), %v", endpoint, err)
	}

	return dynamo.New(*conf), nil
}

// getAwsConfigForDummy は、DynamoDB Localに接続するためのAWS Configを取得します。
func getAwsConfigForDummy(ctx context.Context, endpoint string) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(TEST_REGION),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: endpoint}, nil
			}),
		),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			"dummy",
			"dummy",
			"",
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration, %v", err)
	}

	return &cfg, nil
}
