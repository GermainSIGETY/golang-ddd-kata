package port

import "github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"

const (
	NoRowDeleted = "no row deleted"
)

type ITodosRepository interface {
	ReadTodo(int) (model.Todo, error)
	ReadTodoList() ([]model.ISummaryResponse, error)
	Create(todo model.Todo) (int, error)
	UpdateTodo(todo model.Todo) error
	DeleteTodo(id int) error
	EmptyDatabaseForTests()
	ReadTodosIdsToNotify() ([]int, error)
}
