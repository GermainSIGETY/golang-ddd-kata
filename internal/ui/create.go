package ui

import (
	"github.com/rs/zerolog/log"
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
	AssigneeJson    string `json:"assignee"`
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

func (t todoCreationJSONRequest) Assignee() string {
	return t.AssigneeJson
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
	var jsonRequest todoCreationJSONRequest
	if err := context.ShouldBindJSON(&jsonRequest); err != nil {
		log.Warn().Err(err).Msg("")
		answerError400(context, err)
		return
	}

	ID, err := api.CreateTodo(jsonRequest)
	if err != nil {
		answerError(context, err)
		return
	}
	context.JSON(http.StatusOK, TodoCreationJSONResponse{ID})
}
