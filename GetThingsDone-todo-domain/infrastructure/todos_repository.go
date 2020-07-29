package infrastructure

import (
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/domain"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/presentation"
)

const (
	NoRowDeleted = "no row deleted"
)

type ITodosRepository interface {
	ReadTodo(int) (domain.Todo, error)
	ReadTodoList() ([]presentation.TodoSummaryResponse, error)
	Create(todo domain.Todo) (int, error)
	UpdateTodo(todo domain.Todo) error
	DeleteTodo(id int) error
}
