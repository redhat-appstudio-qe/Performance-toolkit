package controller

import (
	"context"
	"errors"

	"github.com/cucumber/godog"
	"github.com/redhat-appstudio-qe/performance-toolkit/load-tests/steps"
	"github.com/redhat-appstudio-qe/performance-toolkit/metrics"
	"github.com/redhat-appstudio/e2e-tests/pkg/framework"
)

var (
	kcp_config string
	cluster_config string
)
func BeforeSuite(){
	//Add logic here!
	//kcp_config = utils.CheckVarExistsAndReturn("KCP_KUBECONFIG")
	//cluster_config = utils.CheckVarExistsAndReturn("CLUSTER_KUBECONFIG")
}

func BeforeScenarioHook(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	//ctx = context.WithValue(ctx, "kcp_config", kcp_config)
	//ctx = context.WithValue(ctx, "cluster_config", cluster_config)
	framework, err := framework.NewFramework()
	if err != nil {
		return ctx, errors.New("Connection Error")
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
	//sc.Step(`^I should create test namespace "([^"]*)"$`, steps.CreateNS)
	sc.Step(`^I should create (\d+) appstudio users$`, steps.CreateUsers)
	sc.Step(`^I should create user resources with "([^"]*)" component$`, steps.CreateResources)
	sc.Step(`^I should be able to print metrics`, steps.PrintMetrics)
	sc.Step(`^I should create Has Application "([^"]*)"$`, steps.CreateHasApp)
	sc.Step(`^I should validate Has Application "([^"]*)"$`, steps.CheckHasApp)
	sc.Step(`^I should wait for (\d+) min$`, steps.WaitforMins)
	sc.Step(`^I should wait for (\d+) seconds$`, steps.WaitforSecs)
	sc.Step(`^I should create Has component detection query "([^"]*)" with "([^"]*)"$`, steps.CreateCompoenetDetectionQuery)
	sc.Step(`^I should validate Has component detection query for "([^"]*)"$`, steps.CheckComponentDetectionQuery)
	sc.Step(`^I should create Has component "([^"]*)"$`, steps.CreateComponent)
	sc.Step(`^I should validate Has component "([^"]*)"$`, steps.CheckComponent)

}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(BeforeSuite)
	ctx.AfterSuite(AfterSuite)
}
  