package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
)

type TodoListJSONResponse struct {
	TodoList []TodoSummaryJSONResponse `json:"todos"`
}

type TodoSummaryJSONResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	DueDate int64  `json:"dueDate"`
}

// HandleReadTodoList godoc
// @Summary Read todos
// @Description Read all todos
// @id read-todos
// @Tags todos
// @Produce json
// @Success 200 {object} TodoListJSONResponse
// @Failure 500 {object} ErrorsArrayJsonResponse
// @Router /todos [get]
func HandleReadTodoList(context *gin.Context) {
	todos, errs := api.GetTodoApi().ReadTodoList()
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
