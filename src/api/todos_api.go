package api

import (
	. "github.com/GermainSIGETY/golang-ddd-kata/src/domain"
	"github.com/GermainSIGETY/golang-ddd-kata/src/domain/validators"
	"github.com/GermainSIGETY/golang-ddd-kata/src/infrastructure"
	"github.com/GermainSIGETY/golang-ddd-kata/src/presentation"
)

type TodosAPI struct {
	todosRepository infrastructure.ITodosRepository
}

func New(repository infrastructure.ITodosRepository) TodosAPI {
	return TodosAPI{repository}
}

func (api TodosAPI) CreateTodo(request presentation.TodoCreationRequest) (int, error) {

	toPersist := FromTodoCreationRequest(request)
	createdId, err := api.todosRepository.Create(toPersist)
	if err == nil {
		return createdId, nil
	}
	return 0, err
}

func (api TodosAPI) ReadTodo(ID int) (presentation.ReadTodoResponse, error) {
	todo, err := api.todosRepository.ReadTodo(ID)
	if err != nil {
		return presentation.ReadTodoResponse{}, err
	}
	return todo.MapToTodoResponse(), nil
}

func (api TodosAPI) UpdateTodo(request presentation.TodoUpdateRequest) interface{} {
	todo, errRead := api.todosRepository.ReadTodo(request.Id())
	if errRead != nil {
		return errRead
	}

	if todo == (Todo{}) {
		domainError := presentation.NewTodoDomainError(presentation.IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return []presentation.TodoDomainError{domainError}
	}

	todo.UpdateFromTodoUpdateRequest(request)
	if errUpdate := api.todosRepository.UpdateTodo(todo); errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (api TodosAPI) DeleteTodo(id int) interface{} {
	deleteError := api.todosRepository.DeleteTodo(id)
	if deleteError != nil {
		return handleDeleteError(deleteError)
	}
	return nil
}

func (api TodosAPI) ReadTodoList() ([]presentation.TodoSummaryResponse, error) {
	return api.todosRepository.ReadTodoList()
}

func handleDeleteError(err error) interface{} {
	if err.Error() == infrastructure.NoRowDeleted {
		domainError := presentation.NewTodoDomainError(presentation.IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return []presentation.TodoDomainError{domainError}

	}
	return err
}
