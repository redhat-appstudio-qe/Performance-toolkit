package state

import (
	"context"
	"errors"

	config "github.com/redhat-appstudio-qe/performance-toolkit/config"
	"github.com/redhat-appstudio-qe/performance-toolkit/expirements"
	utils "github.com/redhat-appstudio-qe/performance-toolkit/utils"
	"github.com/redhat-appstudio/e2e-tests/pkg/framework"
)



func IsSystemRunning(ctx context.Context) (context.Context, error) {
	return ctx,nil
  }

func SteadyState(ctx context.Context) (context.Context, error) {
	ctx, err := IsSystemRunning(ctx)
	if err != nil {
		errors.New("Something Worng with system!")
	}
	// get a new framework

	framework, err := framework.NewFramework()
	if err != nil {
		return ctx, errors.New("Connection Error")
	}

	ctx = context.WithValue(ctx, "framework", framework)

	//create a namespace to run chaos 

	creatednamespace, err := framework.CommonController.CreateTestNamespace(config.TEST_NAMESPACE)
  	if err != nil {
    	return ctx, errors.New("cannot create ns")
  	}	
    
	ctx = utils.SetTestNamespace(creatednamespace.Name, ctx)

	Explist := expirements.CreateExpirementsList(ctx)

	ctx = context.WithValue(ctx, "ExperimentList", Explist)

	ctx = expirements.PopulateEnvVars(ctx)
	
	return ctx, nil


}