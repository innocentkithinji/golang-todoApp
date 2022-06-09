package repository

import "todoAPI/entity"

type TodoRepository interface {
	Create(todo *entity.Todo) (*entity.Todo, error)
	Get(id string) (*entity.Todo, error)
	GetAll() ([]entity.Todo, error)
	Update(todo *entity.Todo) (*entity.Todo, error)
	Delete(id string) error
}
