package expirements

import (
	common "github.com/redhat-appstudio-qe/performance-toolkit/common"
	config "github.com/redhat-appstudio-qe/performance-toolkit/config"

	//application "github.com/redhat-appstudio-qe/performance-toolkit/expirements/application"
	//"github.com/redhat-appstudio-qe/performance-toolkit/expirements/infrastructure"
	"github.com/redhat-appstudio-qe/performance-toolkit/expirements/network"
	"github.com/redhat-appstudio-qe/performance-toolkit/probes"
)


var ExperimentList = []config.Expirement {
	/* {
		Name: "Application Service Pod Delete - Test", 
		Probe: probes.ProbeCreateHASApplication,
		Inject: application.DeletePodExperiment,
		ChaosIteration: 5,
		ProbeIntervalSecs: 28,
		Before: common.ApplicationServicePodDeleteBefore,
		After: common.ApplicationServicePodDeleteAfter,
		Weight: 10,

	}, */

	//Generated New
	/* {
		Name: "NodeDrain", 
		Probe: probes.NodeDrainProbe,
		Inject: infrastructure.NodeDrainExperiment,
		ChaosIteration: 0,
		ProbeIntervalSecs: 30,
		Before: common.NodeDrainBefore,
		After: common.NodeDrainAfter,
		Weight: 10,
	}, */

	//Generated New
	{
		Name: "NetworkDeny", 
		Probe: probes.NetworkDenyProbe,
		Inject: network.NetworkDenyExperiment,
		ChaosIteration: 0,
		ProbeIntervalSecs: 5,
		Before: common.NetworkDenyBefore,
		After: common.NetworkDenyAfter,
		Weight: 10,
	},
}