package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/validators"
	"github.com/stretchr/testify/assert"
)

const (
	todoID                = 12
	todoTitle             = "Be smart"
	todoDescription       = "even at home"
	todoDueDate     int64 = 12345
)

type creationRequestForTest struct {
	title       string
	description string
	dueDate     int64
	assignee    string
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

func (t creationRequestForTest) Assignee() string {
	return t.assignee
}

func Test_Creation(t *testing.T) {
	tests := []struct {
		name         string
		given        creationRequestForTest
		expected     model.Todo
		expectedErrs []model.DomainError
	}{
		{
			name: "nominal case",
			given: creationRequestForTest{
				title:       todoTitle,
				description: todoDescription,
				dueDate:     todoDueDate,
			},
			expected: model.Todo{
				Title:       todoTitle,
				Description: todoDescription,
				DueDate:     time.Unix(todoDueDate, 0),
			},
			expectedErrs: nil,
		},
		{
			name: "empty description",
			given: creationRequestForTest{
				title:       todoTitle,
				description: "",
				dueDate:     todoDueDate,
			},
			expected: model.Todo{
				Title:       todoTitle,
				Description: "",
				DueDate:     time.Unix(todoDueDate, 0),
			},
			expectedErrs: nil,
		},
		{
			name: "validation errors",
			given: creationRequestForTest{
				title:       "",
				description: fmt.Sprintf("%256v", "foo"),
				dueDate:     0,
			},
			expected: model.Todo{},
			expectedErrs: []model.DomainError{
				model.NewDomainError(model.TitleField, validators.EmptyFieldCode, validators.EmptyFieldDescription),
				model.NewDomainError(model.DescriptionField, validators.FieldTooLongCode, fmt.Sprintf(validators.FieldToLongDescription, 255)),
				model.NewDomainError(model.DueDateField, validators.EmptyFieldCode, validators.EmptyFieldDescription),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FromTodoCreationRequest(tt.given)
			if tt.expectedErrs == nil {
				assert.Empty(t, err)
				assert.Equal(t, tt.expected.Title, result.Title)
				assert.Equal(t, tt.expected.Description, result.Description)
				assert.Equal(t, tt.expected.DueDate.Unix(), result.DueDate.Unix())
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, len(tt.expectedErrs), len(err))
				errorsMap := make(map[string]model.DomainError)
				for _, err := range err {
					errorsMap[err.Field()] = err
				}
				for _, expectedErr := range tt.expectedErrs {
					actualErr := errorsMap[expectedErr.Field()]
					assert.NotNil(t, actualErr)
					assert.Equal(t, expectedErr.Field(), actualErr.Field())
					assert.Equal(t, expectedErr.Code(), actualErr.Code())
					assert.Equal(t, expectedErr.Description(), actualErr.Description())
				}
			}
		})
	}
}

type updateRequestForTest struct {
	id          int
	title       string
	description string
	dueDate     int64
	assignee    string
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

func (t updateRequestForTest) Assignee() string {
	return t.assignee
}

func Test_Update(t *testing.T) {
	tests := []struct {
		name         string
		given        updateRequestForTest
		expectedErrs []model.DomainError
	}{
		{
			name: "nominal case",
			given: updateRequestForTest{
				id:          todoID,
				title:       todoTitle,
				description: todoDescription,
				dueDate:     todoDueDate,
			},
			expectedErrs: nil,
		},
		{
			name: "empty description",
			given: updateRequestForTest{
				id:          todoID,
				title:       todoTitle,
				description: "",
				dueDate:     todoDueDate,
			},
			expectedErrs: nil,
		},
		{
			name: "validation errors",
			given: updateRequestForTest{
				id:          -1,
				title:       "",
				description: fmt.Sprintf("%256v", "foo"),
				dueDate:     0,
			},
			expectedErrs: []model.DomainError{
				model.NewDomainError(model.IDField, validators.InvalidNumberCode, validators.InvalidNumberDescription),
				model.NewDomainError(model.TitleField, validators.EmptyFieldCode, validators.EmptyFieldDescription),
				model.NewDomainError(model.DescriptionField, validators.FieldTooLongCode, fmt.Sprintf(validators.FieldToLongDescription, 255)),
				model.NewDomainError(model.DueDateField, validators.EmptyFieldCode, validators.EmptyFieldDescription),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := ValidateUpdateRequest(tt.given)
			if len(tt.expectedErrs) == 0 {
				assert.Equal(t, 0, len(errs))
			} else {
				assert.Equal(t, len(tt.expectedErrs), len(errs))
				errorsMap := make(map[string]model.DomainError)
				for _, err := range errs {
					errorsMap[err.Field()] = err
				}
				for _, expectedErr := range tt.expectedErrs {
					actualErr := errorsMap[expectedErr.Field()]
					assert.NotNil(t, actualErr)
					assert.Equal(t, expectedErr.Field(), actualErr.Field())
					assert.Equal(t, expectedErr.Code(), actualErr.Code())
					assert.Equal(t, expectedErr.Description(), actualErr.Description())
				}
			}
		})
	}
}
