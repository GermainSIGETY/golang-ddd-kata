package e2e

import (
	"fmt"
	"net/http"
	"strconv"
)

func (world *TodoWorld) UserDeletesPreviouslyCreatedTodo() error {
	resp, err := callHttpDelete(world.todoID)
	if err != nil {
		return fmt.Errorf("error on DELETE todo %v", err)
	}
	world.statusCode = resp.StatusCode
	return nil
}

func (world *TodoWorld) UserDeletesTodoWithID(ID int) error {
	resp, err := callHttpDelete(ID)
	if err != nil {
		return fmt.Errorf("error on DELETE todo %v", err)
	}
	world.statusCode = resp.StatusCode
	return nil
}

func callHttpDelete(ID int) (*http.Response, error) {

	client := &http.Client{}
	req, errDelete := http.NewRequest(http.MethodDelete, serverURL+"/"+strconv.Itoa(ID), nil)
	if errDelete != nil {
		return nil, errDelete
	}

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	return resp, nil
}
