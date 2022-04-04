Feature: When user update a Todo, Todo is saved

  Background:
    Given a Todo with title "Go to the greengrocer", a description "Buy some pickles" and a due date "2018-01-31"

  Scenario: user creates 2 Todos and reads list
    Given a Todo with title "Drink a Mojito", a description "with a straw" and a due date "2025-05-31"
    And a Todo with title "Eat a Vodka Tonic", a description "with ice" and a due date "2016-01-31"
    When User reads todoList
    Then application answers with status code 200
    And answer contains more than 2 Todos
