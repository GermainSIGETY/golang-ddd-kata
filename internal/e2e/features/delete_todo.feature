Feature: When user deletes a Todo, todo is actually deleted

  Background:
    Given a Todo with ID 13, title "Escape from prison", a description "without weapon", a creation date "2021-01-31" and a due date "2021-03-12"

  Scenario: user deletes an existing Todo
    When I send a "DELETE" request to "/todos/13"
    Then the response code should be 204


  Scenario: user deletes a non existing Todo
    When I send a "DELETE" request to "/todos/40000"
    Then the response code should be 422
    And the response should match json:
      """
      {
        "errors": [
          {
            "code": "INVALID_TODO_ID",
            "field": "id",
            "message": "no existing todo with this id"
          }
        ]
      }
      """
