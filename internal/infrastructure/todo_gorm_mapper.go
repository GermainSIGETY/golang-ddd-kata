package infrastructure

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
)

func FromTodo(todo model.Todo) todoGORM {

	descriptionValue := todo.Description
	var descriptionToSave *string
	if descriptionToSave = &descriptionValue; descriptionValue == "" {
		descriptionToSave = nil
	}

	return todoGORM{todo.ID, todo.Title, descriptionToSave,
		todo.CreationDate, todo.DueDate}
}

func FromTodoGORM(todoGORM todoGORM) model.Todo {
	var descriptionValue string
	if todoGORM.Description != nil {
		descriptionValue = *todoGORM.Description
	}

	todo := model.Todo{
		ID: todoGORM.ID,

		CreationDate: todoGORM.CreationDate,
		Description:  descriptionValue,
		DueDate:      todoGORM.DueDate,
		Title:        todoGORM.Title,
	}
	return todo
}
