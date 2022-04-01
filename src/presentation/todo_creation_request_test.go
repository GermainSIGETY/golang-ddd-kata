package presentation

import (
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/src/domain/validators"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	todoTitle             = "Be smart"
	todoDescription       = "even at home"
	todoDueDate     int64 = 12345
)

func TestCreationRequestOk(t *testing.T) {
	title := todoTitle
	dd := todoDueDate
	desc := todoDescription
	request, err := NewTodoCreationRequest(&title, &desc, &dd)
	assert.Empty(t, err)
	assert.Equal(t, todoTitle, request.title)
	assert.Equal(t, todoDescription, *request.description)
	assert.Equal(t, todoDueDate, request.dueDate.Unix())
}

func TestCreationRequestWithNilDescription(t *testing.T) {
	title := todoTitle
	dd := todoDueDate
	request, err := NewTodoCreationRequest(&title, nil, &dd)
	assert.Empty(t, err)
	assert.Equal(t, todoTitle, request.title)
	assert.Nil(t, request.description)
	assert.Equal(t, todoDueDate, request.dueDate.Unix())
}

func TestCreationRequestErrors(t *testing.T) {
	desc := fmt.Sprintf("%256v", "foo")
	request, errs := NewTodoCreationRequest(nil, &desc, nil)
	assert.Empty(t, request)
	assert.Equal(t, 3, len(errs))

	errorsMap := make(map[string]TodoDomainError)
	for _, err := range errs {
		errorsMap[err.Field()] = err
	}

	assert.Equal(t, titleField, errorsMap[titleField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[titleField].Code())
	assert.Equal(t, validators.EmptyFieldDescription, errorsMap[titleField].Description())

	assert.Equal(t, descriptionField, errorsMap[descriptionField].Field())
	assert.Equal(t, validators.FieldTooLongCode, errorsMap[descriptionField].Code())

	assert.Equal(t, dueDateField, errorsMap[dueDateField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[dueDateField].Code())

}
