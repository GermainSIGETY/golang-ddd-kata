package repository

import (
	"github.com/GermainSIGETY/golang-ddd-kata/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	id          = 33
	title       = "tester c'est douter"
	description = "mais sans c'est galérer"
)

var creationDate = time.Date(2013, 3, 1, 12, 30, 0, 0, time.UTC)
var dueDate = time.Date(2031, 3, 1, 12, 30, 0, 0, time.UTC)

func TestFromTodo(t *testing.T) {
	todo := domain.NewTodo(&id, title, &description, creationDate, dueDate)

	todoGORM := FromTodo(todo)

	assert.Equal(t, id, *todoGORM.ID, "The two IDs should be the same.")
	assert.Equal(t, title, todoGORM.Title, "The two titles should be the same.")
	assert.Equal(t, description, *todoGORM.Description, "The two descriptions should be the same.")
	assert.Equal(t, creationDate, todoGORM.CreationDate, "The two creation dates should be the same.")
	assert.Equal(t, dueDate, todoGORM.DueDate, "The two due dates should be the same.")
}

func TestFromTodoWithNilValues(t *testing.T) {
	todo := domain.NewTodo(nil, title, nil, creationDate, dueDate)

	todoGORM := FromTodo(todo)

	assert.Equal(t, title, todoGORM.Title, "The two titles should be the same.")
	assert.Nil(t, todoGORM.ID, "The two IDs should be the same.")
	assert.Nil(t, todoGORM.Description, "The two descriptions should be the same.")
}

func TestFromTodoGORM(t *testing.T) {
	todoGORM := todoGORM{&id, title, &description, creationDate, dueDate}

	todo := FromTodoGORM(todoGORM)

	assert.Equal(t, id, *todo.ID())
	assert.Equal(t, title, todo.Title())
	assert.Equal(t, description, *todo.Description())
	assert.Equal(t, creationDate, todo.CreationDate())
	assert.Equal(t, dueDate, todo.DueDate())
}

func TestFromTodoGORMWithNilvalues(t *testing.T) {
	todoGORM := todoGORM{Title: title, CreationDate: creationDate, DueDate: dueDate}

	todo := FromTodoGORM(todoGORM)

	assert.Equal(t, title, todo.Title())
	assert.Nil(t, todo.ID())
	assert.Nil(t, todo.Description())
}
