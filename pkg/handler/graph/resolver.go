package graph

import (
	"github.com/guregu/dynamo/v2"
)

type Resolver struct {
	DB      *dynamo.DB
	dbNames DBNames
}

type DBNames struct {
	User string
	Todo string
}

func NewResolver(db *dynamo.DB, dbNames DBNames) *Resolver {
	if dbNames.User == "" {
		dbNames.User = "User"
	}
	if dbNames.Todo == "" {
		dbNames.Todo = "Todo"
	}
	return &Resolver{
		DB:      db,
		dbNames: dbNames,
	}
}

func (r *Resolver) GetUserTableName() string {
	return r.dbNames.User
}

func (r *Resolver) GetTodoTableName() string {
	return r.dbNames.Todo
}
