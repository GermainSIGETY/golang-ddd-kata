package api

import (
	"time"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/validators"
)

func UpdateFromTodoUpdateRequest(t *model.Todo, request port.UpdateRequest) {
	// Choix model.Todo -> champ externe
	t.Title = request.Title()
	t.Description = request.Description()
	t.DueDate = time.Unix(request.DueDate(), 0)
}

func FromTodoCreationRequest(request port.CreationRequest) (model.Todo, []model.DomainError) {
	if errors := validateCreationRequest(request); len(errors) > 0 {
		return model.Todo{}, errors
	}
	return model.Todo{
		Title:        request.Title(),
		Description:  request.Description(),
		CreationDate: time.Now(),
		DueDate:      time.Unix(request.DueDate(), 0),
		Assignee:     request.Assignee(),
	}, nil
}

func validateCreationRequest(request port.CreationRequest) []model.DomainError {
	var errs []model.DomainError

	titleErr := validators.ValidateTitle(request.Title())
	if titleErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewDomainError(model.TitleField, titleErr.Code(), titleErr.Description()))
	}
	descriptionErr := validators.ValidateDescription(request.Description())
	if descriptionErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewDomainError(model.DescriptionField, descriptionErr.Code(), descriptionErr.Description()))
	}
	dueDateErr := validators.ValidateDueDate(request.DueDate())
	if dueDateErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewDomainError(model.DueDateField, dueDateErr.Code(), dueDateErr.Description()))
	}
	return errs
}

func ValidateUpdateRequest(request port.UpdateRequest) []model.DomainError {
	var errs []model.DomainError

	idErr := validators.ValidateTodoId(request.Id())
	if idErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewDomainError(model.IDField, idErr.Code(), idErr.Description()))
	}
	titleErr := validators.ValidateTitle(request.Title())
	if titleErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewDomainError(model.TitleField, titleErr.Code(), titleErr.Description()))
	}
	descriptionErr := validators.ValidateDescription(request.Description())
	if descriptionErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewDomainError(model.DescriptionField, descriptionErr.Code(), descriptionErr.Description()))
	}
	dueDateErr := validators.ValidateDueDate(request.DueDate())
	if dueDateErr != (validators.ValidationError{}) {
		errs = append(errs, model.NewDomainError(model.DueDateField, dueDateErr.Code(), dueDateErr.Description()))
	}
	return errs
}
