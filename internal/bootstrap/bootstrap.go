package bootstrap

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/infrastructure"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/ui"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func LaunchApp() {
	launchApp("/tmp/Todos.db", false)
}

func LaunchAppForIntegrationTest() {
	launchApp("/tmp/TodosTest.db", true)
}

func launchApp(url string, drop bool) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	repository, err := infrastructure.NewTodosRepository(url, drop)
	if err != nil {
		log.Fatal().Err(err).Msg("Error during locationSearchApi instantiation")
	}
	todosAPI := api.NewApi(repository)
	ui.NewRouter(todosAPI)
}
