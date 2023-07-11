package core

type TodoService interface {
	CreateTodo(todo Todo) (*Todo, error)
	ReadTodo(id string) (*Todo, error)
	ReadTodos() (*[]Todo, error)
	UpdateTodo(id string) (*Todo, error)
	DeleteTodo(id string) error
}

type TodoRepository interface {
	CreateTodo(todo Todo) (*Todo, error)
	ReadTodo(id string) (*Todo, error)
	ReadTodos() (*[]Todo, error)
	UpdateTodo(id string) (*Todo, error)
	DeleteTodo(id string) error
}
