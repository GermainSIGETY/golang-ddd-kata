package ui

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/gin-gonic/gin"
)

func NewRouter(api api.TodosAPI) {

	r := gin.Default()

	r.POST("/todos", func(c *gin.Context) {
		handleCreate(c, api)
	})

	r.GET("/todos", func(c *gin.Context) {
		HandleReadTodoList(c, api)
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		handleReadTodo(c, ID, api)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		handleUpdate(c, ID, api)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		handleDelete(c, ID, api)
	})
	r.Run()
}
