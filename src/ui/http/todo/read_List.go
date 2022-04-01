package todo

import (
	"github.com/GermainSIGETY/golang-ddd-kata/src/api"
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

// HandleReadTodo godoc
// @Summary Read todos
// @Description Read all todos
// @id read-todos
// @Tags todos
// @Produce json
// @Success 200 {object} TodoListJSONResponse
// @Failure 500 {object} ErrorsArrayJsonResponse
// @Router /todos [get]
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
