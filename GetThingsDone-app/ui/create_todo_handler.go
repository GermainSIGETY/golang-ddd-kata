package ui

import (
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/api"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoCreationJSONRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	DueDate     *int64  `json:"dueDate"`
}

type TodoCreationJSONResponse struct {
	ID int `json:"id"`
}

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
