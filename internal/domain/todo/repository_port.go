package todo

const (
	NoRowDeleted = "no row deleted"
)

type ITodosRepository interface {
	ReadTodo(int) (Todo, error)
	ReadTodoList() ([]SummaryResponse, error)
	Create(todo Todo) (int, error)
	UpdateTodo(todo Todo) error
	DeleteTodo(id int) error
}
