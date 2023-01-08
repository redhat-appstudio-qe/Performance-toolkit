package application

import (
	"context"
	"errors"
	//"sync"
	//"time"
	"github.com/redhat-appstudio/e2e-tests/pkg/framework"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func DeletePod(ctx context.Context, namespace string, labelkey string, labelvalue string)(error){
	framework := ctx.Value("framework").(*framework.Framework)
	PodList, er := framework.CommonController.ListPods(namespace, labelkey, labelvalue,  2)
	if er != nil {
	  return errors.New("error! cannot list pods")
	}
	k := framework.CommonController.K8sClient
	for _, i := range PodList.Items{
	  if i.Namespace == namespace {
		err := k.KubeInterface().CoreV1().Pods(namespace).Delete(ctx, i.Name, metav1.DeleteOptions{})
		if err != nil {
		  return errors.New("error! cannot delete pod")
		}
	  } 
	}
	
	return  nil
}


//Injects the chaos into the cluster - total(no of times chaos needs to be injected)
/* func InjectChaos(ctx context.Context, namespace string, labelkey string, labelvalue string, total int, ch chan<- int, wg *sync.WaitGroup){
	defer wg.Done()
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Injecting Chaos no %d \n", idx)
		err := DeletePod(ctx, namespace, labelkey, labelvalue)
		if err != nil{
			errors.New("Something bad happened!!")
		}
		ch <- idx
	}

} */

func DeletePodExperiment(ctx context.Context){
	chaos_namespace := ctx.Value("chaos_ns").(string)
	podDeleteLabelKey := ctx.Value("podDeleteLabelKey").(string)
	podDeleteLabelValue := ctx.Value("podDeleteLabelValue").(string)
	err := DeletePod(ctx, chaos_namespace, podDeleteLabelKey, podDeleteLabelValue)
	if err != nil{
		errors.New("Something bad happened!!")
	}
}

//Checks after/during the chaos injects
/* func Probe(ctx context.Context, namespace string,  ch <-chan int, wg *sync.WaitGroup){
	defer wg.Done()

	for num := range ch {
		fmt.Printf("Starting Probe %d \n", num)
		applicationName := "app" + namespace
		framework := ctx.Value("framework").(*framework.Framework)
		time.Sleep(10 * time.Second)
		createdApplication, err := framework.HasController.CreateHasApplication(applicationName, namespace)
		if err != nil {
			fmt.Printf(err.Error())
		  }
		  if createdApplication.Spec.DisplayName != applicationName {
			errors.New("error! Application name doesnt match")
		  }
	}

} */

