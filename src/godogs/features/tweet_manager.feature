Feature: publish a text tweet
  I need publish a new text tweet

  Scenario: 
    Given there is a user "francoagarcia"
    When he wants to tweet "Hola mundo"
    Then there should be 1 new tweet for the user