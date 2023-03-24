package bootstrap

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/infrastructure"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/ui"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func InitApp() *gin.Engine {
	router, _ := launchApp("/tmp/Todos.db", false)
	return router
}

func InitAppForIntegrationTest() (*gin.Engine, port.ITodosRepository) {
	return launchApp("/tmp/TodosTest.db", true)
}

func launchApp(url string, drop bool) (*gin.Engine, port.ITodosRepository) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	repository, err := infrastructure.NewTodosRepository(url, drop)
	if err != nil {
		log.Fatal().Err(err).Msg("Error during get Things done initialization")
	}
	todosAPI := api.NewApi(repository)
	router := ui.NewRouter(todosAPI)

	return router, repository
}
