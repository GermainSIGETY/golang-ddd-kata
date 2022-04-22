package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/ui"
)

type jsonRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     int64  `json:"dueDate"`
}

func (world *TodoWorld) ATodoWithTitleADescriptionAndADueDate(title, description string, date string) (err error) {

	dueDate, _ := stringToDate(date)

	req := jsonRequest{title, description, dueDate}

	body, _ := json.Marshal(req)

	res, errPost := http.Post(serverURL, ui.JSONContentType, bytes.NewBuffer(body))
	if errPost != nil {
		return fmt.Errorf("error on POST todo %v", errPost)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error on POST todo, server didn't answered 200, but %v", res.StatusCode)
	}

	answer, errReadBody := ioutil.ReadAll(res.Body)
	if errReadBody != nil {
		return fmt.Errorf("error reading body : %v", errReadBody)
	}

	var jsonAnswer ui.TodoCreationJSONResponse
	err = json.Unmarshal(answer, &jsonAnswer)
	if err != nil {
		return fmt.Errorf("error deserializing body : %v", err)
	}

	if jsonAnswer.ID <= 0 {
		return fmt.Errorf("POST todo, server didn't answered ID")
	}

	world.todoID = jsonAnswer.ID
	return
}

func (world *TodoWorld) UserReadPreviouslyCreatedTodo() error {

	res, errGet := http.Get(serverURL + "/" + strconv.Itoa(world.todoID))
	if errGet != nil {
		return fmt.Errorf("error on getting todo %v", errGet)
	}

	answer, errReadBody := ioutil.ReadAll(res.Body)
	if errReadBody != nil {
		return fmt.Errorf("error reading body : %v", errReadBody)
	}

	var jsonAnswer ui.TodoReadJSONResponse
	errUnmarshall := json.Unmarshal(answer, &jsonAnswer)
	if errUnmarshall != nil {
		return fmt.Errorf("error deserializing body : %v", errUnmarshall)
	}

	world.title = jsonAnswer.Title
	if jsonAnswer.Description != "" {
		world.description = jsonAnswer.Description
	}
	world.dueDate = jsonAnswer.DueDate
	world.statusCode = res.StatusCode
	return nil
}

func (world *TodoWorld) UserReadTodoWithID(todoID int) error {
	res, errGet := http.Get(serverURL + "/" + strconv.Itoa(todoID))
	if errGet != nil {
		return fmt.Errorf("error reading todo : %v", errGet)
	}
	world.statusCode = res.StatusCode
	return nil
}
