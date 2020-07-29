package domain

import (
	. "github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/presentation"
	"time"
)

type Todo struct {
	id           *int
	title        string
	description  *string
	creationDate time.Time
	dueDate      time.Time
}

func NewTodo(ID *int, title string, description *string, creationDate time.Time, dueDate time.Time) Todo {
	return Todo{ID, title, description, creationDate, dueDate}
}

func FromTodoCreationRequest(request TodoCreationRequest) Todo {
	return Todo{title: request.Title(), description: request.Description(), creationDate: time.Now(), dueDate: request.DueDate()}
}

func (t Todo) MapToTodoResponse() ReadTodoResponse {
	if t.id == nil {
		return ReadTodoResponse{}
	}
	return ReadTodoResponse{
		ID:           *t.id,
		Title:        t.title,
		Description:  t.description,
		CreationDate: t.creationDate,
		DueDate:      t.dueDate,
	}
}

func (t Todo) ID() *int {
	return t.id
}

func (t Todo) Title() string {
	return t.title
}

func (t Todo) Description() *string {
	return t.description
}

func (t Todo) CreationDate() time.Time {
	return t.creationDate
}

func (t Todo) DueDate() time.Time {
	return t.dueDate
}

func (t *Todo) UpdateFromTodoUpdateRequest(request TodoUpdateRequest) {
	t.title = request.Title()
	t.description = request.Description()
	t.dueDate = request.DueDate()
}
