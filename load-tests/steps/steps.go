package steps

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/codeready-toolchain/toolchain-e2e/setup/users"
	"github.com/codeready-toolchain/toolchain-e2e/setup/wait"
	config "github.com/redhat-appstudio-qe/performance-toolkit/config"
	metrics "github.com/redhat-appstudio-qe/performance-toolkit/metrics/getters"
	appservice "github.com/redhat-appstudio/application-api/api/v1alpha1"
	"github.com/redhat-appstudio/e2e-tests/pkg/constants"
	"github.com/redhat-appstudio/e2e-tests/pkg/framework"
	"github.com/redhat-appstudio/e2e-tests/pkg/utils"
	"k8s.io/klog"
)

type godogsCtxKey struct{}

var (
	AverageUserCreationTime time.Duration
)

func IsPhysicalSystemRunning(ctx context.Context) (context.Context, error) {
  return ctx,nil
}

func WaitforMins(ctx context.Context, min int64) (context.Context, error){
  time.Sleep(time.Duration(min * int64(time.Minute)))
  return ctx, nil
}

func WaitforSecs(ctx context.Context, secs int64) (context.Context, error){
  time.Sleep(time.Duration(secs * int64(time.Second)))
  return ctx, nil
}


func CreateNS(ctx context.Context, namespace string) (context.Context, error){
  framework := ctx.Value("framework").(*framework.Framework)
  creatednamespace, err := framework.CommonController.CreateTestNamespace(namespace)
  if err != nil {
    return ctx, errors.New("cannot create ns")
  }
  return context.WithValue(ctx, "ns", creatednamespace.Name), nil
}

func CreateResourcesConcurently()(error){
	return nil
}


func CreateUsers(ctx context.Context, number int)(context.Context){
	startTime := time.Now()
	ctx = context.WithValue(ctx, "NumberOfUsers", number)
	framework := ctx.Value("framework").(*framework.Framework)
	userprefix := config.USERNAME_PREFIX
	for i:=1;i<=number;i++{
		username := fmt.Sprintf("%s-%04d", userprefix, i)
		if err := users.Create(framework.CommonController.KubeRest(), username, constants.HostOperatorNamespace, constants.MemberOperatorNamespace); err != nil {
			klog.Fatalf("failed to provision user '%s'", username)
			klog.Errorf(err.Error())
		}
		if err := wait.ForNamespace(framework.CommonController.KubeRest(), username); err != nil {
			klog.Fatalf("failed to find namespace '%s'", username)
			klog.Errorf(err.Error())
		}
		UserCreationTime := time.Since(startTime)
		AverageUserCreationTime += UserCreationTime
	}
	ctx = context.WithValue(ctx, "UserCreationTime", AverageUserCreationTime)

	return ctx

	
}

func CreateResources(ctx context.Context, cmp string)(context.Context){
	framework := ctx.Value("framework").(*framework.Framework)
	numberUsers := ctx.Value("NumberOfUsers").(int)
	userprefix := config.USERNAME_PREFIX
	for i:=1;i<=numberUsers;i++{
		username := fmt.Sprintf("%s-%04d", userprefix, i)
		ApplicationName := fmt.Sprintf("%s-app", username)
		_, errors := framework.CommonController.CreateRegistryAuthSecret(
			"redhat-appstudio-registry-pull-secret",
			username,
			utils.GetDockerConfigJson(),
		)
		if errors != nil {
			klog.Fatalf("Problem Creating the secret: %v", errors)
		}
		app, err := framework.HasController.CreateHasApplication(ApplicationName, username)
			if err != nil {
				klog.Fatalf("Problem Creating the Application: %v", err)
			}
			if err := utils.WaitUntil(framework.CommonController.ApplicationGitopsRepoExists(app.Status.Devfile), 30*time.Second); err != nil {
				klog.Fatalf("timed out waiting for application gitops repo to be created: %v", err)
			}
			d := parseDevfileSource(cmp)
		
		componentName := fmt.Sprintf("%s-component-%d", userprefix,i)
		cdq, err := framework.HasController.CreateComponentDetectionQuery(componentName, username, d, "", false)
		if err != nil {
			klog.Fatalf("error: %v", err)
			return ctx
		}

		ctx = context.WithValue(ctx, "cdq", cdq)
		time.Sleep(5 * time.Second)
		compDetected := appservice.ComponentDetectionDescription{}
		cdqD, err := framework.HasController.GetComponentDetectionQuery(componentName, username)
		if err != nil {
			klog.Fatalf("error! cannot get cdq")
			return ctx
		}
		for _, compDetected = range cdqD.Status.ComponentDetected {
			if !compDetected.DevfileFound {
				klog.Fatalf("error! Devfile not found")
				return ctx
			}
			
		}
		
		ctx = context.WithValue(ctx, "compD", compDetected)
		ctx = context.WithValue(ctx, "cdq", cdqD)
		component, err := framework.HasController.CreateComponentFromStub(compDetected, componentName, username, "", app.Name)
		if err != nil {
			klog.Fatalf("error! cannot create component %v", err)
			return ctx
		}

		componet, err := framework.HasController.GetHasComponent(component.Name, username)
		if err != nil {
			klog.Fatalf("error! cannot get component %v", err)
			return ctx
		}
		ctx = context.WithValue(ctx, "component", componet)
	//time.Sleep(5 * time.Second)
	}
	return ctx
}

