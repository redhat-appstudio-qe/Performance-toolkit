package probes

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redhat-appstudio/e2e-tests/pkg/framework"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	func NodeDrainProbe(ctx context.Context){
		//Add your logic here
		fmt.Println("Starting Node Drain Probe!")
		TargetNode := ctx.Value("NodeToBeDrained").(string)
		framework := ctx.Value("framework").(*framework.Framework)
		k := framework.CommonController.K8sClient
		nodeSpec, err := k.KubeInterface().CoreV1().Nodes().Get(ctx, TargetNode, v1.GetOptions{})
		if err != nil{
			log.Fatal(err)
		}
		if !nodeSpec.Spec.Unschedulable{
			fmt.Println("Node is not Unschedulable!")
			fmt.Println("ProbeSuccessful")
		}
	}
	// New Generated Probe
	func NetworkDenyProbe(ctx context.Context){
		//Add your logic here
		//ProbeCreateHASApplication(ctx)
	}