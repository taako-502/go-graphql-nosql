package graph

import "os"

func GetTodoTableName() string {
	if os.Getenv("TODO_TABLE_NAME") == "" {
		return "Todo"
	}
	return os.Getenv("TODO_TABLE_NAME")
}

func GetUserTableName() string {
	if os.Getenv("UserTableName") == "" {
		return "Todo"
	}
	return os.Getenv("UserTableName")
}
