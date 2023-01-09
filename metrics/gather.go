package metrics

import (
	"context"
	"time"

	"github.com/codeready-toolchain/toolchain-e2e/setup/auth"
	"github.com/codeready-toolchain/toolchain-e2e/setup/metrics/queries"
	metrics "github.com/redhat-appstudio-qe/performance-toolkit/metrics/getters"
	"github.com/redhat-appstudio/e2e-tests/pkg/constants"
	"github.com/redhat-appstudio/e2e-tests/pkg/framework"
	"k8s.io/klog"
)



func StartCollection(ctx context.Context) (chan struct{}, *metrics.Gatherer) {
	framework := ctx.Value("framework").(*framework.Framework)

		token, err := auth.GetTokenFromOC()
		if err != nil {
			tokenRequestURI, err := auth.GetTokenRequestURI(framework.CommonController.KubeRest()) // authorization.FindTokenRequestURI(framework.CommonController.KubeRest())
			if err != nil {
				klog.Fatalf("a token is required to capture metrics, use oc login to log into the cluster: %v", err)
			}
			klog.Fatalf("a token is required to capture metrics, use oc login to log into the cluster. alternatively request a token and use the token flag: %v", tokenRequestURI)
		}
	
	metricsInstance := metrics.NewEmpty(framework.CommonController.KubeRest(), 5*time.Minute)

	prometheusClient := metrics.GetPrometheusClient(framework.CommonController.KubeRest(), token)

	metricsInstance.AddQueries(
		queries.QueryClusterCPUUtilisation(prometheusClient),
		queries.QueryClusterMemoryUtilisation(prometheusClient),
		queries.QueryNodeMemoryUtilisation(prometheusClient),
		queries.QueryEtcdMemoryUsage(prometheusClient),
		queries.QueryWorkloadCPUUsage(prometheusClient, constants.OLMOperatorNamespace, constants.OLMOperatorWorkload),
		queries.QueryWorkloadMemoryUsage(prometheusClient, constants.OLMOperatorNamespace, constants.OLMOperatorWorkload),
		queries.QueryOpenshiftKubeAPIMemoryUtilisation(prometheusClient),
		queries.QueryWorkloadCPUUsage(prometheusClient, constants.OSAPIServerNamespace, constants.OSAPIServerWorkload),
		queries.QueryWorkloadMemoryUsage(prometheusClient, constants.OSAPIServerNamespace, constants.OSAPIServerWorkload),
		queries.QueryWorkloadCPUUsage(prometheusClient, constants.HostOperatorNamespace, constants.HostOperatorWorkload),
		queries.QueryWorkloadMemoryUsage(prometheusClient, constants.HostOperatorNamespace, constants.HostOperatorWorkload),
		queries.QueryWorkloadCPUUsage(prometheusClient, constants.MemberOperatorNamespace, constants.MemberOperatorWorkload),
		queries.QueryWorkloadMemoryUsage(prometheusClient, constants.MemberOperatorNamespace, constants.MemberOperatorWorkload),
		//queries.QueryWorkloadCPUUsage(prometheusClient, "application-service", "application-service-controller-manager"),
		//queries.QueryWorkloadMemoryUsage(prometheusClient, "application-service", "application-service-controller-manager"),
		//queries.QueryWorkloadCPUUsage(prometheusClient, "build-service", "build-service-controller-manager"),
		//queries.QueryWorkloadMemoryUsage(prometheusClient, "build-service", "build-service-controller-manager"),
	)
	
	stopMetrics := metricsInstance.StartGathering()
	
	return stopMetrics, metricsInstance

}