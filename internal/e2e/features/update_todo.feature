Feature: When user update a Todo, Todo is saved

  Background:
    Given a Todo with ID 14, title "Go on Holidays", a description "without smartphone", a creation date "2023-04-29" and a due date "2024-08-04"

  Scenario: user updates and reads Todo
    When I send a "PUT" request to "/todos/14" with JSON:
    """
    {
        "title": "Go on Holidays in september",
        "description": "without laptop",
        "dueDate": 1693591248
    }
    """
    And the response code should be 204
    And I send a "GET" request to "/todos/14"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "creationDate": 1682726400,
        "description": "without laptop",
        "dueDate": 1693591248,
        "id": 14,
        "title": "Go on Holidays in september"
      }
      """

  Scenario: user updates a non existing Todo
    When I send a "PUT" request to "/todos/400000" with JSON:
    """
    {
        "title": "Go on Holidays in september",
        "description": "without laptop",
        "dueDate": 1693591248
    }
    """
    And the response code should be 422
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

  Scenario: user updates a Todo with malformed Json
    When I send a "PUT" request to "/todos/400000" with JSON:
    """
    {badJson}
    """
    And the response code should be 400
    And the response should match json:
      """
      {
        "errors": [
          {
            "code": "BAD_REQUEST",
            "message": "invalid character 'b' looking for beginning of object key string"
          }
        ]
      }
      """