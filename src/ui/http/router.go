package http

import (
	"github.com/GermainSIGETY/golang-ddd-kata/src/api"
	"github.com/GermainSIGETY/golang-ddd-kata/src/ui/http/todo"
	"github.com/gin-gonic/gin"
)

func NewRouter(api api.TodosAPI) {

	r := gin.Default()

	r.POST("/todos", func(c *gin.Context) {
		todo.HandleCreate(c, api)
	})

	r.GET("/todos", func(c *gin.Context) {
		todo.HandleReadTodoList(c, api)
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		todo.HandleReadTodo(c, ID, api)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		todo.HandleUpdate(c, ID, api)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		ID := c.Param("id")
		todo.HandleDelete(c, ID, api)
	})

	r.Run()

}
