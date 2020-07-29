package app

import (
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-app/ui"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-infra/repository"
	"github.com/GermainSIGETY/golang-ddd-kata/GetThingsDone-todo-domain/api"
)

func LaunchApp() {
	launchApp("sqlite3", "/tmp/Todos.db", false)
}

func LaunchAppForIntegrationTest() {
	launchApp("sqlite3", "/tmp/TodosTest.db", true)
}

func launchApp(dialect string, url string, drop bool) {

	repository := repository.TodosRepository{}

	if err := repository.InitDatabase(dialect, url, drop); err != nil {
		panic(err)
	}

	todosAPI := api.New(&repository)

	ui.NewRouter(todosAPI)
}
