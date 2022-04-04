package infrastructure

import "github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo"

func FromTodo(todo todo.Todo) todoGORM {

	descriptionValue := todo.Description()
	var descriptionToSave *string
	if descriptionToSave = &descriptionValue; descriptionValue == "" {
		descriptionToSave = nil
	}

	return todoGORM{todo.ID(), todo.Title(), descriptionToSave,
		todo.CreationDate(), todo.DueDate()}
}

func FromTodoGORM(todoGORM todoGORM) todo.Todo {
	var descriptionValue string
	if todoGORM.Description != nil {
		descriptionValue = *todoGORM.Description
	}
	var todo = todo.NewTodo(todoGORM.ID, todoGORM.Title, descriptionValue,
		todoGORM.CreationDate, todoGORM.DueDate)
	return todo
}
