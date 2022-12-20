package ui

import (
	"errors"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/GermainSIGETY/golang-ddd-kata/pkg/domain_error"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

const contentType = "Content-Type"
const JSONContentType = "application/json"

type ErrorsArrayJsonResponse struct {
	ErrorJsonResponse []ErrorJsonResponse `json:"errors"`
}
type ErrorJsonResponse struct {
	Code    *string `json:"code,omitempty"`
	Field   *string `json:"field,omitempty"`
	Message *string `json:"message,omitempty"`
}

func answerError(context *gin.Context, err error) {
	if errors.As(err, &domain_error.DomainError{}) {
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
	var errorsJson ErrorsArrayJsonResponse

	switch errorType := err.(type) {
	case domain_error.DomainError:
		errorsJson = buildDomainErrorsResponseBody([]domain_error.DomainError{errorType})
	case interface{ Unwrap() []error }:
		domainErrors := castToDomainErrors(errorType.Unwrap())
		errorsJson = buildDomainErrorsResponseBody(domainErrors)
	default:
		{
			log.Error().Interface("level", "ui").
				Msg("error has been previously checked as DomainError but cannot be casted as one DomainError nor []error, WTF")
			answerError500(context)
			return
		}
	}
	context.JSON(http.StatusUnprocessableEntity, errorsJson)
}

func castToDomainErrors(errors []error) []domain_error.DomainError {
	var domainErrors []domain_error.DomainError
	for _, err := range errors {
		if domainError, is := err.(domain_error.DomainError); is {
			domainErrors = append(domainErrors, domainError)
		} else {
			log.Error().Str("level", "ui").
				Err(err).
				Msg("Cast to a DomainError failed for this error in joined errors")
		}
	}
	return domainErrors
}

func buildDomainErrorsResponseBody(errs []domain_error.DomainError) ErrorsArrayJsonResponse {
	jsonErrors := make([]ErrorJsonResponse, len(errs))
	for i, v := range errs {
		field := v.Field()
		description := v.Description()
		jsonErrors[i] = ErrorJsonResponse{Code: &field, Message: &description}
	}
	response := ErrorsArrayJsonResponse{jsonErrors}
	return response
}

func answerError404(context *gin.Context) {
	notFound := "NOT_FOUND"
	errorJson := ErrorJsonResponse{Code: &notFound}
	context.JSON(http.StatusNotFound, ErrorsArrayJsonResponse{[]ErrorJsonResponse{errorJson}})
}

func answerError400(context *gin.Context, message string) {
	context.Header(contentType, JSONContentType)
	context.String(http.StatusBadRequest, "{\"errors\":[{\"message\": \"%v\"}]}", message)
}

func answerError500(context *gin.Context) {
	context.Header(contentType, JSONContentType)
	context.String(http.StatusInternalServerError, "{\"errors\":[{\"code\": \"INTERNAL_ERROR\"}]}")
}
