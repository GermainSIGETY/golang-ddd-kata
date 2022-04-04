package bootstrap

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo"
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
	repository := infrastructure.TodosRepository{}
	if err := repository.InitDatabase(url, drop); err != nil {
		panic(err)
	}
	todosAPI := todo.NewApi(&repository)
	ui.NewRouter(todosAPI)
}
