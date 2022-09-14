package bootstrap

import (
	_ "github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
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
	err := infrastructure.InitTodosRepository(url, drop)
	if err != nil {
		panic(err)
	}
	ui.InitAndRunRouter()
}
