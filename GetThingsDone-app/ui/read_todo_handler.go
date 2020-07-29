package ui

import (
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/api"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TodoReadJSONResponse struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Description  *string `json:"description,omitempty"`
	CreationDate int64   `json:"creationDate"`
	DueDate      int64   `json:"dueDate"`
}

func HandleReadTodo(context *gin.Context, IDAsString string, api api.TodosAPI) {
	ID, err := strconv.Atoi(IDAsString)
	if err != nil {
		answerBadRequest(context, "todo ID in path must be an integer")
		return
	}

	todo, errs := api.ReadTodo(ID)
	if errs != nil {
		answerError(context, errs)
		return
	}
	if todo == (presentation.ReadTodoResponse{}) {
		answerResourceNotFound(context, fmt.Sprintf("no todo with ID : %v", ID))
		return
	}

	response := TodoReadJSONResponse{todo.ID, todo.Title, todo.Description, todo.CreationDate.Unix(), todo.DueDate.Unix()}
	context.JSON(http.StatusOK, response)
}
