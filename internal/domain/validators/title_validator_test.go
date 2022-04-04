package validators

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTitleCorrect(t *testing.T) {
	title := "Come back from weekend"
	err := ValidateTitle(title)
	assert.Empty(t, err)
}

func TestTitleEmpty(t *testing.T) {
	title := ""
	err := ValidateTitle(title)
	assert.Equal(t, EmptyFieldCode, err.Code())
	assert.Equal(t, EmptyFieldDescription, err.Description())
}

func TestTitleNil(t *testing.T) {
	err := ValidateTitle("")
	assert.Equal(t, EmptyFieldCode, err.Code())
	assert.Equal(t, EmptyFieldDescription, err.Description())
}

func TestTitleTooLong(t *testing.T) {
	titleTooLong := fmt.Sprintf("%256v", "foo")
	err := ValidateTitle(titleTooLong)
	assert.Equal(t, FieldTooLongCode, err.Code())
	assert.Equal(t, fmt.Sprintf(FieldToLongDescription, titleMaxSize), err.Description())
}
