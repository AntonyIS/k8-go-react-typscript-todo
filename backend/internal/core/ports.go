package core

import "github.com/gin-gonic/gin"

type TodoService interface {
	CreateTodo(todo *Todo) (*Todo, error)
	ReadTodo(id string) (*Todo, error)
	ReadTodos() (*[]Todo, error)
	UpdateTodo(todo Todo) (*Todo, error)
	DeleteTodo(id string) error
}

type TodoRepository interface {
	CreateTodo(todo *Todo) (*Todo, error)
	ReadTodo(id string) (*Todo, error)
	ReadTodos() (*[]Todo, error)
	UpdateTodo(todo *Todo) (*Todo, error)
	DeleteTodo(id string) error
}

type HTTPHandler interface {
	PostTodo(ctx *gin.Context)
	GetTodo(ctx *gin.Context)
	GetTodos(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
}
