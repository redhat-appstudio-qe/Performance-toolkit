package probes

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redhat-appstudio/e2e-tests/pkg/framework"
)


func ProbeCreateHASApplication(ctx context.Context){
	fmt.Println("Starting Probe")
	ns := ctx.Value("test-namespace").(string)
	time := ctx.Value("time").(time.Time)
	applicationName := "app" + ns + time.Format("20060102150405")
	framework := ctx.Value("framework").(*framework.Framework)
	createdApplication, err := framework.HasController.CreateHasApplication(applicationName, ns)
	if err != nil {
	 fmt.Printf(err.Error())
   }
   if createdApplication.Spec.DisplayName != applicationName {
	 errors.New("error! Application name doesnt match")
   }
 }
	// New Generated Probe
	func NodespikeProbe(ctx context.Context){
		//Add your logic here
		fmt.Println("this is a probe")
	}