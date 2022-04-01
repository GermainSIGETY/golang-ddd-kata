package presentation

import "fmt"

type TodoDomainError struct {
	field       string
	code        string
	description string
}

func NewTodoDomainError(field string, code string, description string) TodoDomainError {
	return TodoDomainError{field: field, code: code, description: description}
}

func (e TodoDomainError) Field() string {
	return e.field
}

func (e TodoDomainError) Code() string {
	return e.code
}

func (e TodoDomainError) Description() string {
	return e.description
}

func (e *TodoDomainError) Error() string {
	return fmt.Sprintf("%v %v %v", e.field, e.code, e.description)
}
