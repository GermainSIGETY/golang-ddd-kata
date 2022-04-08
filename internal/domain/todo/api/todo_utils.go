package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/validators"
	"github.com/stretchr/testify/assert"
)

func UpdateFromTodoUpdateRequest(t *model.Todo, request port.UpdateRequest) {
	// Choix model.Todo -> champ externe
	t.Title = request.Title()
	t.Description = request.Description()
	t.DueDate = time.Unix(request.DueDate(), 0)
}

func FromTodoCreationRequest(request port.CreationRequest) (model.Todo, []model.DomainError) {
	if errors := validateCreationRequest(request); errors != nil {
		return model.Todo{}, errors
	}
	return model.Todo{
		Title:        request.Title(),
		Description:  request.Description(),
		CreationDate: time.Now(),
		DueDate:      time.Unix(request.DueDate(), 0),
	}, nil
}

func validateCreationRequest(request port.CreationRequest) []model.DomainError {
	var errs []model.DomainError

	titleErr := validators.ValidateTitle(request.Title())
	if titleErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewTodoDomainError(model.TitleField, titleErr.Code(), titleErr.Description()))
	}
	descriptionErr := validators.ValidateDescription(request.Description())
	if descriptionErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewTodoDomainError(model.DescriptionField, descriptionErr.Code(), descriptionErr.Description()))
	}
	dueDateErr := validators.ValidateDueDate(request.DueDate())
	if dueDateErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewTodoDomainError(model.DueDateField, dueDateErr.Code(), dueDateErr.Description()))
	}
	return errs
}

func FromTodoUpdateRequest(request port.UpdateRequest) (model.Todo, []model.DomainError) {
	if errors := validateUpdateRequest(request); errors != nil {
		return model.Todo{}, errors
	}
	// wait Todo immutable ou champ exporté
	return model.Todo{
		ID: request.Id(),

		CreationDate: time.Now(),
		Description:  request.Description(),
		DueDate:      time.Unix(request.DueDate(), 0),
		Title:        request.Title(),
	}, nil
}

func validateUpdateRequest(request port.UpdateRequest) []model.DomainError {
	var errs []model.DomainError

	idErr := validators.ValidateTodoId(request.Id())
	if idErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewTodoDomainError(model.IDField, idErr.Code(), idErr.Description()))
	}
	titleErr := validators.ValidateTitle(request.Title())
	if titleErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewTodoDomainError(model.TitleField, titleErr.Code(), titleErr.Description()))
	}
	descriptionErr := validators.ValidateDescription(request.Description())
	if descriptionErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewTodoDomainError(model.DescriptionField, descriptionErr.Code(), descriptionErr.Description()))
	}
	dueDateErr := validators.ValidateDueDate(request.DueDate())
	if dueDateErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewTodoDomainError(model.DueDateField, dueDateErr.Code(), dueDateErr.Description()))
	}
	return errs
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
	assert.Equal(t, updateTodoId, todo.ID)
	assert.Equal(t, updateTodoTitle, todo.Title)
	assert.Equal(t, updateTodoDescription, todo.Description)
	assert.Equal(t, time.Unix(updateTodoDueDate, 0), todo.DueDate)
}

func TestTodoFromUpdateRequestWithNilDescription(t *testing.T) {
	title := updateTodoTitle
	dd := updateTodoDueDate
	id := updateTodoId
	todo, err := FromTodoUpdateRequest(updateRequestForTest{id, title, "", dd})
	assert.Empty(t, err)
	assert.Equal(t, updateTodoId, todo.ID)
	assert.Equal(t, updateTodoTitle, todo.Title)
	assert.Empty(t, todo.Description)
	assert.Equal(t, time.Unix(updateTodoDueDate, 0), todo.DueDate)
}

func TestTodoFromUpdateRequestErrors(t *testing.T) {
	id := -1
	desc := fmt.Sprintf("%256v", "foo")
	todo, errs := FromTodoUpdateRequest(updateRequestForTest{id, "", desc, 0})
	assert.Empty(t, todo)
	assert.Equal(t, 4, len(errs))

	errorsMap := make(map[string]model.DomainError)
	for _, err := range errs {
		errorsMap[err.Field()] = err
	}

	assert.Equal(t, model.IDField, errorsMap[model.IDField].Field())
	assert.Equal(t, validators.InvalidNumberCode, errorsMap[model.IDField].Code())
	assert.Equal(t, validators.InvalidNumberDescription, errorsMap[model.IDField].Description())

	assert.Equal(t, model.TitleField, errorsMap[model.TitleField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[model.TitleField].Code())
	assert.Equal(t, validators.EmptyFieldDescription, errorsMap[model.TitleField].Description())

	assert.Equal(t, model.DescriptionField, errorsMap[model.DescriptionField].Field())
	assert.Equal(t, validators.FieldTooLongCode, errorsMap[model.DescriptionField].Code())

	assert.Equal(t, model.DueDateField, errorsMap[model.DueDateField].Field())
	assert.Equal(t, validators.EmptyFieldCode, errorsMap[model.DueDateField].Code())
}
