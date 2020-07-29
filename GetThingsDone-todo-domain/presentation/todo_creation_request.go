package presentation

import (
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/domain/validators"
	"time"
)

type TodoCreationRequest struct {
	title       string
	description *string
	dueDate     time.Time
}

func (t TodoCreationRequest) Title() string {
	return t.title
}

func (t TodoCreationRequest) Description() *string {
	return t.description
}

func (t TodoCreationRequest) DueDate() time.Time {
	return t.dueDate
}

func NewTodoCreationRequest(title *string, description *string, dueDate *int64) (TodoCreationRequest, []TodoDomainError) {
	var request TodoCreationRequest
	if errors := validate(title, description, dueDate); errors != nil {
		return request, errors
	}

	return TodoCreationRequest{*title, description, time.Unix(*dueDate, 0)}, nil
}

func validate(title *string, description *string, dueDate *int64) []TodoDomainError {
	var errs []TodoDomainError

	titleErr := validators.ValidateTitle(title)
	if titleErr != (validators.ValidationError{}) {
		errs = append(errs, TodoDomainError{titleField, titleErr.Code(), titleErr.Description()})
	}
	descriptionErr := validators.ValidateDescription(description)
	if descriptionErr != (validators.ValidationError{}) {
		errs = append(errs, TodoDomainError{descriptionField, descriptionErr.Code(), descriptionErr.Description()})
	}
	dueDateErr := validators.ValidateDueDate(dueDate)
	if dueDateErr != (validators.ValidationError{}) {
		errs = append(errs, TodoDomainError{dueDateField, dueDateErr.Code(), dueDateErr.Description()})
	}
	return errs
}
