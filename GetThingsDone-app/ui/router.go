package ui

import (
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/api"
	"github.com/gin-gonic/gin"
)

func NewRouter(api api.TodosAPI) {

	r := gin.Default()

	r.POST("/todos", func(c *gin.Context) {
		HandleCreate(c, api)
	})

	r.GET("/todos", func(c *gin.Context) {
		HandleReadTodoList(c, api)
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		HandleReadTodo(c, ID, api)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		HandleUpdate(c, ID, api)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		HandleDelete(c, ID, api)
	})

	r.Run()

}
