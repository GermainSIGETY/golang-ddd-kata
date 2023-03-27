package bootstrap

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/api"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/services"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/infrastructure"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/ui"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func InitApp() *gin.Engine {
	notificationSender := infrastructure.NewNotificationSender()
	router, _ := launchApp("/tmp/Todos.db", false, notificationSender, 20)
	return router
}

func InitAppForIntegrationTest(notificationSender port.INotificationSender) (*gin.Engine, port.ITodosRepository) {
	return launchApp("/tmp/TodosTest.db", true, notificationSender, 1)
}

func launchApp(url string, drop bool, notificationSender port.INotificationSender, notificationFailureTickerTime int) (*gin.Engine, port.ITodosRepository) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	repository, err := infrastructure.NewTodosRepository(url, drop)
	if err != nil {
		log.Fatal().Err(err).Msg("Error during get Things done initialization")
	}
	notificationChannel := services.NewNotificationService(repository, notificationSender, notificationFailureTickerTime)
	todosAPI := api.NewApi(repository, notificationChannel)
	router := ui.NewRouter(todosAPI)
	return router, repository
}
