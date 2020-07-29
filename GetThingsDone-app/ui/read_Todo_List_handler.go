package ui

import (
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoListJSONResponse struct {
	TodoList []TodoSummaryJSONResponse `json:"todos"`
}

type TodoSummaryJSONResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	DueDate int64  `json:"dueDate"`
}

func HandleReadTodoList(context *gin.Context, api api.TodosAPI) {
	todos, errs := api.ReadTodoList()
	if errs != nil {
		answerError(context, errs)
		return
	}

	var jsonSummaries = make([]TodoSummaryJSONResponse, len(todos))
	for i, todo := range todos {
		jsonSummaries[i] = TodoSummaryJSONResponse{todo.Id(), todo.Title(), todo.DueDate().Unix()}
	}
	context.JSON(http.StatusOK, TodoListJSONResponse{jsonSummaries})
}
