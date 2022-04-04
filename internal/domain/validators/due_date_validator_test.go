package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDueDateCorrect(t *testing.T) {
	var dueDate int64 = 1
	err := ValidateDueDate(dueDate)
	assert.Empty(t, err)
}

func TestDueDateEmpty(t *testing.T) {
	err := ValidateDueDate(0)
	assert.Equal(t, EmptyFieldCode, err.Code())
	assert.Equal(t, EmptyFieldDescription, err.Description())
}

func TestDueDateInvalid(t *testing.T) {
	var dueDate int64 = -1
	err := ValidateDueDate(dueDate)
	assert.Equal(t, InvalidNumberCode, err.Code())
	assert.Equal(t, InvalidNumberDescription, err.Description())
}
