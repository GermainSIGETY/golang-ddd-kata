package todo

import (
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/validators"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	todoTitle             = "Be smart"
	todoDescription       = "even at home"
	todoDueDate     int64 = 12345
)

type creationRequestForTest struct {
	title       string
	description string
	dueDate     int64
}

func (t creationRequestForTest) Title() string {
	return t.title
}

func (t creationRequestForTest) Description() string {
	return t.description
}

func (t creationRequestForTest) DueDate() int64 {
	return t.dueDate
}

func TestCreationOk(t *testing.T) {
	title := todoTitle
	dd := todoDueDate
	desc := todoDescription

	todo, err := FromTodoCreationRequest(creationRequestForTest{title, desc, dd})
	assert.Empty(t, err)
	assert.Equal(t, todoTitle, todo.Title())
	assert.Equal(t, todoDescription, todo.Description())
	assert.Equal(t, todoDueDate, todo.DueDate().Unix())
}

func TestCreationRequestWithNilDescription(t *testing.T) {
	title := todoTitle
	dd := todoDueDate
	request, err := FromTodoCreationRequest(creationRequestForTest{title, "", dd})
	assert.Empty(t, err)
	assert.Equal(t, todoTitle, request.Title())
	assert.Empty(t, request.Description())
	assert.Equal(t, todoDueDate, request.DueDate().Unix())
}

func TestCreationRequestErrors(t *testing.T) {
	desc := fmt.Sprintf("%256v", "foo")
	request, errs := FromTodoCreationRequest(creationRequestForTest{"", desc, 0})
	assert.Empty(t, request)
	assert.Equal(t, 3, len(errs))

	errorsMap := make(map[string]DomainError)
	for _, err := range errs {
		errorsMap[err.Field()] = err
	}

	assert.Equal(t, TitleField, errorsMap[TitleField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[TitleField].Code())
	assert.Equal(t, validators.EmptyFieldDescription, errorsMap[TitleField].Description())

	assert.Equal(t, DescriptionField, errorsMap[DescriptionField].Field())
	assert.Equal(t, validators.FieldTooLongCode, errorsMap[DescriptionField].Code())

	assert.Equal(t, DueDateField, errorsMap[DueDateField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[DueDateField].Code())

}

const (
	updateTodoId                = 12
	updateTodoTitle             = "Be funny"
	updateTodoDescription       = "even at work"
	updateTodoDueDate     int64 = 123456
)

type updateRequestForTest struct {
	id          int
	title       string
	description string
	dueDate     int64
}

func (t updateRequestForTest) Id() int {
	return t.id
}

func (t updateRequestForTest) Title() string {
	return t.title
}

func (t updateRequestForTest) Description() string {
	return t.description
}

func (t updateRequestForTest) DueDate() int64 {
	return t.dueDate
}

func TestTodoFromUpdateRequestOk(t *testing.T) {
	title := updateTodoTitle
	dd := updateTodoDueDate
	desc := updateTodoDescription
	id := updateTodoId
	todo, err := FromTodoUpdateRequest(updateRequestForTest{id, title, desc, dd})
	assert.Empty(t, err)
	assert.Equal(t, updateTodoId, todo.ID())
	assert.Equal(t, updateTodoTitle, todo.Title())
	assert.Equal(t, updateTodoDescription, todo.Description())
	assert.Equal(t, time.Unix(updateTodoDueDate, 0), todo.DueDate())
}

func TestTodoFromUpdateRequestWithNilDescription(t *testing.T) {
	title := updateTodoTitle
	dd := updateTodoDueDate
	id := updateTodoId
	todo, err := FromTodoUpdateRequest(updateRequestForTest{id, title, "", dd})
	assert.Empty(t, err)
	assert.Equal(t, updateTodoId, todo.ID())
	assert.Equal(t, updateTodoTitle, todo.Title())
	assert.Empty(t, todo.Description())
	assert.Equal(t, time.Unix(updateTodoDueDate, 0), todo.DueDate())
}

func TestTodoFromUpdateRequestErrors(t *testing.T) {
	id := -1
	desc := fmt.Sprintf("%256v", "foo")
	todo, errs := FromTodoUpdateRequest(updateRequestForTest{id, "", desc, 0})
	assert.Empty(t, todo)
	assert.Equal(t, 4, len(errs))

	errorsMap := make(map[string]DomainError)
	for _, err := range errs {
		errorsMap[err.Field()] = err
	}

	assert.Equal(t, IDField, errorsMap[IDField].Field())
	assert.Equal(t, validators.InvalidNumberCode, errorsMap[IDField].Code())
	assert.Equal(t, validators.InvalidNumberDescription, errorsMap[IDField].Description())

	assert.Equal(t, TitleField, errorsMap[TitleField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[TitleField].Code())
	assert.Equal(t, validators.EmptyFieldDescription, errorsMap[TitleField].Description())

	assert.Equal(t, DescriptionField, errorsMap[DescriptionField].Field())
	assert.Equal(t, validators.FieldTooLongCode, errorsMap[DescriptionField].Code())

	assert.Equal(t, DueDateField, errorsMap[DueDateField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[DueDateField].Code())
}
