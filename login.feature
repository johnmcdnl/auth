#noinspection CucumberUndefinedStep
Feature: Login

  Background: Login
    Given I have a valid user

  Scenario: Valid Login
    When I request "POST" "/auth/login" with request body
    """
    {
      "username" : "{username}",
      "password" : "{password}"
    }
    """
    Then I get a 200 response
    And I get a response body with a valid JWT

  Scenario: User does not exist
    When I request "POST" "/auth/login" with request body
    """
    {
      "username" : "madeUpUser",
      "password" : "{password}"
    }
    """
    Then I get a 401 response

  Scenario: Incorrect Password
    When I request "POST" "/auth/login" with request body
    """
    {
      "username" : "{username}",
      "password" : "badpassword"
    }
    """
    Then I get a 401 response

  Scenario: No Username provided
    When I request "POST" "/auth/login" with request body
    """
    {
      "password" : "{password}"
    }
    """
    Then I get a 400 response

  Scenario: No Password provided
    When I request "POST" "/auth/login" with request body
    """
    {
      "username" : "{username}"
    }
    """
    Then I get a 400 response

  Scenario: Invalid JSON provided
    When I request "POST" "/auth/login" with request body
    """
      "username" : "{username}"
    """
    Then I get a 400 response

  Scenario: No Body provided
    When I request "POST" "/auth/login"
    Then I get a 400 response
