package api

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/validators"
)

type TodosAPI struct {
	todosRepository port.ITodosRepository
}

func NewApi(repository port.ITodosRepository) TodosAPI {
	return TodosAPI{repository}
}

func (api TodosAPI) CreateTodo(request port.CreationRequest) (int, interface{}) {
	toPersist, err := FromTodoCreationRequest(request)
	if err != nil {
		return 0, err
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
		return model.ReadTodoResponse{}, err
	}
	return todo.
		MapToTodoResponse(), nil
}

func (api TodosAPI) UpdateTodo(request port.UpdateRequest) interface{} {
	todo, errRead := api.todosRepository.ReadTodo(request.Id())
	if errRead != nil {
		return errRead
	}

	if todo == (model.Todo{}) {
		domainError := model.NewTodoDomainError(model.IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return []model.DomainError{domainError}
	}

	UpdateFromTodoUpdateRequest(&todo, request)
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

func (api TodosAPI) ReadTodoList() ([]model.ISummaryResponse, error) {
	return api.todosRepository.ReadTodoList()
}

func handleDeleteError(err error) interface{} {
	if err.Error() == port.NoRowDeleted {
		domainError := model.NewTodoDomainError(model.IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return []model.DomainError{domainError}

	}
	return err
}
