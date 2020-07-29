Feature: When user update a Todo, Todo is saved

  Background:
    Given a Todo with title "Go to the greengrocer", a description "Buy some pickles" and a due date "2018-01-31"

  Scenario: user read Todo
    When User read previously created Todo
    Then application answers with status code 200
    And title is "Go to the greengrocer", description is "Buy some pickles" and a due date is "2018-01-31"

  Scenario: user read non Existent Todo
    When User read todo with ID 123432
    Then application answers with status code 404