package bootstrap

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/infrastructure"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/ui"
)

func LaunchApp() {
	launchApp("/tmp/Todos.db", false)
}

func LaunchAppForIntegrationTest() {
	launchApp("/tmp/TodosTest.db", true)
}

func launchApp(url string, drop bool) {
	repository, err := infrastructure.NewTodosRepository(url, drop)
	if err != nil {
		panic(err)
	}
	todosAPI := api.NewApi(repository)
	ui.NewRouter(todosAPI)
}
