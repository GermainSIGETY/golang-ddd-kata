package steps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/src/ui/http/todo"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (world *TodoWorld) ATodoWithTitleADescriptionAndADueDate(title, description string, date string) (err error) {

	_, dueDate := stringToDate(date)

	req := todo.TodoCreationJSONRequest{Title: &title, Description: &description, DueDate: &dueDate}

	body, _ := json.Marshal(req)

	res, errPost := http.Post(serverURL, todo.JSONContentType, bytes.NewBuffer(body))
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

	var jsonAnswer todo.TodoCreationJSONResponse
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

	var jsonAnswer todo.TodoReadJSONResponse
	errUnmarshall := json.Unmarshal(answer, &jsonAnswer)
	if errUnmarshall != nil {
		return fmt.Errorf("error deserializing body : %v", errUnmarshall)
	}

	world.title = jsonAnswer.Title
	if jsonAnswer.Description != nil {
		world.description = *jsonAnswer.Description
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
