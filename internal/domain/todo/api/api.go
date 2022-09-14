package api

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/validators"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/infrastructure"
)

type TodosAPI struct{}

var todoAPI *TodosAPI

func init() {
	todoAPI = new(TodosAPI)
}

func GetTodoApi() TodosAPI {
	return *todoAPI
}

func (api TodosAPI) CreateTodo(request port.CreationRequest) (int, interface{}) {
	toPersist, err := FromTodoCreationRequest(request)
	if err != nil {
		return 0, err
	}
	createdId, creationError := infrastructure.GetTodosRepository().Create(toPersist)
	if creationError != nil {
		return 0, creationError
	}
	return createdId, nil
}

func (api TodosAPI) ReadTodo(ID int) (model.ReadTodoResponse, error) {
	todo, err := infrastructure.GetTodosRepository().ReadTodo(ID)
	if err != nil {
		return model.ReadTodoResponse{}, err
	}
	return todo.
		MapToTodoResponse(), nil
}

func (api TodosAPI) UpdateTodo(request port.UpdateRequest) interface{} {
	todo, errRead := infrastructure.GetTodosRepository().ReadTodo(request.Id())
	if errRead != nil {
		return errRead
	}

	if todo == (model.Todo{}) {
		domainError := model.NewTodoDomainError(model.IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return []model.DomainError{domainError}
	}

	UpdateFromTodoUpdateRequest(&todo, request)
	if errUpdate := infrastructure.GetTodosRepository().UpdateTodo(todo); errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (api TodosAPI) DeleteTodo(id int) interface{} {
	deleteError := infrastructure.GetTodosRepository().DeleteTodo(id)
	if deleteError != nil {
		return handleDeleteError(deleteError)
	}
	return nil
}

func (api TodosAPI) ReadTodoList() ([]model.ISummaryResponse, error) {
	return infrastructure.GetTodosRepository().ReadTodoList()
}

func handleDeleteError(err error) interface{} {
	if err.Error() == port.NoRowDeleted {
		domainError := model.NewTodoDomainError(model.IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return []model.DomainError{domainError}

	}
	return err
}
