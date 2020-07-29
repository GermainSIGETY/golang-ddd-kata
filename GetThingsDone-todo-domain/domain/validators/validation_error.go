package validators

import "fmt"

const (
	EmptyFieldCode           = "EMPTY_FIELD"
	FieldTooLongCode         = "FIELD_TOO_LONG"
	InvalidNumberCode        = "INVALID_NUMBER"
	InvalidTodoIdCode        = "INVALID_TODO_ID"
	EmptyFieldDescription    = "please fill this field"
	FieldToLongDescription   = "this field must be less than %v caracters"
	InvalidNumberDescription = "invalid number format"
	InvalidTodoIdDescription = "no existing todo with this id"
)

type ValidationError struct {
	code        string
	description string
}

func (v ValidationError) Error() string {
	return v.code + " : " + v.description
}

func EmptyField() ValidationError {
	return ValidationError{EmptyFieldCode, EmptyFieldDescription}
}

func FieldTooLong(maxSize int) ValidationError {
	return ValidationError{FieldTooLongCode, fmt.Sprintf(FieldToLongDescription, maxSize)}
}

func InvalidNumber() ValidationError {
	return ValidationError{InvalidNumberCode, InvalidNumberDescription}
}

func (v ValidationError) Code() string {
	return v.code
}

func (v ValidationError) Description() string {
	return v.description
}
