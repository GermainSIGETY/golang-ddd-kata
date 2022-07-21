package ui

import (
	"errors"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/gin-gonic/gin"
)

// We do not use validation tags like : validate:"required"
// Because :
// - it only works with pointers,
// - and we centralize requests validation in domain : validations are business rules !
type todoCreationJSONRequest struct {
	TitleJson       string `json:"title"`
	DescriptionJson string `json:"description"`
	DueDateJson     int64  `json:"dueDate"`
}

func (t todoCreationJSONRequest) Title() string {
	return t.TitleJson
}

func (t todoCreationJSONRequest) Description() string {
	return t.DescriptionJson
}

func (t todoCreationJSONRequest) DueDate() int64 {
	return t.DueDateJson
}

type TodoCreationJSONResponse struct {
	ID int `json:"id"`
}

// HandleCreate godoc
// @Summary Create a todo
// @Description Create a new todo
// @id create-todo
// @Tags todos
// @Accept json
// @Produce json
// @Param body body todoCreationJSONRequest true "todo infos"
// @Success 200 {object} TodoCreationJSONResponse
// @Failure 422 {object} ErrorsArrayJsonResponse
// @Failure 500 {object} ErrorsArrayJsonResponse
// @Router /todos [post]
func handleCreate(context *gin.Context, api api.TodosAPI) {
	op := model.Operation("ui.handleCreate")
	var jsonRequest todoCreationJSONRequest
	if errs := context.ShouldBindJSON(&jsonRequest); errs != nil {
		errorMessage := "unable to parse TODO Creation JSON body"
		err := model.New(op, model.FATAL, errors.New(errorMessage))
		logError(err)
		answerBadRequest(context, errorMessage)
		return
	}

	ID, err := api.CreateTodo(jsonRequest)
	if err != nil {
		err.AppendOperation(op)
		logError(err)
		answerError(context, err)
		return
	}
	context.JSON(http.StatusOK, TodoCreationJSONResponse{ID})
}

// Consommation:
func logError(err *model.Error) {
	switch err.Severity {
	case model.FATAL:
		log.WithError(err).WithFields(log.Fields{
			"operations": err.Op,
		}).Error("Fatal error")
	case model.WARNING:
		log.WithError(err).WithFields(log.Fields{
			"operations": err.Op,
		}).Warn("Warning error")
	case model.INFO:
		log.WithError(err).WithFields(log.Fields{
			"operations": err.Op,
		}).Info("Info error")
	default:
		log.WithError(err).WithFields(log.Fields{
			"operations": err.Op,
		}).Error("Fatal error")
	}

}
