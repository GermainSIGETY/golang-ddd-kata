package todo

import (
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/src/api"
	"github.com/GermainSIGETY/golang-ddd-kata/src/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoCreationJSONRequest struct {
	Title       *string `json:"title" validate:"required"`
	Description *string `json:"description"`
	DueDate     *int64  `json:"dueDate" validate:"required"`
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
// @Param body body TodoCreationJSONRequest true "todo infos"
// @Success 200 {object} TodoCreationJSONResponse
// @Failure 422 {object} ErrorsArrayJsonResponse
// @Failure 500 {object} ErrorsArrayJsonResponse
// @Router /todos [post]
func HandleCreate(context *gin.Context, api api.TodosAPI) {
	var jsonRequest TodoCreationJSONRequest
	if errs := context.ShouldBindJSON(&jsonRequest); errs != nil {
		fmt.Print(errs)
		answerBadRequest(context, "unable to parse TODO Creation JSON body")
		return
	}

	request, errs := presentation.NewTodoCreationRequest(jsonRequest.Title, jsonRequest.Description, jsonRequest.DueDate)
	if errs != nil {
		answerError(context, errs)
		return
	}

	ID, err := api.CreateTodo(request)
	if err != nil {
		answerError(context, err)
		return
	}
	context.JSON(http.StatusOK, TodoCreationJSONResponse{ID})
}
