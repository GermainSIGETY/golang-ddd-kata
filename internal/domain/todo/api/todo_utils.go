package api

import (
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

func FromTodoCreationRequest(request port.CreationRequest) (model.Todo, *model.Error) {
	op := model.Operation("api.convertTodoCreationRequest")
	if err := validateCreationRequest(request); err != nil {
		return model.Todo{}, err.AppendOperation(op)
	}
	return model.Todo{
		Title:        request.Title(),
		Description:  request.Description(),
		CreationDate: time.Now(),
		DueDate:      time.Unix(request.DueDate(), 0),
	}, nil
}

func validateCreationRequest(request port.CreationRequest) *model.Error {
	myError := model.New(model.Operation("api.validateCreationRequest"), model.FATAL)
	titleErr := validators.ValidateTitle(request.Title())
	if titleErr != (validators.ValidationError{}) {
		myError.RootCauses = append(myError.RootCauses, model.NewTodoDomainError(model.TitleField, titleErr.Code(), titleErr.Description()))
	}
	descriptionErr := validators.ValidateDescription(request.Description())
	if descriptionErr != (validators.ValidationError{}) {
		myError.RootCauses = append(myError.RootCauses, model.NewTodoDomainError(model.DescriptionField, descriptionErr.Code(), descriptionErr.Description()))
	}
	dueDateErr := validators.ValidateDueDate(request.DueDate())
	if dueDateErr != (validators.ValidationError{}) {
		myError.RootCauses = append(myError.RootCauses, model.NewTodoDomainError(model.DueDateField, dueDateErr.Code(), dueDateErr.Description()))
	}
	if len(myError.RootCauses) > 0 {
		// TODO: handle payload
		//var payload model.TodoErrorPayload
		return myError
	}

	return nil
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
		errs = append(errs, *model.NewTodoDomainError(model.IDField, idErr.Code(), idErr.Description()))
	}
	titleErr := validators.ValidateTitle(request.Title())
	if titleErr != (validators.ValidationError{}) {
		errs = append(errs, *model.NewTodoDomainError(model.TitleField, titleErr.Code(), titleErr.Description()))
	}
	descriptionErr := validators.ValidateDescription(request.Description())
	if descriptionErr != (validators.ValidationError{}) {
		errs = append(errs, *model.NewTodoDomainError(model.DescriptionField, descriptionErr.Code(), descriptionErr.Description()))
	}
	dueDateErr := validators.ValidateDueDate(request.DueDate())
	if dueDateErr != (validators.ValidationError{}) {
		errs = append(errs, *model.NewTodoDomainError(model.DueDateField, dueDateErr.Code(), dueDateErr.Description()))
	}
	return errs
}
