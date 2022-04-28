package ui

import (
	"github.com/gin-gonic/gin"
)

func InitAndRunRouter() {

	r := gin.Default()

	r.POST("/todos", func(c *gin.Context) {
		handleCreate(c)
	})

	r.GET("/todos", func(c *gin.Context) {
		HandleReadTodoList(c)
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		handleReadTodo(c, ID)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		handleUpdate(c, ID)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		handleDelete(c, ID)
	})
	r.Run()
}
