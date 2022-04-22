package model

import "fmt"

const (
	IDField          = "id"
	TitleField       = "title"
	DescriptionField = "description"
	DueDateField     = "dueDate"
)

// Transformer en struct immutable ou champ exportable ?
type DomainError struct {
	field       string
	code        string
	description string
}

func NewTodoDomainError(field string, code string, description string) DomainError {
	return DomainError{field: field, code: code, description: description}
}

func (e DomainError) Field() string {
	return e.field
}

func (e DomainError) Code() string {
	return e.code
}

func (e DomainError) Description() string {
	return e.description
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("%v %v %v", e.field, e.code, e.description)
}
