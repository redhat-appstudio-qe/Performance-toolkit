package common

import (
	"context"
	"fmt"
	//"math/rand"
	"os/exec"
	"strings"

	"github.com/redhat-appstudio-qe/performance-toolkit/config"
	"github.com/redhat-appstudio-qe/performance-toolkit/utils"
)

func ApplicationServicePodDeleteBefore(ctx context.Context) context.Context {
	fmt.Println("Setting Required Variables")
	ctx = utils.SetChaosNamespace(config.CHAOS_NAMESPACE, ctx)
	ctx = utils.SetpodDeleteLabelKey(config.POD_DELETE_LABEL_KEY, ctx)
	ctx = utils.SetpodDeleteLabelValue(config.POD_DELETE_LABEL_VALUE, ctx)
	return ctx

}


func ApplicationServicePodDeleteAfter(ctx context.Context){
	fmt.Println("i am in after ")
}
	
	//Generated Before
	func NodeDrainBefore(ctx context.Context) context.Context {
		//Add your logic here
		//kubectl get no -o json | jq -r '[.items[] | {name:.metadata.name}]' >> temp.json 
		fmt.Println("Running Before Block")
		cmd := exec.Command("kubectl", "get", "no", "-o=name", "--selector=!node-role.kubernetes.io/master")
    	stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		// Print the output
		temp := strings.Split(string(stdout), "\n")
		/* for i:=0;i<len(temp);i++{
			println(temp[i])
		} */
		//nodeno := rand.Intn(len(temp)-1) + 1
		ctx = context.WithValue(ctx, "NodeList", temp)
		TargetNode := temp[2]
		TargetNode = TargetNode[5:]
		fmt.Println("here", TargetNode)
		ctx = context.WithValue(ctx, "NodeToBeDrained", TargetNode)
		return ctx
	}
	
	//Generated Before
	func NodeDrainAfter(ctx context.Context){
		//Add your logic 
		fmt.Println("Running After Block")
		TargetNode := ctx.Value("NodeToBeDrained").(string)
		cmd := exec.Command("kubectl", "uncordon", TargetNode)
    	stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(stdout))
	}
	
	//Generated Before
	func NetworkDenyBefore(ctx context.Context) context.Context {
		//Add your logic here
		fmt.Println("Setting Before Block - Network Deny")
		ctx = utils.SetChaosNamespace(config.CHAOS_NAMESPACE, ctx)
		ctx = utils.SetpodDeleteLabelKey(config.POD_DELETE_LABEL_KEY, ctx)
		ctx = utils.SetpodDeleteLabelValue(config.POD_DELETE_LABEL_VALUE, ctx)
		ctx = context.WithValue(ctx, "PolicyFile", "templates/network-deny.yaml")
		ctx = context.WithValue(ctx, "NetworkPolicyName", "web-deny-all")
		return ctx
	}
	
	//Generated Before
	func NetworkDenyAfter(ctx context.Context){
		fmt.Println("Running After Block - Network Deny")
		TargetPolicy := ctx.Value("NetworkPolicyName").(string)
		ns := ctx.Value("chaos_ns").(string)
		cmd := exec.Command("kubectl", "delete", "networkpolicy", TargetPolicy, "-n", ns)
    	stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(stdout))
	}
	