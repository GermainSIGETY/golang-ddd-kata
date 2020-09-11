package repository

import "github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/domain"

func FromTodo(todo domain.Todo) todoGORM {
	return todoGORM{todo.ID(), todo.Title(), todo.Description(),
		todo.CreationDate(), todo.DueDate()}
}

func FromTodoGORM(todoGORM todoGORM) domain.Todo {
	todo := domain.NewTodo(todoGORM.ID, todoGORM.Title, todoGORM.Description,
		todoGORM.CreationDate, todoGORM.DueDate)
	return todo
}
