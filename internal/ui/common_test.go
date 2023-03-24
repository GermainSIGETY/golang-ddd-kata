package ui

import (
	"errors"
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fromOneDomainErrorToErrorsArrayJsonResponse(t *testing.T) {
	trickyfield := "trickyfield"
	damnItCode := "DAMN_IT"
	verySorryDescription := "I am very very sorry"
	domainError := model.NewDomainError(trickyfield, damnItCode, verySorryDescription)

	errorsJson := fromDomainErrorsToErrorsArrayJsonResponse(domainError)
	assert.Equal(t, 1, len(errorsJson.ErrorsJson))
	assert.Equal(t, trickyfield, *errorsJson.ErrorsJson[0].Field)
	assert.Equal(t, damnItCode, *errorsJson.ErrorsJson[0].Code)
	assert.Equal(t, verySorryDescription, *errorsJson.ErrorsJson[0].Message)
}

func Test_fromTwoDomainErrorsToErrorsArrayJsonResponse(t *testing.T) {
	trickyfield := "trickyfield"
	repulsiveField := "repulsiveField"
	damnItCode := "DAMN_IT"
	verySorryDescription := "I am very very sorry"
	domainError1 := model.NewDomainError(trickyfield, damnItCode, verySorryDescription)
	domainError2 := model.NewDomainError(repulsiveField, damnItCode, verySorryDescription)

	errorsJson := fromDomainErrorsToErrorsArrayJsonResponse(errors.Join(domainError1, domainError2))
	assert.Equal(t, 2, len(errorsJson.ErrorsJson))
	assert.Equal(t, trickyfield, *errorsJson.ErrorsJson[0].Field)
	assert.Equal(t, damnItCode, *errorsJson.ErrorsJson[0].Code)
	assert.Equal(t, verySorryDescription, *errorsJson.ErrorsJson[0].Message)
	assert.Equal(t, repulsiveField, *errorsJson.ErrorsJson[1].Field)

}

func Test_fromBasicErrorToNothing(t *testing.T) {
	basicError := fmt.Errorf("hello I'm a basic error")
	errorsJson := fromDomainErrorsToErrorsArrayJsonResponse(basicError)
	assert.Equal(t, 0, len(errorsJson.ErrorsJson))
}
