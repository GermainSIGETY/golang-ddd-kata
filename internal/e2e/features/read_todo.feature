Feature: When user update a Todo, Todo is saved

  Background:
    Given a Todo with ID 12, title "Go to the greengrocer", a description "Buy some pickles", a creation date "2021-01-31" and a due date "2021-03-12"

  Scenario: user reads existing Todo
    When I send a "GET" request to "/todos/12"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "creationDate": 1612051200,
        "description": "Buy some pickles",
        "dueDate": 1615507200,
        "id": 12,
        "title": "Go to the greengrocer"
      }
      """

  Scenario: user reads non existing Todo
    When I send a "GET" request to "/todos/40000"
    Then the response code should be 404
    And the response should match json:
      """
      {
        "errors": [
          {
            "code": "NOT_FOUND"
          }
        ]
      }
      """
