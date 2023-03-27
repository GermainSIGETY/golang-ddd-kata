Feature: Failure recovery for Todos that has to be notified, but not yet notified

  Background:
    Given an empty Database
    And a Todo with ID 30, title "Deliver Rambo", a creation date "2024-04-29", an assignee "joe@biden.com" and notified flag to "false"
    And a Todo with ID 31, title "Deliver ... princess Zelda", a creation date "2024-04-29", an assignee "joe@biden.com" and notified flag to "false"
    And a Todo with ID 32, title "Deliver Belgium", a creation date "1944-04-29", an assignee "delano@rosevelt.com" and notified flag to "true"

  Scenario: failure recovery
    Then NotificationService is called 2 times after 1000 ms
