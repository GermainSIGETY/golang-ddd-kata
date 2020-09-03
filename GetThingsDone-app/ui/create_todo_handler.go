package ui

import (
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/api"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/presentation"
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

// @Summary Create a todo
// @Description Create a todo
// @Accept  json
// @Produce  json
// @Param todo body TodoCreationJSONRequest true "todo fields"
// @Success 200 {object} TodoCreationJSONResponse
// @Failure 500 {object} ErrorsArrayJsonResponse
// @Failure 422 {object} ErrorsArrayJsonResponse
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
