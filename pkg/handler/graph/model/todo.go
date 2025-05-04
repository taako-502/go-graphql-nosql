package model

type Todo struct {
	ID          string `dynamo:"ID,hash"`
	Title       string `dynamo:"Title"`
	Description string `dynamo:"Description"`
	Done        bool   `dynamo:"Done"`
	UserID      string `dynamo:"UserID"`
	DueDateTime string `dynamo:"DueDateTime"`
	Status      string `dynamo:"Status"`
	CreatedAt   string `dynamo:"CreatedAt"`
	UpdatedAt   string `dynamo:"UpdatedAt"`
}
