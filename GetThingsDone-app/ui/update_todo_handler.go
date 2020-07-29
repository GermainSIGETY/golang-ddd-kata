package ui

import (
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/api"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TodoUpdateJSONRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	DueDate     *int64  `json:"dueDate"`
}

func HandleUpdate(context *gin.Context, IDAsString string, api api.TodosAPI) {
	ID, err := strconv.Atoi(IDAsString)
	if err != nil {
		answerBadRequest(context, "todo ID in path must be an integer")
		return
	}

	var jsonRequest TodoUpdateJSONRequest
	if errs := context.ShouldBindJSON(&jsonRequest); errs != nil {
		fmt.Print(errs)
		answerBadRequest(context, "unable to parse TODO update JSON body")
		return
	}

	request, errs := presentation.NewTodoUpdateRequest(ID, jsonRequest.Title, jsonRequest.Description, jsonRequest.DueDate)
	if errs != nil {
		answerError(context, errs)
		return
	}

	updateErrors := api.UpdateTodo(request)
	if updateErrors != nil {
		answerError(context, updateErrors)
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
