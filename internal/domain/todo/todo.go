package todo

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/validators"
	"time"
)

type Todo struct {
	id           int
	title        string
	description  string
	creationDate time.Time
	dueDate      time.Time
}

func (t Todo) MapToTodoResponse() ReadTodoResponse {
	if t.id == 0 {
		return ReadTodoResponse{}
	}
	return ReadTodoResponse{
		ID:           t.id,
		Title:        t.title,
		Description:  t.description,
		CreationDate: t.creationDate,
		DueDate:      t.dueDate,
	}
}

func (t Todo) ID() int {
	return t.id
}

func (t Todo) Title() string {
	return t.title
}

func (t Todo) Description() string {
	return t.description
}

func (t Todo) CreationDate() time.Time {
	return t.creationDate
}

func (t Todo) DueDate() time.Time {
	return t.dueDate
}

func NewTodo(id int, title string, description string, creationDate time.Time, dueDate time.Time) Todo {
	return Todo{id, title, description, creationDate, dueDate}
}

func (t *Todo) UpdateFromTodoUpdateRequest(request UpdateRequest) {
	t.title = request.Title()
	t.description = request.Description()
	t.dueDate = time.Unix(request.DueDate(), 0)
}

func FromTodoCreationRequest(request CreationRequest) (Todo, []DomainError) {
	if errors := validateCreationRequest(request); errors != nil {
		return Todo{}, errors
	}
	return Todo{title: request.Title(), description: request.Description(), creationDate: time.Now(), dueDate: time.Unix(request.DueDate(), 0)}, nil
}

func validateCreationRequest(request CreationRequest) []DomainError {
	var errs []DomainError

	titleErr := validators.ValidateTitle(request.Title())
	if titleErr != (validators.ValidationError{}) {
		errs = append(errs, NewTodoDomainError(TitleField, titleErr.Code(), titleErr.Description()))
	}
	descriptionErr := validators.ValidateDescription(request.Description())
	if descriptionErr != (validators.ValidationError{}) {
		errs = append(errs, NewTodoDomainError(DescriptionField, descriptionErr.Code(), descriptionErr.Description()))
	}
	dueDateErr := validators.ValidateDueDate(request.DueDate())
	if dueDateErr != (validators.ValidationError{}) {
		errs = append(errs, NewTodoDomainError(DueDateField, dueDateErr.Code(), dueDateErr.Description()))
	}
	return errs
}

func FromTodoUpdateRequest(request UpdateRequest) (Todo, []DomainError) {
	if errors := validateUpdateRequest(request); errors != nil {
		return Todo{}, errors
	}
	return Todo{id: request.Id(), title: request.Title(), description: request.Description(), creationDate: time.Now(), dueDate: time.Unix(request.DueDate(), 0)}, nil
}

func validateUpdateRequest(request UpdateRequest) []DomainError {
	var errs []DomainError

	idErr := validators.ValidateTodoId(request.Id())
	if idErr != (validators.ValidationError{}) {
		errs = append(errs, NewTodoDomainError(IDField, idErr.Code(), idErr.Description()))
	}
	titleErr := validators.ValidateTitle(request.Title())
	if titleErr != (validators.ValidationError{}) {
		errs = append(errs, NewTodoDomainError(TitleField, titleErr.Code(), titleErr.Description()))
	}
	descriptionErr := validators.ValidateDescription(request.Description())
	if descriptionErr != (validators.ValidationError{}) {
		errs = append(errs, NewTodoDomainError(DescriptionField, descriptionErr.Code(), descriptionErr.Description()))
	}
	dueDateErr := validators.ValidateDueDate(request.DueDate())
	if dueDateErr != (validators.ValidationError{}) {
		errs = append(errs, NewTodoDomainError(DueDateField, dueDateErr.Code(), dueDateErr.Description()))
	}
	return errs
}
