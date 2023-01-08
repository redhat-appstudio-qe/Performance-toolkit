package expirements

import (
	common "github.com/redhat-appstudio-qe/performance-toolkit/common"
	config "github.com/redhat-appstudio-qe/performance-toolkit/config"
	application "github.com/redhat-appstudio-qe/performance-toolkit/expirements/application"
	"github.com/redhat-appstudio-qe/performance-toolkit/expirements/infrastructure"
	"github.com/redhat-appstudio-qe/performance-toolkit/probes"
)


var ExperimentList = []config.Expirement {
	{
		Name: "Application Service Pod Delete - Test", 
		Probe: probes.ProbeCreateHASApplication,
		Inject: application.DeletePodExperiment,
		ChaosIteration: 5,
		ProbeIntervalSecs: 28,
		Before: common.ApplicationServicePodDeleteBefore,
		After: common.ApplicationServicePodDeleteAfter,
		Weight: 10,

	},
	//Generated New
	{
		Name: "Nodespike", 
		Probe: probes.NodespikeProbe,
		Inject: infrastructure.NodespikeExperiment,
		ChaosIteration: 2,
		ProbeIntervalSecs: 10,
		Before: common.NodespikeBefore,
		After: common.NodespikeAfter,
		Weight: 10,
	},
}