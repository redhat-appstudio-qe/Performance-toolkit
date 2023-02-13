Feature: Load Tests For AppStudio

  Scenario: Only Create Users 
    Given system is running
    Then I should Configure Infinite Concurent Tests with RPS 15 and timeout of 10 secs
    Then I should ramp up users with infinite controller
    Then I should Stop And Print Metrics

  Scenario: Run a Simple dotnet Batch Concurent Test
      Given system is running
      Then I should Configure Batch Concurent Tests with max requests 200 and 20 batches
      Then I should run Batch Concurent Tests with "dotnet"
      Then I should Stop And Print Metrics
  
  Scenario: Run a Simple Quarkus Infinite Concurent Test
      Given system is running
      Then I should Configure Infinite Concurent Tests with RPS 25 and timeout of 5 secs
      Then I should run Infinite Concurent Tests with "Quarkus"
      Then I should Stop And Print Metrics
  
  Scenario: Run a Simple node Infinite Concurent Test
    Given system is running
    Then I should Configure Infinite Concurent Tests with RPS 2 and timeout of 1 secs
    Then I should run Infinite Concurent Tests with "node"
    Then I should Stop And Print Metrics
  
  Scenario: Run a Simple python Spike Concurent Test
    Given system is running
    Then I should Configure Spike Concurent Tests with max RPS 5 and timeout of 1 secs
    Then I should run Spike Concurent Tests with "python"
    Then I should Stop And Print Metrics

    
