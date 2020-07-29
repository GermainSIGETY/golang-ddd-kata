package ui

import (
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
