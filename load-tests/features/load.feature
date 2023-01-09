Feature: Load Tests For AppStudio
  
  Scenario: Validate Creation on Application with quarkus component
    Given system is running
    Then I should wait for 10 seconds
    #create a toolchain user here
    When I should create 2 appstudio users
    Then I should wait for 5 seconds
    Then I should create user resources with "Quarkus" component
    # When I should create test namespace "chaostests"
    Then I should be able to print metrics
    # When I should create Has Application "chaos1"
    # Then I should validate Has Application "chaos1"
    # When I should create Has component detection query "cmp-1" with "Quarkus"
    # Then I should wait for 10 seconds
    # Then I should validate Has component detection query for "cmp-1"
    # When I should create Has component "cmp-1"
    # Then I should validate Has component "cmp-1"

    
