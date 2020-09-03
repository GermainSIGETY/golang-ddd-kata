package ui

import (
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Delete a todo
// @Description Delete a todo from its id
// @Produce json
// @Param id path int true "Todo ID"
// @Success 204
// @Failure 500 {object} ErrorsArrayJsonResponse
// @Router /todos/{id} [delete]
func HandleDelete(context *gin.Context, IDAsString string, api api.TodosAPI) {
	ID, err := strconv.Atoi(IDAsString)
	if err != nil {
		answerBadRequest(context, "todo ID in path must be an integer")
		return
	}

	deleteErrors := api.DeleteTodo(ID)
	if deleteErrors != nil {
		answerError(context, deleteErrors)
		return
	}
	context.JSON(http.StatusNoContent, nil)
}
