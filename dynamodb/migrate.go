package ddbmanager

import (
	"go-graphql-nosql/graph/model"
	"log"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

type DDBMnager struct {
	DB *dynamo.DB
}

func (d *DDBMnager) Migration() error {
	// exist, err := d.TableExists("Todo")
	// if err != nil {
	// 	return err
	// }
	// if !exist {
	// 	d.TableCreate("Todo")
	// }
	if err := d.TableCreate("Todo"); err != nil {
		return err
	}
	return nil
}

func (d *DDBMnager) TableExists(tableName string) (bool, error) {
	if _, err := d.DB.Table("Todos").Describe().Run(); err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (d *DDBMnager) TableCreate(tableName string) error {
	if err := d.DB.CreateTable("Todo", model.Todo{}).Run(); err != nil {
		log.Fatalf("Unable to create table: %s", err)
		return err
	}
	return nil
}
