package utils

import "fmt"

const (
	DEFAULT_CHAOS_ITERATION int = 2
	DEFAULT_PROBE_INTERVAL_SECS int = 10
	DEFAULT_WEIGHT int = 10
)

func GetExperimentTemplate(Package string, MainFunction string) string {
	temp := `package %s

	import (
		"context"
		"errors"
		"fmt"
	)
	
	//MainChaosFunction
	//This Function executes/injects chaos into the system
	func %[2]v(ctx context.Context) (error) {
		fmt.Println("replace with your logic")
		// Add your code here
		return nil
	}
	
	//RunnerFunction
	//This is the Runner Function which runs the above functions in a collabrative way
	func %[2]vExperiment(ctx context.Context){
		err := %[2]v(ctx)
		if err != nil{
			errors.New("Something bad happened!!")
		}
	}`
	result := fmt.Sprintf(temp, Package, MainFunction)

	return result
}

func GetProbeTemplate(ProbeName string) string {
	temp := `
	// New Generated Probe
	func %sProbe(ctx context.Context){
		//Add your logic here
		fmt.Println("this is a probe")
	}`
	result := fmt.Sprintf(temp, ProbeName)

	return result
}

func GetBeforeTemplate(Name string) string {
	temp := `
	//Generated Before
	func %sBefore(ctx context.Context) context.Context {
		//Add your logic here
		return ctx
	}
	`
	result := fmt.Sprintf(temp, Name)

	return result
}

func GetAfterTemplate(Name string) string {
	temp := `
	//Generated Before
	func %sAfter(ctx context.Context){
		//Add your logic here
	}
	`
	result := fmt.Sprintf(temp, Name)

	return result
}

func GetAppendFeatureTemplate(Name string, module string, Expname string) string {
	temp := `
	//Generated New
	{
		Name: "%[6]v", 
		Probe: probes.%[6]vProbe,
		Inject: %[2]v.%[1]vExperiment,
		ChaosIteration: %[3]v,
		ProbeIntervalSecs: %[4]v,
		Before: common.%[6]vBefore,
		After: common.%[6]vAfter,
		Weight: %[5]v,
	},
}`
result := fmt.Sprintf(temp, Name, module, DEFAULT_CHAOS_ITERATION, DEFAULT_PROBE_INTERVAL_SECS, DEFAULT_WEIGHT, Expname)
return result
}

