package ui

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
)

// HandleDelete godoc
// @Summary Delete a todo
// @Description Delete a todo by its Id
// @id delete-todo
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 204 "it's ok"
// @Failure 422 {object} ErrorsArrayJsonResponse
// @Failure 500 {object} ErrorsArrayJsonResponse
// @Router /todos/{id} [delete]
func handleDelete(context *gin.Context, IDAsString string) {
	ID, err := strconv.Atoi(IDAsString)
	if err != nil {
		answerBadRequest(context, "todo ID in path must be an integer")
		return
	}

	deleteErrors := api.GetTodoApi().DeleteTodo(ID)
	if deleteErrors != nil {
		answerError(context, deleteErrors)
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
