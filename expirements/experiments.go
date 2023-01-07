package expirements

import (
	"context"

	config "github.com/redhat-appstudio-qe/performance-toolkit/config"
	applicationChaos "github.com/redhat-appstudio-qe/performance-toolkit/expirements/application"
	"github.com/redhat-appstudio-qe/performance-toolkit/probes"
	"github.com/redhat-appstudio-qe/performance-toolkit/utils"
)



func CreateExpirementsList(ctx context.Context) []config.Expirement{
	expi := []config.Expirement{
		{
			Name: "Application Service Pod Delete - Test", 
			Probe: probes.ProbeCreateHASApplication,
			Inject: applicationChaos.DeletePodExperiment,
			ChaosIteration: 5,
			ProbeIntervalSecs: 28,
		},
	}
	return expi
}

func PopulateEnvVars(ctx context.Context) context.Context{
	
	ctx = utils.SetChaosNamespace(config.CHAOS_NAMESPACE, ctx)
	ctx = utils.SetpodDeleteLabelKey(config.POD_DELETE_LABEL_KEY, ctx)
	ctx = utils.SetpodDeleteLabelValue(config.POD_DELETE_LABEL_VALUE, ctx)

	return ctx
}





