package steps

import (
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/src/ui/http/todo"
	"time"
)

const (
	serverURL = "http://localhost:8080/todos"
	layoutISO = "2006-01-02"
)

type TodoWorld struct {
	todoID        int
	title         string
	description   string
	dueDate       int64
	statusCode    int
	todoSummaries []todo.TodoSummaryJSONResponse
}

func stringToDate(date string) (error, int64) {
	t, err := time.Parse(layoutISO, date)
	dueDate := t.Unix()
	return err, dueDate
}

func (world *TodoWorld) ApplicationAnswersWithStatusCode(statusCode int) (err error) {
	if world.statusCode != statusCode {
		err = fmt.Errorf("error status code %v is not %v", world.statusCode, statusCode)
	}
	return err
}

func (world *TodoWorld) TitleIsDescriptionIsAndADueDateIs(title, description string, dueDate string) (err error) {

	if world.title != title {
		return fmt.Errorf("error title %v is not %v", world.title, title)
	}
	if world.description != description {
		return fmt.Errorf("error description %v is not %v", world.description, description)
	}
	_, date := stringToDate(dueDate)
	if world.dueDate != date {
		return fmt.Errorf("error date %v is not %v", world.dueDate, date)
	}
	return
}
