package ddbmanager

import (
	"context"
	"fmt"
	"log"

	"github.com/taako-502/go-graphql-nosql/handler/graph/model"

	"github.com/guregu/dynamo/v2"
)

type DDBMnager struct {
	DB *dynamo.DB
}

func (d *DDBMnager) Migration(ctx context.Context) error {
	tables := []string{"Todo", "User"}
	for _, table := range tables {
		exist, err := d.TableExists(ctx, table)
		if err != nil {
			return fmt.Errorf("DDBMnager.TableExists: %w", err)
		}
		if exist {
			log.Printf("Table %s already exists", table)
			continue
		}

		if err := d.TableCreate(ctx, table); err != nil {
			return fmt.Errorf("DDBMnager.TableCreate: %w", err)
		}
	}
	return nil
}

func (d *DDBMnager) TableExists(ctx context.Context, tableName string) (bool, error) {
	if _, err := d.DB.Table(tableName).Describe().Run(ctx); err != nil {
		// FIXME: テーブルが存在しない場合のエラーを確認する
		// var codeResourceNotFoundException = dynamodb.ErrCodeResourceNotFoundException
		// if errors.As(err, &codeResourceNotFoundException) {
		// 	return false, nil
		// }
		return false, fmt.Errorf("DDBMnager.TableExists: %w", err)
	}
	return true, nil
}

func (d *DDBMnager) TableCreate(ctx context.Context, tableName string) error {
	if err := d.DB.CreateTable(tableName, model.Todo{}).Run(ctx); err != nil {
		log.Fatalf("Unable to create table: %s", err)
		return fmt.Errorf("DDBMnager.TableCreate: %w", err)
	}
	return nil
}
