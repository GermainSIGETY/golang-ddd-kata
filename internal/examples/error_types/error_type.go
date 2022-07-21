package error_types

import (
	"errors"
	"fmt"
)

// PathError records an error and the operation
// and file path that caused it.
type PathError struct {
	Op   string
	Path string
	Err  error // the cause
}

func (e *PathError) Error() string {
	return fmt.Sprintf("error")
}

func usage() {
	err := something()
	switch err := err.(type) {
	case nil:
		// call succeeded, nothing to do
	case *PathError:
		fmt.Println("error occurred on line:", err.Path)
	default:
		// unknown error
	}
}

func something() error {
	return errors.New("error")
}

// Avantage / inconvénient ?

// ++++ Extra context
// ---- Dépendances entre les packages
