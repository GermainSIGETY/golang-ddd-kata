package e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/bootstrap"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/port"
	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

var (
	router                   *gin.Engine
	repository               port.ITodosRepository
	resp                     *httptest.ResponseRecorder
	mockedNotificationSender *port.NotificationSenderMock
)

func initApp() {

	mockedNotificationSender = &port.NotificationSenderMock{}
	router, repository = bootstrap.InitAppForIntegrationTest(mockedNotificationSender)
	log.Info().Msg("app init")
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	fmt.Println("Before test suite")
	ctx.BeforeSuite(initApp)
	ctx.ScenarioContext().Before(beforeScenario)
}

func beforeScenario(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	mockedNotificationSender.ResetCallsCounter()
	return ctx, nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	fmt.Println("Init scenarios")

	ctx.Step(`^an empty Database$`, anEmptyDatabase)
	ctx.Step(`^a Todo with ID (\d+), title "([^"]*)", a description "([^"]*)", a creation date "([^"]*)" and a due date "([^"]*)"$`, aTodoWithTitleADescriptionAndADueDate)
	ctx.Step(`^a Todo with ID (\d+), title "([^"]*)", a creation date "([^"]*)", an assignee "([^"]*)" and notified flag to "([^"]*)"$`, aTodoWithTitleAssigneeAndFlag)
	ctx.Step(`^I send a "([^"]*)" request to "([^"]*)"$`, iSendARequestTo)
	ctx.Step(`^I send a "([^"]*)" request to "([^"]*)" with JSON:$`, iSendARequestToWithJSON)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, theResponseShouldMatchJSON)
	ctx.Step(`^NotificationService is called (\d+) times after (\d+) ms$`, notificationServiceIsCalled)

}

func anEmptyDatabase() {
	repository.EmptyDatabaseForTests()
}

func aTodoWithTitleADescriptionAndADueDate(todoID int, title, description string, creationDate string, dueDate string) (err error) {
	creationDateInt64, _ := stringToDate(creationDate)
	dueDateInt64, _ := stringToDate(dueDate)
	return repository.UpdateTodo(model.Todo{
		ID:               todoID,
		CreationDate:     time.Unix(creationDateInt64, 0),
		Description:      description,
		DueDate:          time.Unix(dueDateInt64, 0),
		Title:            title,
		NotificationSent: true,
	})
}

func aTodoWithTitleAssigneeAndFlag(todoID int, title, creationDate string, assignee string, notifiedFlag string) (err error) {
	creationDateInt64, _ := stringToDate(creationDate)
	notifiedFlagBool, err := strconv.ParseBool(notifiedFlag)
	if err != nil {
		return err
	}
	return repository.UpdateTodo(model.Todo{
		ID:               todoID,
		Title:            title,
		CreationDate:     time.Unix(creationDateInt64, 0),
		Assignee:         assignee,
		NotificationSent: notifiedFlagBool,
	})
}

func iSendARequestTo(action, endpoint string) error {
	fmt.Printf("Send request to %s\n", endpoint)
	ds := &godog.DocString{}
	ds.Content = "{}"
	return iSendARequestToWithJSON(action, endpoint, ds)
}

func iSendARequestToWithJSON(action string, endpoint string, body *godog.DocString) (err error) {
	// create a request
	req, _ := http.NewRequest(action, endpoint, strings.NewReader(body.Content))

	// create a response recorder
	resp = httptest.NewRecorder()

	// execute de request
	router.ServeHTTP(resp, req)

	return nil
}

func theResponseCodeShouldBe(code int) error {
	if code != resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d with body: %v", code, resp.Code, resp.Body)
	}
	return nil
}

func theResponseShouldMatchJSON(body *godog.DocString) (err error) {
	var expected, actual []byte
	var exp, act interface{}

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &exp); err != nil {
		return
	}
	if expected, err = json.MarshalIndent(exp, "", "  "); err != nil {
		return
	}

	// re-encode actual response too
	if err = json.Unmarshal(resp.Body.Bytes(), &act); err != nil {
		return
	}
	if actual, err = json.MarshalIndent(act, "", "  "); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if len(actual) != len(expected) {
		return fmt.Errorf(
			"expected json length: %d does not match actual: %d:\n%s",
			len(expected),
			len(actual),
			string(actual),
		)
	}

	for i, b := range actual {
		if b != expected[i] {
			return fmt.Errorf(
				"expected JSON does not match actual, showing up to last matched character:\n%s",
				string(actual[:i+1]),
			)
		}
	}
	return nil
}

func stringToDate(date string) (int64, error) {
	t, err := time.Parse(layoutISO, date)
	dueDate := t.Unix()
	return dueDate, err
}

func notificationServiceIsCalled(awaitedCalls int, timeToWait int) error {
	//dummy wait in order to wait message in channel to be received
	time.Sleep(time.Millisecond * time.Duration(timeToWait))
	if awaitedCalls != mockedNotificationSender.Calls {
		return fmt.Errorf("expected calls to notificationService: %d, but actual is: %d", awaitedCalls, mockedNotificationSender.Calls)
	}
	return nil
}
