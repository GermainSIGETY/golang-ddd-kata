package model

import "strings"

type Operation string

type TodoErrorPayload struct {
	task []interface{}
}

type Severity int

const (
	FATAL Severity = iota
	WARNING
	INFO
)

type Code struct {
}

type Error struct {
	Op         []Operation // Nom
	RootCauses []error     // Erreur source
	Severity   Severity    // Erreur, Warning, Info

	Context TodoErrorPayload
	Code    Code //
}

//TODO: check if there is a better way to print list of errors with logrus
func (e Error) Error() string {
	var errorMessages []string
	for _, cause := range e.RootCauses {
		errorMessages = append(errorMessages, cause.Error())
	}
	return strings.Join(errorMessages, ",")
}

func (e *Error) AppendOperation(op Operation) *Error {
	e.Op = append(e.Op, op)
	return e
}

func New(args ...interface{}) *Error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Operation:
			e.Op = append(e.Op, arg)
		case DomainError:
			e.RootCauses = append(e.RootCauses, &arg)
		case error:
			e.RootCauses = append(e.RootCauses, arg)
		case Severity:
			e.Severity = arg
		case Code:
			e.Code = arg
		case TodoErrorPayload:
			e.Context = arg
		default:
			panic("bad call to New")
		}
	}

	return e
}
