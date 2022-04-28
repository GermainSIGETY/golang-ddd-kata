package ui

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
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
func handleUpdate(context *gin.Context, IDAsString string) {
	id, err := strconv.Atoi(IDAsString)
	if err != nil {
		answerBadRequest(context, "todo ID in path must be an integer")
		return
	}

	var jsonRequest todoUpdateJSONRequest
	if errs := context.ShouldBindJSON(&jsonRequest); errs != nil {
		fmt.Print(errs)
		answerBadRequest(context, "unable to parse TODO update JSON body")
		return
	}

	jsonRequest.id = id
	errs := api.GetTodoApi().UpdateTodo(jsonRequest)
	if errs != nil {
		answerError(context, errs)
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
