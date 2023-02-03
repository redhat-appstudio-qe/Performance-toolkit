Feature: Load Tests For AppStudio
  
  Scenario: Validate Creation on Application with quarkus component
    Given system is running
    Then I should wait for 10 seconds
    When I should create 1 appstudio users
    Then I should wait for 5 seconds
    Then I should create user resources with "node" component
    Then I should be able to print metrics

    
