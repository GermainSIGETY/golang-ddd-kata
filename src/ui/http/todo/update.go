package todo

import (
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/src/api"
	"github.com/GermainSIGETY/golang-ddd-kata/src/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TodoUpdateJSONRequest struct {
	Title       *string `json:"title" validate:"required"`
	Description *string `json:"description"`
	DueDate     *int64  `json:"dueDate" validate:"required"`
}

// HandleUpdate godoc
// @Summary Update a todo
// @Description Update a todo by its Id
// @id update-todo
// @Tags todos
// @Accept json
// @Param id path int true "Todo ID"
// @Param body body TodoCreationJSONRequest true "todo infos"
// @Success 204 "it's ok"
// @Failure 422 {object} ErrorsArrayJsonResponse
// @Failure 500 {object} ErrorsArrayJsonResponse
// @Router /todos/{id} [put]
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
