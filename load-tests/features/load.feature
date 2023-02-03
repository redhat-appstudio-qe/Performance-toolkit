Feature: Load Tests For AppStudio
  
  # Scenario: Validate Creation on Application with quarkus component
  #   Given system is running
  #   Then I should wait for 10 seconds
  #   When I should create 1 appstudio users
  #   Then I should wait for 5 seconds
  #   Then I should create user resources with "node" component
  #   Then I should be able to print metrics

  # Scenario: Run a Simple Batch Concurent Test
  #     Given system is running
  #     Then I should Configure Batch Concurent Tests with max requests 200 and 20 batches
  #     Then I should run Batch Concurent Tests with "node"
  
  Scenario: Run a Simple Infinite Concurent Test
      Given system is running
      Then I should Configure Infinite Concurent Tests with RPS 3 and timeout of 10 secs
      Then I should run Infinite Concurent Tests with "node"
  
  # Scenario: Run a Simple Spike Concurent Test
  #     Given system is running
  #     Then I should Configure Spike Concurent Tests with max RPS 30 and timeout of 100 secs
  #     Then I should run Spike Concurent Tests with "node"

    
