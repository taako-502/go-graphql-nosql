package ddbmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/guregu/dynamo/v2"
)

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
