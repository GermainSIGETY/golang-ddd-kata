Feature: When user update a Todo, Todo is saved

  Background:
    Given an empty Database
    And a Todo with ID 24, title "Drink a Mojito", a description "with a straw", a creation date "2024-04-29" and a due date "2024-08-04"
    And a Todo with ID 25, title "Sip a Vodka Tonic", a description "with ice", a creation date "2025-04-29" and a due date "2025-08-04"


  Scenario: user reads list
    And I send a "GET" request to "/todos"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "todos": [
          {
            "dueDate": 1722729600,
            "id": 24,
            "title": "Drink a Mojito"
          },
          {
            "dueDate": 1754265600,
            "id": 25,
            "title": "Sip a Vodka Tonic"
          }
        ]
      }
      """