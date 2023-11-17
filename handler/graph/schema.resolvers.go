package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"errors"
	"fmt"
	"go-graphql-nosql/handler/graph/model"
	"go-graphql-nosql/handler/utility"
	"time"

	"github.com/google/uuid"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.User, error) {
	fmt.Println("ログイン処理を開始します。")
	var users []*model.User
	if err := r.DB.Table("User").Scan().Filter("'Username' = ?", username).All(&users); err != nil {
		return nil, errors.New("failed to get user")
	}

	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	user := users[0]
	if !utility.CheckPasswordHash(password, user.PasswordHash) {
		return nil, ErrCodeLoginFailed
	}

	return user, nil
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	currentTime := utility.FormatDateForDynamoDB(time.Now())
	uuid := uuid.NewString()
	todo := &model.Todo{
		ID:          uuid,
		Title:       input.Title,
		Description: input.Description,
		Done:        false,
		UserID:      input.UserID,
		DueDateTime: input.DueDateTime,
		Status:      "CREATED",
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	if err := r.DB.Table("Todo").Put(todo).Run(); err != nil {
		return nil, err
	}

	return todo, nil
}

// UpdateTodoStatus is the resolver for the updateTodoStatus field.
func (r *mutationResolver) UpdateTodoStatus(ctx context.Context, id string, status string) (*model.Todo, error) {
	if err := r.DB.Table("Todo").Update("ID", id).Set("Status", status).Run(); err != nil {
		return nil, err
	}
	return &model.Todo{ID: id, Status: status}, nil
}

// UpdateTodoDone is the resolver for the updateTodoDone field.
func (r *mutationResolver) UpdateTodoDone(ctx context.Context, id string, done bool) (*model.Todo, error) {
	if err := r.DB.Table("Todo").Update("ID", id).Set("Done", done).Run(); err != nil {
		return nil, err
	}
	return &model.Todo{ID: id, Done: done}, nil
}

// DeleteTodoByID is the resolver for the deleteTodoById field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	if err := r.DB.Table("Todo").Delete("ID", id).Run(); err != nil {
		return nil, err
	}
	return &model.Todo{ID: id}, nil
}

// DeleteTodoByUserID is the resolver for the deleteTodoByUserId field.
func (r *mutationResolver) DeleteTodoByUserID(ctx context.Context, userID string) (int, error) {
	// 削除するTodoのリストを取得する。
	var todos []*model.Todo
	if err := r.DB.Table("Todo").Scan().Filter("'UserID' = ?", userID).All(&todos); err != nil {
		return 0, err
	}

	if len(todos) == 0 {
		return 0, nil
	}

	// 削除するアイテムの数をカウントする。
	deletedCount := 0
	for _, todo := range todos {
		if err := r.DB.Table("Todo").Delete("ID", todo.ID).Run(); err != nil {
			// 一つでも削除に失敗した場合はエラーを返す。
			return deletedCount, err
		}
		deletedCount++
	}

	// 削除したアイテムの数を返す。
	return deletedCount, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var users []*model.User
	r.DB.Table("User").Scan().Filter("'Username' = ?", input.Username).All(&users)
	if len(users) > 0 {
		return nil, errors.New("user already exists")
	}

	currentTime := utility.FormatDateForDynamoDB(time.Now())
	passwordHash, err := utility.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	uuid := uuid.NewString()
	user := &model.User{
		ID:           uuid,
		Username:     input.Username,
		PasswordHash: passwordHash,
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
	}

	if err := r.DB.Table("User").Put(user).Run(); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	var todos []*model.Todo
	r.DB.Table("Todo").Scan().Filter("'UserID' = ?", id).All(&todos)
	if len(todos) > 0 {
		return nil, errors.New("user still has todos")
	}

	if err := r.DB.Table("User").Delete("ID", id).Run(); err != nil {
		return nil, err
	}
	return &model.User{ID: id}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.DB.Table("Todo").Scan().All(&todos); err != nil {
		return nil, err
	}
	return todos, nil
}

// TodosByUserID is the resolver for the todosByUserId field.
func (r *queryResolver) TodosByUserID(ctx context.Context, userID string) ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.DB.Table("Todo").Scan().Filter("'UserID' = ?", userID).All(&todos); err != nil {
		return nil, err
	}
	return todos, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.DB.Table("User").Scan().All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// UserByID is the resolver for the userById field.
func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.DB.Table("User").Get("ID", id).One(&user); err != nil {
		if err.Error() == "dynamo: no item found" {
			return &model.User{}, nil
		}
		return nil, err
	}
	return &user, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	var user model.User
	if err := r.DB.Table("User").Get("ID", obj.UserID).One(&user); err != nil {
		if err.Error() == "dynamo: no item found" {
			return &model.User{}, nil
		}
		return nil, err
	}
	return &user, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
