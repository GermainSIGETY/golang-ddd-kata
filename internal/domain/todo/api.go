package todo

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/validators"
)

type TodosAPI struct {
	todosRepository ITodosRepository
}

func NewApi(repository ITodosRepository) TodosAPI {
	return TodosAPI{repository}
}

func (api TodosAPI) CreateTodo(request CreationRequest) (int, interface{}) {
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

func (api TodosAPI) ReadTodo(ID int) (ReadTodoResponse, error) {
	todo, err := api.todosRepository.ReadTodo(ID)
	if err != nil {
		return ReadTodoResponse{}, err
	}
	return todo.
		MapToTodoResponse(), nil
}

func (api TodosAPI) UpdateTodo(request UpdateRequest) interface{} {
	todo, errRead := api.todosRepository.ReadTodo(request.Id())
	if errRead != nil {
		return errRead
	}

	if todo == (Todo{}) {
		domainError := NewTodoDomainError(IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return []DomainError{domainError}
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

func (api TodosAPI) ReadTodoList() ([]SummaryResponse, error) {
	return api.todosRepository.ReadTodoList()
}

func handleDeleteError(err error) interface{} {
	if err.Error() == NoRowDeleted {
		domainError := NewTodoDomainError(IDField, validators.InvalidTodoIdCode, validators.InvalidTodoIdDescription)
		return []DomainError{domainError}

	}
	return err
}
