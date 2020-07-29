Feature: When user deletes a Todo, todo is actually deleted

  Background:
    Given a Todo with title "escape from prison", a description "without weapon" and a due date "2021-08-04"

  Scenario: user deletes Todo
    When User deletes previously created Todo
    Then application answers with status code 204
    Then User read previously created Todo
    And application answers with status code 404

  Scenario: user deletes non Existent Todo
    When User deletes todo with ID 123432
    Then application answers with status code 422