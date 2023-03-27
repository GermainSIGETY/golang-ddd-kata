package infrastructure

import (
	"testing"
	"time"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/stretchr/testify/assert"
)

var (
	id          = 33
	title       = "tester c'est douter"
	description = "mais sans c'est galérer"
	assignee    = "A poor burnout-ed guy"
)

var creationDate = time.Date(2013, 3, 1, 12, 30, 0, 0, time.UTC)
var dueDate = time.Date(2031, 3, 1, 12, 30, 0, 0, time.UTC)

func TestFromTodo(t *testing.T) {
	todo := model.Todo{
		ID: id,

		CreationDate:     creationDate,
		Description:      description,
		DueDate:          dueDate,
		Title:            title,
		Assignee:         assignee,
		NotificationSent: true,
	}
	todoGORM := FromTodo(todo)

	assert.Equal(t, id, todoGORM.ID, "The two IDs should be the same.")
	assert.Equal(t, title, todoGORM.Title, "The two titles should be the same.")
	assert.Equal(t, description, *todoGORM.Description, "The two descriptions should be the same.")
	assert.Equal(t, creationDate, todoGORM.CreationDate, "The two creation dates should be the same.")
	assert.Equal(t, dueDate, todoGORM.DueDate, "The two due dates should be the same.")
	assert.Equal(t, assignee, *todoGORM.Assignee, "The two assignee should be the same.")
	assert.Truef(t, todoGORM.NotificationSent, "NotificationSent should be true")
}

func TestFromTodoWithEmptyValues(t *testing.T) {
	todo := model.Todo{
		ID: 0,

		CreationDate: creationDate,
		Description:  "",
		DueDate:      dueDate,
		Title:        title,
	}

	todoGORM := FromTodo(todo)

	assert.Equal(t, title, todoGORM.Title, "The two titles should be the same.")
	assert.Empty(t, todoGORM.ID, "The two IDs should be the same.")
	assert.Nil(t, todoGORM.Description, "The two descriptions should be the same.")
	assert.Nil(t, todoGORM.Assignee, "The two Assignee should be the same.")
}

func TestFromTodoGORM(t *testing.T) {
	todoGORM := todoGORM{id, title, &description, creationDate, dueDate, &assignee, true}

	todo := FromTodoGORM(todoGORM)

	assert.Equal(t, id, todo.ID)
	assert.Equal(t, title, todo.Title)
	assert.Equal(t, description, todo.Description)
	assert.Equal(t, creationDate, todo.CreationDate)
	assert.Equal(t, dueDate, todo.DueDate)
	assert.Equal(t, assignee, todo.Assignee)
	assert.True(t, todo.NotificationSent)
}

func TestFromTodoGORMWithEmptyValues(t *testing.T) {
	todoGORM := todoGORM{Title: title, CreationDate: creationDate, DueDate: dueDate}

	todo := FromTodoGORM(todoGORM)

	assert.Equal(t, title, todo.Title)
	assert.Empty(t, todo.ID)
	assert.Empty(t, todo.Description)
	assert.Empty(t, todo.Assignee)
	assert.False(t, todo.NotificationSent)
}
