package ddbmanager

import (
	"go-graphql-nosql/handler/graph/model"
	"log"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

type DDBMnager struct {
	DB *dynamo.DB
}

func (d *DDBMnager) Migration() error {
	tables := []string{"Todo", "User"}
	for _, table := range tables {
		exist, err := d.TableExists(table)
		if err != nil {
			return err
		}
		if exist {
			log.Printf("Table %s already exists", table)
			continue
		}

		if err := d.TableCreate(table); err != nil {
			return err
		}
	}
	return nil
}

func (d *DDBMnager) TableExists(tableName string) (bool, error) {
	if _, err := d.DB.Table(tableName).Describe().Run(); err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (d *DDBMnager) TableCreate(tableName string) error {
	if err := d.DB.CreateTable(tableName, model.Todo{}).Run(); err != nil {
		log.Fatalf("Unable to create table: %s", err)
		return err
	}
	return nil
}
