package api

import (
	"errors"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/validators"
	"github.com/rs/zerolog/log"
)

const (
	NotFoundErrorMessage = "Not found"
)

type TodosAPI struct {
	todosRepository     port.ITodosRepository
	notificationChannel chan<- int
}

func NewApi(repository port.ITodosRepository, notificationChannel chan<- int) TodosAPI {
	return TodosAPI{repository, notificationChannel}
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
	// send id to notification channel : a notification will be (eventually) sent
	api.notificationChannel <- createdId
	return createdId, nil
}

func (api TodosAPI) ReadTodo(ID int) (model.ReadTodoResponse, error) {
	todo, err := api.todosRepository.ReadTodo(ID)
	if err != nil {
		log.Err(err).Msg("")
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
		log.Err(errRead).Msg("")
		return errRead
	}

	if todo == (model.Todo{}) {
		domainError := model.NewDomainError(model.IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return domainError
	}

	UpdateFromTodoUpdateRequest(&todo, request)
	if errUpdate := api.todosRepository.UpdateTodo(todo); errUpdate != nil {
		log.Err(errUpdate).Msg("")
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
		log.Err(err).Msg("")
		return err
	}
}
