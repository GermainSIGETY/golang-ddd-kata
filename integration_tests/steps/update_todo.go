package steps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/src/ui/http/todo"
	"net/http"
	"strconv"
)

func (world *TodoWorld) UserUpdatesPreviouslyCreatedTodoWithTitleDescriptionAndDueDate(title, description, dueDateAsString string) error {
	_, dueDate := stringToDate(dueDateAsString)

	jsonRequest := todo.TodoUpdateJSONRequest{Title: &title, Description: &description, DueDate: &dueDate}

	resp, err := callHttpPut(world.todoID, jsonRequest)
	if err != nil {
		return fmt.Errorf("error on PUT todo %v", err)
	}
	world.statusCode = resp.StatusCode
	return nil
}

func (world *TodoWorld) UserUpdatesTodoWithID(ID int) error {

	title := "toto"
	description := "toto"
	var dueDate int64 = 1234

	jsonRequest := todo.TodoUpdateJSONRequest{Title: &title, Description: &description, DueDate: &dueDate}
	resp, err := callHttpPut(ID, jsonRequest)
	if err != nil {
		return fmt.Errorf("error on PUT todo %v", err)
	}
	world.statusCode = resp.StatusCode
	return nil
}

func callHttpPut(ID int, jsonRequest todo.TodoUpdateJSONRequest) (*http.Response, error) {
	requestBody, _ := json.Marshal(jsonRequest)

	client := &http.Client{}
	req, errPut := http.NewRequest(http.MethodPut, serverURL+"/"+strconv.Itoa(ID), bytes.NewBuffer(requestBody))
	if errPut != nil {
		return nil, errPut
	}

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	return resp, nil
}
