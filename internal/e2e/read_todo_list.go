package e2e

import (
	"encoding/json"
	"fmt"
	"github.com/GermainSIGETY/golang-ddd-kata/internal/ui"
	"io/ioutil"
	"net/http"
)

func (world *TodoWorld) AnswerContainsMoreThanTodos(numberOfTodos int) error {
	if len(world.todoSummaries) < numberOfTodos {
		return fmt.Errorf("error reading todo list ; list should contains at least : %v", numberOfTodos)
	}
	return nil
}

func (world *TodoWorld) UserReadsTodoList() error {
	res, errGet := http.Get(serverURL)
	if errGet != nil {
		return fmt.Errorf("error on POST todo %v", errGet)
	}
	answer, errReadBody := ioutil.ReadAll(res.Body)
	if errReadBody != nil {
		return fmt.Errorf("error reading body : %v", errReadBody)
	}

	var jsonAnswer ui.TodoListJSONResponse
	errUnmarshall := json.Unmarshal(answer, &jsonAnswer)
	if errUnmarshall != nil {
		return fmt.Errorf("error deserializing body : %v", errUnmarshall)
	}

	world.statusCode = res.StatusCode
	world.todoSummaries = jsonAnswer.TodoList

	return nil
}
