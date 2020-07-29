package presentation

import (
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/domain/validators"
	"time"
)

type TodoUpdateRequest struct {
	id          int
	title       string
	description *string
	dueDate     time.Time
}

func (t TodoUpdateRequest) Id() int {
	return t.id
}

func (t TodoUpdateRequest) Title() string {
	return t.title
}

func (t TodoUpdateRequest) Description() *string {
	return t.description
}

func (t TodoUpdateRequest) DueDate() time.Time {
	return t.dueDate
}

func NewTodoUpdateRequest(id int, title *string, description *string, dueDate *int64) (TodoUpdateRequest, []TodoDomainError) {
	var request TodoUpdateRequest
	if errors := validateUpdateRequest(id, title, description, dueDate); errors != nil {
		return request, errors
	}

	return TodoUpdateRequest{id, *title, description, time.Unix(*dueDate, 0)}, nil
}

func validateUpdateRequest(id int, title *string, description *string, dueDate *int64) []TodoDomainError {
	var errs []TodoDomainError

	idErr := validators.ValidateTodoId(id)
	if idErr != (validators.ValidationError{}) {
		errs = append(errs, TodoDomainError{IDField, idErr.Code(), idErr.Description()})
	}
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
