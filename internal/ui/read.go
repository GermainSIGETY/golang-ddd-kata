package ui

import (
	"net/http"
	"strconv"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/gin-gonic/gin"
)

type TodoReadJSONResponse struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	CreationDate int64  `json:"creationDate"`
	DueDate      int64  `json:"dueDate"`
	Assignee     string `json:"assignee,omitempty"`
}

// HandleReadTodo godoc
// @Summary Read a todo
// @Description Read a todo by its Id
// @id read-todo
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} TodoReadJSONResponse
// @Failure 404 {object} ErrorsArrayJsonResponse
// @Failure 500 {object} ErrorsArrayJsonResponse
// @Router /todos/{id} [get]
func handleReadTodo(context *gin.Context, IDAsString string, api api.TodosAPI) {
	ID, err := strconv.Atoi(IDAsString)
	if err != nil {
		answerError400(context, err)
		return
	}
	response, error := api.ReadTodo(ID)
	if error != nil {
		answerError(context, error)
		return
	}
	jsonResponse := TodoReadJSONResponse{response.ID, response.Title, response.Description,
		response.CreationDate.Unix(), response.DueDate.Unix(), response.Assignee}
	context.JSON(http.StatusOK, jsonResponse)
}
