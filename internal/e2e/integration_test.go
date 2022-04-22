package e2e

import (
	"os"
	"testing"
	"time"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/bootstrap"
	"github.com/cucumber/godog"
)

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	go bootstrap.LaunchAppForIntegrationTest()
	time.Sleep(time.Second)
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	var world TodoWorld

	ctx.BeforeScenario(func(*godog.Scenario) {
		world = TodoWorld{}
	})

	ctx.Step(`^a Todo with title "([^"]*)", a description "([^"]*)" and a due date "([^"]*)"$`, world.ATodoWithTitleADescriptionAndADueDate)
	ctx.Step(`^title is "([^"]*)", description is "([^"]*)" and a due date is "([^"]*)"$`, world.TitleIsDescriptionIsAndADueDateIs)
	ctx.Step(`^User read previously created Todo$`, world.UserReadPreviouslyCreatedTodo)
	ctx.Step(`^application answers with status code (\d+)$`, world.ApplicationAnswersWithStatusCode)
	ctx.Step(`^User read todo with ID (\d+)$`, world.UserReadTodoWithID)
	ctx.Step(`^User updates previously created Todo with title "([^"]*)", description "([^"]*)" and due date "([^"]*)"$`, world.UserUpdatesPreviouslyCreatedTodoWithTitleDescriptionAndDueDate)
	ctx.Step(`^User updates todo with ID (\d+)$`, world.UserUpdatesTodoWithID)
	ctx.Step(`^User deletes previously created Todo$`, world.UserDeletesPreviouslyCreatedTodo)
	ctx.Step(`^User deletes todo with ID (\d+)$`, world.UserDeletesTodoWithID)
	ctx.Step(`^answer contains more than (\d+) Todos$`, world.AnswerContainsMoreThanTodos)
	ctx.Step(`^User reads todoList$`, world.UserReadsTodoList)
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "pretty",
		Paths:     []string{"features"},
		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	}

	status := godog.TestSuite{
		Name:                 "Todos integration tests suite",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
