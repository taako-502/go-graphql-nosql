package graph

import "os"

func GetTodoTableName() string {
	if os.Getenv("TODO_TABLE_NAME") == "" {
		return "Todo"
	}
	return os.Getenv("TODO_TABLE_NAME")
}

func GetUserTableName() string {
	if os.Getenv("USER_TABLE_NAME") == "" {
		return "User"
	}
	return os.Getenv("USER_TABLE_NAME")
}
