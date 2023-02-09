package config

import "context"

const (
	TEST_NAMESPACE string = "chaos-tests"
	CHAOS_NAMESPACE string = "application-service"
	POD_DELETE_LABEL_KEY string = "control-plane"
	POD_DELETE_LABEL_VALUE string = "controller-manager"
	USERNAME_PREFIX string = "load"
	QuarkusDevfileSource string = "https://github.com/devfile-samples/devfile-sample-code-with-quarkus"
  	NodejsDevfileSource string = "https://github.com/sawood14012/simple-nodejs-app"
	PythonDevfileSource string = "https://github.com/devfile-samples/devfile-sample-python-basic"
	DotnetDevfileSource string = "https://github.com/devfile-samples/devfile-sample-dotnet60-basic"
)

type Inject func(ctx context.Context)
type Probe func(ctx context.Context)
type Before func(ctx context.Context)(context.Context)
type After func(ctx context.Context)

type Expirement struct{
        
    // defining struct variables
    Name string
    Probe Probe
    Inject Inject
	Before Before
	After After
	ChaosIteration int
	ProbeIntervalSecs int
	Weight int
}
