package service

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"todoAPI/entity"
	"todoAPI/repository"
)

type TodoService interface {
	Validate(todo *entity.Todo) error
	Create(todo *entity.Todo) (*entity.Todo, error)
	RetrieveTodo(id string) (*entity.Todo, error)
	RetrieveAll() ([]entity.Todo, error)
	UpdateTodo(id string, todo *entity.Todo) (*entity.Todo, error)
	DeleteTodo(id string) error
}
type service struct{}

var repo repository.TodoRepository

func (s service) Validate(todo *entity.Todo) error {
	if todo != nil {
		err := errors.New("the post is empty")
		return err
	}
	if todo.Title == "" {
		err := errors.New("todo is missing a title")
		return err
	}
	return nil
}

func (s service) Create(todo *entity.Todo) (*entity.Todo, error) {
	todo.Id = uuid.New().String()
	todo.Done = false
	return repo.Create(todo)
}

func (s service) RetrieveTodo(id string) (*entity.Todo, error) {
	return repo.Get(id)
}

func (s service) RetrieveAll() ([]entity.Todo, error) {

	todos, err := repo.GetAll()
	if err != nil {
		log.Println("Couldn't retrieve all todos")
	}
	return todos, err
}

func (s service) UpdateTodo(id string, todo *entity.Todo) (*entity.Todo, error) {
	oldtodo, err := repo.Get(id)
	if err != nil {
		log.Printf("Could not find todo with ID: %s", id)
		return nil, errors.New("couldn't find the request todo")
	}

	if todo.Title != "" {
		oldtodo.Title = todo.Title
	}
	if todo.Description != "" {
		oldtodo.Description = todo.Description
	}
	if todo.Done != oldtodo.Done {
		oldtodo.Done = todo.Done
	}
	return repo.Update(oldtodo)
}

func (s service) DeleteTodo(id string) error {
	_, err := repo.Get(id)
	if err != nil {
		log.Printf("Could not find todo with ID: %s", id)
		return errors.New("couldn't find the request todo")
	}
	return repo.Delete(id)
}

func NewTodoService(serviceRepo repository.TodoRepository) TodoService {
	repo = serviceRepo
	return &service{}
}
