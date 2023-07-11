package core

type TodoApplicationService struct {
	repo TodoRepository
}

func NewTodoService(repo *TodoRepository) *TodoApplicationService {
	return &TodoApplicationService{
		repo: *repo,
	}
}

func (t *TodoApplicationService) CreateTodo(todo *Todo) (*Todo, error) {
	/*
		Receives todo item from the API and saves the data into the database.
		This method will add an ID and created at attribute values
		Returns the todo item after successful creation , else returns an error
	*/
	return nil, nil
}

func (t *TodoApplicationService) ReadTodo(id string) (*Todo, error) {
	/*
		Receives todo id from the API,fetches the data from the database.
		Returns the todo item with the id else returns an error
	*/
	return nil, nil
}

func (t *TodoApplicationService) ReadTodos() (*[]Todo, error) {
	/*
		Reads all todos in the database
		Returns all todos from the database else returns an error
	*/
	return nil, nil
}

func (t *TodoApplicationService) UpdateTodo(todo *Todo) (*[]Todo, error) {
	/*
		Updates the todo in the database if found
		Returns an updated todo after a successful update else returns an err
	*/
	return nil, nil
}

func (t *TodoApplicationService) DeleteTodo(id string) error {
	/*
		Deletes a todo in the database with id
		Returns an error if the opetation fails
	*/
	return nil
}
