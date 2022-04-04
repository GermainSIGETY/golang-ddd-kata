package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type updateJsonRequest struct {
	Title       *string `json:"title" validate:"required"`
	Description *string `json:"description"`
	DueDate     *int64  `json:"dueDate" validate:"required"`
}

func (world *TodoWorld) UserUpdatesPreviouslyCreatedTodoWithTitleDescriptionAndDueDate(title, description, dueDateAsString string) error {
	_, dueDate := stringToDate(dueDateAsString)

	jsonRequest := updateJsonRequest{&title, &description, &dueDate}

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

	jsonRequest := updateJsonRequest{&title, &description, &dueDate}
	resp, err := callHttpPut(ID, jsonRequest)
	if err != nil {
		return fmt.Errorf("error on PUT todo %v", err)
	}
	world.statusCode = resp.StatusCode
	return nil
}

func callHttpPut(ID int, request updateJsonRequest) (*http.Response, error) {
	requestBody, _ := json.Marshal(request)

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
