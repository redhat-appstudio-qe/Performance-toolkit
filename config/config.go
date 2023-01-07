package config

import "context"

const (
	TEST_NAMESPACE string = "chaos-tests"
	CHAOS_NAMESPACE string = "application-service"
	POD_DELETE_LABEL_KEY string = "control-plane"
	POD_DELETE_LABEL_VALUE string = "controller-manager"
)

type Inject func(ctx context.Context)
type Probe func(ctx context.Context)

type Expirement struct{
        
    // defining struct variables
    Name string
    Probe Probe
    Inject Inject
	ChaosIteration int
	ProbeIntervalSecs int
}
