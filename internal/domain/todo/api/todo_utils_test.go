package api

import (
	"fmt"
	"testing"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/validators"
	"github.com/stretchr/testify/assert"
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
	fmt.Printf("todo: %v\n", todo)
	// Je préfère utiliser des valeur directement plutôt que des const dans les "expected" d'error
	// exemple: assert.Equal(t, "Be smart", todo.Title())
	assert.Equal(t, todoTitle, todo.Title)
	assert.Equal(t, todoDescription, todo.Description)
	assert.Equal(t, todoDueDate, todo.DueDate.Unix())
}

func TestCreationRequestWithNilDescription(t *testing.T) {
	title := todoTitle
	dd := todoDueDate
	request, err := FromTodoCreationRequest(creationRequestForTest{title, "", dd})
	assert.Empty(t, err)
	assert.Equal(t, todoTitle, request.Title)
	assert.Empty(t, request.Description)
	assert.Equal(t, todoDueDate, request.DueDate.Unix())
}

func TestCreationRequestErrors(t *testing.T) {
	desc := fmt.Sprintf("%256v", "foo")
	request, errs := FromTodoCreationRequest(creationRequestForTest{"", desc, 0})
	assert.Empty(t, request)
	assert.Equal(t, 3, len(errs))

	errorsMap := make(map[string]model.DomainError)
	for _, err := range errs {
		errorsMap[err.Field()] = err
	}

	assert.Equal(t, model.TitleField, errorsMap[model.TitleField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[model.TitleField].Code())
	assert.Equal(t, validators.EmptyFieldDescription, errorsMap[model.TitleField].Description())

	assert.Equal(t, model.DescriptionField, errorsMap[model.DescriptionField].Field())
	assert.Equal(t, validators.FieldTooLongCode, errorsMap[model.DescriptionField].Code())

	assert.Equal(t, model.DueDateField, errorsMap[model.DueDateField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[model.DueDateField].Code())

}
