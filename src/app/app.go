package app

import (
	"github.com/GermainSIGETY/golang-ddd-kata/src/api"
	"github.com/GermainSIGETY/golang-ddd-kata/src/repository"
	"github.com/GermainSIGETY/golang-ddd-kata/src/ui/http"
)

func LaunchApp() {
	launchApp("/tmp/Todos.db", false)
}

func LaunchAppForIntegrationTest() {
	launchApp("/tmp/TodosTest.db", true)
}

func launchApp(url string, drop bool) {

	repository := repository.TodosRepository{}
	if err := repository.InitDatabase(url, drop); err != nil {
		panic(err)
	}
	todosAPI := api.New(&repository)
	http.NewRouter(todosAPI)
}
