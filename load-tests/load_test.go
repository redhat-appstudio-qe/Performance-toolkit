package main

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/redhat-appstudio-qe/performance-toolkit/load-tests/controller"
)

func TestFeatures(t *testing.T) {
  suite := godog.TestSuite{
    ScenarioInitializer: controller.InitializeScenario,
	  TestSuiteInitializer: controller.InitializeTestSuite, 
    Options: &godog.Options{
      Format:   "pretty",
      Paths:    []string{"features"},
      TestingT: t, // Testing instance that will run subtests.
    },
  }

  if suite.Run() != 0 {
    t.Fatal("non-zero status returned, failed to run feature tests")
  }
}