package repository

import "github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/domain"

func FromTodo(todo domain.Todo) TodoGORM {
	return TodoGORM{todo.ID(), todo.Title(), todo.Description(),
		todo.CreationDate(), todo.DueDate()}
}

func FromTodoGORM(todoGORM TodoGORM) domain.Todo {
	todo := domain.NewTodo(todoGORM.ID, todoGORM.Title, todoGORM.Description,
		todoGORM.CreationDate, todoGORM.DueDate)
	return todo
}
