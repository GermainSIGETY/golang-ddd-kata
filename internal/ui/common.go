package ui

import (
	"errors"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

const contentType = "Content-Type"
const JSONContentType = "application/json"

type ErrorsArrayJsonResponse struct {
	ErrorsJson []ErrorJsonResponse `json:"errors"`
}
type ErrorJsonResponse struct {
	Code    *string `json:"code,omitempty"`
	Field   *string `json:"field,omitempty"`
	Message *string `json:"message,omitempty"`
}

func answerError(context *gin.Context, err error) {
	if errors.As(err, &model.DomainError{}) {
		answerUnprocessableEntity(context, err)
		return
	}
	if strings.Contains(err.Error(), api.NotFoundErrorMessage) {
		answerError404(context)
		return
	}
	answerError500(context)
}

func answerUnprocessableEntity(context *gin.Context, err error) {
	errorsJsonResponse := fromDomainErrorsToErrorsArrayJsonResponse(err)
	if len(errorsJsonResponse.ErrorsJson) == 0 {
		log.Error().Interface("level", "ui").
			Msg("error has been previously checked as DomainError but cannot be casted as one DomainError nor []error, WTF")
		answerError500(context)
		return
	}
	context.JSON(http.StatusUnprocessableEntity, errorsJsonResponse)
}
func fromDomainErrorsToErrorsArrayJsonResponse(err error) ErrorsArrayJsonResponse {
	var errorsJson ErrorsArrayJsonResponse
	switch errorType := err.(type) {
	case model.DomainError:
		errorsJson = buildDomainErrorsResponseBody([]model.DomainError{errorType})
	case interface{ Unwrap() []error }:
		domainErrors := castToDomainErrors(errorType.Unwrap())
		errorsJson = buildDomainErrorsResponseBody(domainErrors)
	}
	return errorsJson
}

func castToDomainErrors(errors []error) []model.DomainError {
	var domainErrors []model.DomainError
	for _, err := range errors {
		if domainError, is := err.(model.DomainError); is {
			domainErrors = append(domainErrors, domainError)
		} else {
			log.Error().Str("level", "ui").
				Err(err).
				Msg("Cast to a DomainError failed for this error in joined errors")
		}
	}
	return domainErrors
}

func buildDomainErrorsResponseBody(errs []model.DomainError) ErrorsArrayJsonResponse {
	jsonErrors := make([]ErrorJsonResponse, len(errs))
	for i, v := range errs {
		code := v.Code()
		field := v.Field()
		description := v.Description()
		jsonErrors[i] = ErrorJsonResponse{Code: &code, Field: &field, Message: &description}
	}
	response := ErrorsArrayJsonResponse{jsonErrors}
	return response
}

func answerError404(context *gin.Context) {
	notFound := "NOT_FOUND"
	errorJson := ErrorJsonResponse{Code: &notFound}
	context.JSON(http.StatusNotFound, ErrorsArrayJsonResponse{[]ErrorJsonResponse{errorJson}})
}

func answerError400(context *gin.Context, err error) {
	badRequestCode := "BAD_REQUEST"
	message := err.Error()
	jsonErrors := []ErrorJsonResponse{{Code: &badRequestCode, Message: &message}}
	context.JSON(http.StatusBadRequest, ErrorsArrayJsonResponse{jsonErrors})
}

func answerError500(context *gin.Context) {
	internalErrorCode := "INTERNAL_ERROR"
	jsonErrors := []ErrorJsonResponse{{Code: &internalErrorCode}}
	context.JSON(http.StatusInternalServerError, ErrorsArrayJsonResponse{jsonErrors})
}
