package validators

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDescriptionCorrect(t *testing.T) {
	s := "Eat a sandwich"
	err := ValidateDescription(&s)
	assert.Empty(t, err)
}

func TestDescriptionTooLong(t *testing.T) {
	tooLong := fmt.Sprintf("%256v", "foo")
	err := ValidateDescription(&tooLong)
	assert.Equal(t, FieldTooLongCode, err.Code())
	assert.Equal(t, fmt.Sprintf(FieldToLongDescription, descriptionMaxSize), err.Description())
}
