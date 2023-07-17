@AddQuestion
Feature: Add a new question using the "add" subcommand

  @UserInput
  Scenario: User adds a new question
    Given the Go program is running
    When the user executes the "add" subcommand
    Then the program should prompt the user to enter question details
    And the user provides the question, solution, and hint

  @SaveQuestion @FileCreation
  Scenario: User saves the question
    Given the user has provided the question details
    When the user confirms to save the question
    Then a markdown file is created in the "questions" directory
    And the markdown file contains the user's question, solution, and hint
