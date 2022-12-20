package api

import (
	"github.com/GermainSIGETY/golang-ddd-kata/pkg/domain_error"
	"time"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/validators"
)

func UpdateFromTodoUpdateRequest(t *model.Todo, request port.UpdateRequest) {
	// Choix model.Todo -> champ externe
	t.Title = request.Title()
	t.Description = request.Description()
	t.DueDate = time.Unix(request.DueDate(), 0)
}

func FromTodoCreationRequest(request port.CreationRequest) (model.Todo, []domain_error.DomainError) {
	if errors := validateCreationRequest(request); len(errors) > 0 {
		return model.Todo{}, errors
	}
	return model.Todo{
		Title:        request.Title(),
		Description:  request.Description(),
		CreationDate: time.Now(),
		DueDate:      time.Unix(request.DueDate(), 0),
	}, nil
}

func validateCreationRequest(request port.CreationRequest) []domain_error.DomainError {
	var errs []domain_error.DomainError

	titleErr := validators.ValidateTitle(request.Title())
	if titleErr != (validators.ValidationError{}) {
		errs = append(errs, domain_error.NewDomainError(model.TitleField, titleErr.Code(), titleErr.Description()))
	}
	descriptionErr := validators.ValidateDescription(request.Description())
	if descriptionErr != (validators.ValidationError{}) {
		errs = append(errs, domain_error.NewDomainError(model.DescriptionField, descriptionErr.Code(), descriptionErr.Description()))
	}
	dueDateErr := validators.ValidateDueDate(request.DueDate())
	if dueDateErr != (validators.ValidationError{}) {
		errs = append(errs, domain_error.NewDomainError(model.DueDateField, dueDateErr.Code(), dueDateErr.Description()))
	}
	return errs
}

func ValidateUpdateRequest(request port.UpdateRequest) []domain_error.DomainError {
	var errs []domain_error.DomainError

	idErr := validators.ValidateTodoId(request.Id())
	if idErr != (validators.ValidationError{}) {
		errs = append(errs, domain_error.NewDomainError(model.IDField, idErr.Code(), idErr.Description()))
	}
	titleErr := validators.ValidateTitle(request.Title())
	if titleErr != (validators.ValidationError{}) {
		errs = append(errs, domain_error.NewDomainError(model.TitleField, titleErr.Code(), titleErr.Description()))
	}
	descriptionErr := validators.ValidateDescription(request.Description())
	if descriptionErr != (validators.ValidationError{}) {
		errs = append(errs, domain_error.NewDomainError(model.DescriptionField, descriptionErr.Code(), descriptionErr.Description()))
	}
	dueDateErr := validators.ValidateDueDate(request.DueDate())
	if dueDateErr != (validators.ValidationError{}) {
		errs = append(errs, domain_error.NewDomainError(model.DueDateField, dueDateErr.Code(), dueDateErr.Description()))
	}
	return errs
}
