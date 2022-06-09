package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"todoAPI/entity"
	"todoAPI/errors"
	"todoAPI/service"
)

type TodoController interface {
	CreateTodo(c echo.Context) error
	RetrieveTodo(c echo.Context) error
	RetrieveAll(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type controller struct{}

var todoService service.TodoService

func NewTodoController(service service.TodoService) TodoController {
	todoService = service
	return &controller{}
}

func (*controller) CreateTodo(c echo.Context) error {
	var todo = new(entity.Todo)
	if err := c.Bind(todo); err != nil {
		log.Printf("Failed to Bind request data to user data: %v", err)
		return c.JSON(
			http.StatusBadRequest,
			errors.ServiceError{Message: "Error binding the todo"})
	}

	create, err := todoService.Create(todo)
	if err != nil {
		log.Printf("Failed to Create your todo: %v", err)
		return c.JSON(
			http.StatusInternalServerError,
			errors.ServiceError{Message: "Error Creating your todo"})
	}

	return c.JSON(http.StatusCreated, create)
}

func (*controller) RetrieveTodo(c echo.Context) error {
	todoID := c.Param("id")
	todo, err := todoService.RetrieveTodo(todoID)
	if err != nil || todo == nil {
		log.Printf("Failed to Retrieve your todo: %v", err)
		return c.JSON(
			http.StatusNotFound,
			errors.ServiceError{Message: "Couldn't find your todo"})
	}

	return c.JSON(http.StatusOK, todo)
}

func (*controller) RetrieveAll(c echo.Context) error {

	todos, err := todoService.RetrieveAll()
	if err != nil {
		log.Printf("Failed to Retrieve all todo: %v", err)
		return c.JSON(
			http.StatusNotFound,
			errors.ServiceError{Message: "Couldn't find your todo"})
	}

	return c.JSON(http.StatusOK, todos)
}

func (*controller) UpdateTodo(c echo.Context) error {
	todoID := c.Param("id")
	var todo = new(entity.Todo)
	if err := c.Bind(todo); err != nil {
		log.Printf("Failed to Bind request data to user data: %v", err)
		return c.JSON(
			http.StatusBadRequest,
			errors.ServiceError{Message: "Error binding the todo"})
	}

	updatedTodo, err := todoService.UpdateTodo(todoID, todo)
	if err != nil {
		log.Printf("Failed to Update your todo: %v", err)
		return c.JSON(
			http.StatusInternalServerError,
			errors.ServiceError{Message: "Error Updating your todo"})
	}

	return c.JSON(http.StatusAccepted, updatedTodo)
}

func (*controller) DeleteTodo(c echo.Context) error {
	todoID := c.Param("id")

	err := todoService.DeleteTodo(todoID)
	if err != nil {
		log.Printf("Failed to Delete your todo: %v", err)
		return c.JSON(
			http.StatusInternalServerError,
			errors.ServiceError{Message: "Error Deleting your todo"})
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{"message": fmt.Sprintf("Deleted %s", todoID)})
}
