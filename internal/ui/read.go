package ui

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
)

type TodoReadJSONResponse struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	CreationDate int64  `json:"creationDate"`
	DueDate      int64  `json:"dueDate"`
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
func handleReadTodo(context *gin.Context, IDAsString string) {
	ID, err := strconv.Atoi(IDAsString)
	if err != nil {
		answerBadRequest(context, "todo ID in path must be an integer")
		return
	}

	response, errs := api.GetTodoApi().ReadTodo(ID)
	if errs != nil {
		answerError(context, errs)
		return
	}
	if response == (model.ReadTodoResponse{}) {
		answerResourceNotFound(context, fmt.Sprintf("no todo with ID : %v", ID))
		return
	}

	jsonResponse := TodoReadJSONResponse{response.ID, response.Title, response.Description, response.CreationDate.Unix(), response.DueDate.Unix()}
	context.JSON(http.StatusOK, jsonResponse)
}
