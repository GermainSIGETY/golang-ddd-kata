package ui

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
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
func handleCreate(context *gin.Context) {
	var jsonRequest todoCreationJSONRequest
	if errs := context.ShouldBindJSON(&jsonRequest); errs != nil {
		fmt.Print(errs)
		answerBadRequest(context, "unable to parse TODO Creation JSON body")
		return
	}

	ID, err := api.GetTodoApi().CreateTodo(jsonRequest)
	if err != nil {
		answerError(context, err)
		return
	}
	context.JSON(http.StatusOK, TodoCreationJSONResponse{ID})
}
