package infrastructure

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
)

func FromTodo(todo model.Todo) todoGORM {
	var descriptionToSave, assigneeToSave *string
	if descriptionToSave = &todo.Description; todo.Description == "" {
		descriptionToSave = nil
	}
	if assigneeToSave = &todo.Assignee; todo.Assignee == "" {
		assigneeToSave = nil
	}

	return todoGORM{todo.ID, todo.Title, descriptionToSave,
		todo.CreationDate, todo.DueDate, assigneeToSave, todo.NotificationSent}
}

func FromTodoGORM(todoGORM todoGORM) model.Todo {
	var descriptionValue, assigneeValue string
	if todoGORM.Description != nil {
		descriptionValue = *todoGORM.Description
	}
	if todoGORM.Assignee != nil {
		assigneeValue = *todoGORM.Assignee
	}

	todo := model.Todo{
		ID: todoGORM.ID,

		CreationDate:     todoGORM.CreationDate,
		Description:      descriptionValue,
		DueDate:          todoGORM.DueDate,
		Title:            todoGORM.Title,
		Assignee:         assigneeValue,
		NotificationSent: todoGORM.NotificationSent,
	}
	return todo
}
