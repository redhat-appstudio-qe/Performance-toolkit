package common

import (
	"context"
	"fmt"

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
	func NodespikeBefore(ctx context.Context) context.Context {
		//Add your logic here
		return ctx
	}
	
	//Generated Before
	func NodespikeAfter(ctx context.Context){
		//Add your logic here
	}
	