func PrintMetrics(ctx context.Context){
	closeMetrics := ctx.Value("closeMetrics").(chan struct{})
	defer close(closeMetrics)
	metricsInstance := ctx.Value("metricsInstance").(*metrics.Gatherer)
	metricsInstance.PrintResults()
}

func CreateHasApp(ctx context.Context, applicationName string) (context.Context, error){
  framework := ctx.Value("framework").(*framework.Framework)
  ns := ctx.Value("ns").(string)
  createdApplication, err := framework.HasController.CreateHasApplication(applicationName, ns)
	if err != nil {
    fmt.Printf(err.Error())
    return ctx, errors.New("error! cannot create has application")
  }
  if createdApplication.Spec.DisplayName != applicationName {
    return ctx, errors.New("error! Application name doesnt match")
  }

  return ctx, nil
}

func CheckHasApp(ctx context.Context, applicationName string) (context.Context, error){
  framework := ctx.Value("framework").(*framework.Framework)
  ns := ctx.Value("ns").(string)
  application, err := framework.HasController.GetHasApplication(applicationName, ns)
  if err != nil {
    return ctx, errors.New("error! cannot get the application")
  }
  ctx = context.WithValue(ctx, "application", application)

  return ctx, nil
}

func parseDevfileSource(DevfileSource string) string {
  QuarkusDevfileSource := "https://github.com/redhat-appstudio-qe/devfile-sample-code-with-quarkus"
  NodejsDevfileSource := "https://github.com/sawood14012/simple-nodejs-app"
  if DevfileSource != "Quarkus"{
    return NodejsDevfileSource
  }
  return QuarkusDevfileSource

}
func CreateCompoenetDetectionQuery(ctx context.Context, componentName string, DevfileSource string) (context.Context, error){
  framework := ctx.Value("framework").(*framework.Framework)
  ns := ctx.Value("ns").(string)
  d := parseDevfileSource(DevfileSource)
  cdq, err := framework.HasController.CreateComponentDetectionQuery(componentName, ns, d, "", false)
  if err != nil {
    return ctx, errors.New("error! cannot create component detection query")
  }
  return context.WithValue(ctx, "cdq", cdq), nil

}

func CheckComponentDetectionQuery(ctx context.Context, componentName string) (context.Context, error){
  framework := ctx.Value("framework").(*framework.Framework)
  ns := ctx.Value("ns").(string)
  compDetected := appservice.ComponentDetectionDescription{}
  cdq, err := framework.HasController.GetComponentDetectionQuery(componentName, ns)
  if err != nil {
    return ctx, errors.New("error! cannot get cdq")
  }
  for _, compDetected = range cdq.Status.ComponentDetected {
    if !compDetected.DevfileFound {
      return ctx, errors.New("devfile not found in cdq")
    }
    
  }
  ctx = context.WithValue(ctx, "compD", compDetected)
  return context.WithValue(ctx, "cdq", cdq), nil
}

func CreateComponent(ctx context.Context, componentName string) (context.Context, error){
  framework := ctx.Value("framework").(*framework.Framework)
  ns := ctx.Value("ns").(string)
  application := ctx.Value("application").(*appservice.Application)
  compDetected := ctx.Value("compD").(appservice.ComponentDetectionDescription)
  fmt.Println("cmp", componentName)
  fmt.Println("cmpd", compDetected)
  component, err := framework.HasController.CreateComponentFromStub(compDetected, componentName, ns, "", application.Name)
  if err != nil {
    fmt.Printf(err.Error())
    return ctx, errors.New("error! cannot create component")
  }
  return context.WithValue(ctx, "component", component), nil
}

func CheckComponent(ctx context.Context, componentName string)(context.Context, error){
  framework := ctx.Value("framework").(*framework.Framework)
  ns := ctx.Value("ns").(string)
  componet, err := framework.HasController.GetHasComponent(componentName, ns)
  if err != nil {
    return ctx, errors.New("error! cannot get has component")
  }
  return context.WithValue(ctx, "component", componet), nil
}



