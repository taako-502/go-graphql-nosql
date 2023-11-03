package model

type Todo struct {
	ID     string `dynamo:"ID,hash"`
	Text   string `dynamo:"Text"`
	Done   bool   `dynamo:"Done"`
	UserID string `dynamo:"UserID"`
}
