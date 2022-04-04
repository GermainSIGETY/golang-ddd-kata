Feature: When user update a Todo, Todo is saved

  Background:
    Given a Todo with title "Go on Holidays", a description "without smartphone" and a due date "2020-08-04"

  Scenario: user updates and reads Todo
    When User updates previously created Todo with title "Go on Holidays in september", description "without laptop" and due date "2020-09-04"
    Then application answers with status code 204
    Then User read previously created Todo
    And application answers with status code 200
    And title is "Go on Holidays in september", description is "without laptop" and a due date is "2020-09-04"

  Scenario: user updates non Existent Todo
    When User updates todo with ID 123432
    Then application answers with status code 422