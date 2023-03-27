package ui

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/gin-gonic/gin"
)

// We do not use validation tags like : validate:"required"
// Because :
// - it only works with pointers,
// - and we centralize requests validation in domain : validations are business rules !
type todoUpdateJSONRequest struct {
	id              int
	TitleJson       string `json:"title"`
	DescriptionJson string `json:"description"`
	DueDateJson     int64  `json:"dueDate"`
	AssigneeJson    string `json:"assignee"`
}

func (t todoUpdateJSONRequest) Id() int {
	return t.id
}

func (t todoUpdateJSONRequest) Title() string {
	return t.TitleJson
}

func (t todoUpdateJSONRequest) Description() string {
	return t.DescriptionJson
}

func (t todoUpdateJSONRequest) DueDate() int64 {
	return t.DueDateJson
}

func (t todoUpdateJSONRequest) Assignee() string {
	return t.AssigneeJson
}

// HandleUpdate godoc
// @Summary Update a todo
// @Description Update a todo by its Id
// @id update-todo
// @Tags todos
// @Accept json
// @Param id path int true "Todo ID"
// @Param body body todoUpdateJSONRequest true "todo infos"
// @Success 204 "it's ok"
// @Failure 422 {object} ErrorsArrayJsonResponse
// @Failure 500 {object} ErrorsArrayJsonResponse
// @Router /todos/{id} [put]
func handleUpdate(context *gin.Context, IDAsString string, api api.TodosAPI) {
	id, err := strconv.Atoi(IDAsString)
	if err != nil {
		answerError400(context, err)
		return
	}

	var jsonRequest todoUpdateJSONRequest
	if err := context.ShouldBindJSON(&jsonRequest); err != nil {
		log.Err(err).Msg("")
		answerError400(context, err)
		return
	}

	jsonRequest.id = id
	errs := api.UpdateTodo(jsonRequest)
	if errs != nil {
		answerError(context, errs)
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
