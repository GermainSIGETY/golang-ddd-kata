package domain

import (
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/presentation"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	id          = 1234
	title       = "Get a Job"
	description = "In a team that develops in GO"
)

var (
	creationDate = time.Date(2015, time.November, 21, 0, 0, 0, 0, time.UTC)
	dueDate      = time.Date(2015, time.November, 24, 0, 0, 0, 0, time.UTC)
)

func TestTodo_create(t *testing.T) {
	todo := Todo{&id, title, &description, creationDate, dueDate}
	assert.Equal(t, title, todo.Title())
	assert.Equal(t, description, *todo.Description())
	assert.Equal(t, creationDate, todo.CreationDate())
	assert.Equal(t, dueDate, todo.DueDate())
}

func TestTodo_update(t *testing.T) {
	todo := Todo{&id, title, &description, creationDate, dueDate}

	newTitle := "get quickly a new job"
	newDescription := "in a team that tell jokes"
	newDueDate := time.Date(2066, time.November, 24, 0, 0, 0, 0, time.UTC)
	newDueDateAsInt := newDueDate.Unix()

	updateRequest, _ := presentation.NewTodoUpdateRequest(id, &newTitle, &newDescription, &newDueDateAsInt)
	todo.UpdateFromTodoUpdateRequest(updateRequest)

	assert.Equal(t, newTitle, todo.Title())
	assert.Equal(t, newDescription, *todo.Description())
	assert.True(t, newDueDate.Equal(todo.DueDate()))
}

func TestTodo_updateNulDescription(t *testing.T) {
	todo := Todo{&id, title, &description, creationDate, dueDate}

	dueDAteInt := todo.dueDate.Unix()

	updateRequest, _ := presentation.NewTodoUpdateRequest(id, &todo.title, nil, &dueDAteInt)
	todo.UpdateFromTodoUpdateRequest(updateRequest)

	assert.Nil(t, todo.Description())
}
