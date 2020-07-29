package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTodoIDCorrect(t *testing.T) {
	var ID int = 1
	err := ValidateTodoId(ID)
	assert.Empty(t, err)
}

func TestTodoIDInvalid(t *testing.T) {
	var ID = 0
	err := ValidateTodoId(ID)
	assert.Equal(t, InvalidNumberCode, err.Code())
	assert.Equal(t, InvalidNumberDescription, err.Description())
}
