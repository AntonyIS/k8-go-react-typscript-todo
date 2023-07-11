package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	return t.repo.CreateTodo(*todo)
}

func (t *TodoApplicationService) ReadTodo(id string) (*Todo, error) {
	/*
		Receives todo id from the API,fetches the data from the database.
		Returns the todo item with the id else returns an error
	*/
	return t.repo.ReadTodo(id)
}

func (t *TodoApplicationService) ReadTodos() (*[]Todo, error) {
	/*
		Reads all todos in the database
		Returns all todos from the database else returns an error
	*/
	return t.repo.ReadTodos()
}

func (t *TodoApplicationService) UpdateTodo(todo *Todo) (*Todo, error) {
	/*
		Updates the todo in the database if found
		Returns an updated todo after a successful update else returns an err
	*/
	return t.repo.UpdateTodo(todo)
}

func (t *TodoApplicationService) DeleteTodo(id string) error {
	/*
		Deletes a todo in the database with id
		Returns an error if the opetation fails
	*/
	return t.repo.DeleteTodo(id)
}

type HTTPHandlerService struct {
	svc *TodoApplicationService
}

func NewHTTPService(svc *TodoApplicationService) HTTPHandler {
	return HTTPHandlerService{
		svc: svc,
	}
}

func (h HTTPHandlerService) PostTodo(ctx *gin.Context) {
	var todo Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.svc.CreateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
	return
}

func (h HTTPHandlerService) GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.svc.ReadTodo(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
		})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
	return
}
func (h HTTPHandlerService) GetTodos(ctx *gin.Context) {
	users, err := h.svc.ReadTodos()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
	return
}

func (h HTTPHandlerService) UpdateTodo(ctx *gin.Context) {
	var todo Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.svc.UpdateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
	return
}
func (h HTTPHandlerService) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.svc.DeleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "todo deleted successfully",
	})
	return
}
