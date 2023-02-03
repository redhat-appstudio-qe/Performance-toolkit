package expirements

import (
	"context"
	config "github.com/redhat-appstudio-qe/performance-toolkit/config"
)


func CreateExpirementsList() []config.Expirement{
	return ExperimentList;
}

func PopulateGlobalEnvVars(ctx context.Context) context.Context{

	// add global vars here
	return ctx
}
