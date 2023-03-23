package api

import (
	"errors"
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/validators"
)

const (
	NotFoundErrorMessage = "Not found"
)

type TodosAPI struct {
	todosRepository port.ITodosRepository
}

func NewApi(repository port.ITodosRepository) TodosAPI {
	return TodosAPI{repository}
}

func (api TodosAPI) CreateTodo(request port.CreationRequest) (int, error) {
	toPersist, errs := FromTodoCreationRequest(request)
	if len(errs) > 0 {
		return 0, model.JoinDomainErrors(errs)
	}
	createdId, creationError := api.todosRepository.Create(toPersist)
	if creationError != nil {
		return 0, creationError
	}
	return createdId, nil
}

func (api TodosAPI) ReadTodo(ID int) (model.ReadTodoResponse, error) {
	todo, err := api.todosRepository.ReadTodo(ID)
	if err != nil {
		fmt.Printf("ERROR : %v", err)
		return model.ReadTodoResponse{}, err
	} else if todo == (model.Todo{}) {
		return model.ReadTodoResponse{}, errors.New(NotFoundErrorMessage)
	}
	return todo.MapToTodoResponse(), nil
}

func (api TodosAPI) UpdateTodo(request port.UpdateRequest) error {
	if errs := ValidateUpdateRequest(request); len(errs) > 0 {
		return model.JoinDomainErrors(errs)
	}

	todo, errRead := api.todosRepository.ReadTodo(request.Id())
	if errRead != nil {
		fmt.Printf("ERROR : %v", errRead)
		return errRead
	}

	if todo == (model.Todo{}) {
		domainError := model.NewDomainError(model.IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return domainError
	}

	UpdateFromTodoUpdateRequest(&todo, request)
	if errUpdate := api.todosRepository.UpdateTodo(todo); errUpdate != nil {
		fmt.Printf("ERROR : %v", errRead)
		return errUpdate
	}
	return nil
}

func (api TodosAPI) DeleteTodo(id int) error {
	deleteError := api.todosRepository.DeleteTodo(id)
	if deleteError != nil {
		return handleDeleteError(deleteError)
	}
	return nil
}

func (api TodosAPI) ReadTodoList() ([]model.ISummaryResponse, error) {
	return api.todosRepository.ReadTodoList()
}

func handleDeleteError(err error) error {
	if err.Error() == port.NoRowDeleted {
		domainError := model.NewDomainError(model.IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return domainError
	} else {
		fmt.Printf("ERROR : %v", err)
		return err
	}
}
