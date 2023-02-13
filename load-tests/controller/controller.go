package controller

import (
	"context"
	"errors"

	"github.com/cucumber/godog"
	"github.com/redhat-appstudio-qe/performance-toolkit/load-tests/steps"
	"github.com/redhat-appstudio-qe/performance-toolkit/metrics"
	"github.com/redhat-appstudio/e2e-tests/pkg/framework"
)

func BeforeSuite(){
	//Add logic here!
}

func BeforeScenarioHook(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	framework, err := framework.NewFramework()
	if err != nil {
		return ctx, errors.New("connection error")
	}

	
	ctx = context.WithValue(ctx, "framework", framework)

	closeMetrics, metricsInstance := metrics.StartCollection(ctx)

	ctx = context.WithValue(ctx, "closeMetrics", closeMetrics)
	ctx = context.WithValue(ctx, "metricsInstance", metricsInstance)
	return ctx, nil
}

func AfterSuite() {
	// Add Cleanup logic!
}


func InitializeScenario(sc *godog.ScenarioContext) {

	sc.Before(BeforeScenarioHook)
	sc.Step(`^system is running`, steps.IsPhysicalSystemRunning)
	sc.Step(`^I should Configure Batch Concurent Tests with max requests (\d+) and (\d+) batches$`, steps.ConfigureBatchConcurentTests)
	sc.Step(`^I should run Batch Concurent Tests with "([^"]*)"$`, steps.StartBatchConcurentTests)
	sc.Step(`^I should ramp up users with batch controller$`, steps.StartBatchConcurentUserTests)
	sc.Step(`^I should Configure Infinite Concurent Tests with RPS (\d+) and timeout of (\d+) secs$`, steps.ConfigureInfiniteConcurentTests)
	sc.Step(`^I should run Infinite Concurent Tests with "([^"]*)"$`, steps.StartInfiniteConcurentTests)
	sc.Step(`^I should ramp up users with infinite controller$`, steps.StartInfiniteConcurentUserTests)
	sc.Step(`^I should Configure Spike Concurent Tests with max RPS (\d+) and timeout of (\d+) secs$$`, steps.ConfigureSpikeConcurentTests)
	sc.Step(`^I should run Spike Concurent Tests with "([^"]*)"$`, steps.StartSpikeConcurentTests)
	sc.Step(`^I should ramp up users with spike controller$`, steps.StartSpikeConcurentUserTests)
	sc.Step(`^I should Stop And Print Metrics`, steps.PrintMetrics)

}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(BeforeSuite)
	ctx.AfterSuite(AfterSuite)
}
  