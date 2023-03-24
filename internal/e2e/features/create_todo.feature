Feature: When user updates a Todo, Todo is saved

  Scenario: create TODO : valid input
    When I send a "POST" request to "/todos" with JSON:
    """
    {
        "title": "Plant a tree",
        "dueDate": 1557847007,
        "description" : "and water it with 10 liters"
    }
    """
    Then the response code should be 200

  Scenario: create TODO : invalid input
    When I send a "POST" request to "/todos" with JSON:
    """
    {
        "title": "Plant a tree"
    }
    """
    Then the response code should be 422
    And the response should match json:
      """
      {
        "errors": [
          {
            "code": "EMPTY_FIELD",
            "field": "dueDate",
            "message": "please fill this field"
          }
        ]
      }
      """
