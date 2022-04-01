package presentation

import (
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/src/domain/validators"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	updateTodoId                = 12
	updateTodoTitle             = "Be funny"
	updateTodoDescription       = "even at work"
	updateTodoDueDate     int64 = 123456
)

func TestUpdateRequestOk(t *testing.T) {
	title := updateTodoTitle
	dd := updateTodoDueDate
	desc := updateTodoDescription
	request, err := NewTodoUpdateRequest(updateTodoId, &title, &desc, &dd)
	assert.Empty(t, err)
	assert.Equal(t, updateTodoId, request.id)
	assert.Equal(t, updateTodoTitle, request.title)
	assert.Equal(t, updateTodoDescription, *request.description)
	assert.Equal(t, updateTodoDueDate, request.dueDate.Unix())
}

func TestUpdateRequestWithNilDescription(t *testing.T) {
	title := updateTodoTitle
	dd := updateTodoDueDate
	request, err := NewTodoUpdateRequest(updateTodoId, &title, nil, &dd)
	assert.Empty(t, err)
	assert.Equal(t, updateTodoId, request.id)
	assert.Equal(t, updateTodoTitle, request.title)
	assert.Nil(t, request.description)
	assert.Equal(t, updateTodoDueDate, request.dueDate.Unix())
}

func TestUpdateRequestErrors(t *testing.T) {
	desc := fmt.Sprintf("%256v", "foo")
	request, errs := NewTodoUpdateRequest(-1, nil, &desc, nil)
	assert.Empty(t, request)
	assert.Equal(t, 4, len(errs))

	errorsMap := make(map[string]TodoDomainError)
	for _, err := range errs {
		errorsMap[err.Field()] = err
	}

	assert.Equal(t, IDField, errorsMap[IDField].Field())
	assert.Equal(t, validators.InvalidNumberCode, errorsMap[IDField].Code())
	assert.Equal(t, validators.InvalidNumberDescription, errorsMap[IDField].Description())

	assert.Equal(t, titleField, errorsMap[titleField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[titleField].Code())
	assert.Equal(t, validators.EmptyFieldDescription, errorsMap[titleField].Description())

	assert.Equal(t, descriptionField, errorsMap[descriptionField].Field())
	assert.Equal(t, validators.FieldTooLongCode, errorsMap[descriptionField].Code())

	assert.Equal(t, dueDateField, errorsMap[dueDateField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[dueDateField].Code())
}
