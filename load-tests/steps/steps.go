package steps

import (
	"context"
	"errors"
	"fmt"

	// "log"
	"time"

	"github.com/codeready-toolchain/toolchain-e2e/setup/users"
	"github.com/codeready-toolchain/toolchain-e2e/setup/wait"
	controller "github.com/redhat-appstudio-qe/concurency-controller/controller"
	config "github.com/redhat-appstudio-qe/performance-toolkit/config"
	metrics "github.com/redhat-appstudio-qe/performance-toolkit/metrics/getters"
	localutils "github.com/redhat-appstudio-qe/performance-toolkit/utils"
	appservice "github.com/redhat-appstudio/application-api/api/v1alpha1"
	"github.com/redhat-appstudio/e2e-tests/pkg/constants"
	"github.com/redhat-appstudio/e2e-tests/pkg/framework"
	"github.com/redhat-appstudio/e2e-tests/pkg/utils"
	"k8s.io/klog"
)

type godogsCtxKey struct{}

type Concurent struct{
	Framework *framework.Framework
	GlobalCtx context.Context
	Controller  *controller.LoadController
}

type InnerTestVars struct {
	component string
	user string

}

func NewConcurent(Framework *framework.Framework, GlobalCtx context.Context,Controller  *controller.LoadController) (*Concurent) {
	return &Concurent{Framework: Framework, GlobalCtx: GlobalCtx, Controller: Controller}
}

var (
	AverageUserCreationTime time.Duration
	ConcurentCtx *Concurent
)

func SimpleTest() error {
	ctx := ConcurentCtx.GlobalCtx
	user := localutils.RandomString(config.USERNAME_PREFIX)
	framework := ctx.Value("framework").(*framework.Framework)
	err := CreateConcurentUser(framework, user)
	if err != nil{
		return err
	}
	return nil
}

func ConcurrentTestRunner() error {
	ctx := ConcurentCtx.GlobalCtx
	user := localutils.RandomString(config.USERNAME_PREFIX)
	framework := ctx.Value("framework").(*framework.Framework)
	err := CreateConcurentUser(framework, user)
	if err != nil{
		return err
	}
	ctx, err = CreateAppstudioApp(ctx, framework, user)
	if err != nil {
		return err
	}
	devfile := ctx.Value("devfile").(string)
	ctx, err = CreateAppstudioComponents(ctx, framework, user, devfile)
	if err != nil {
		return err
	}
	return nil
}

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

func ConfigureBatchConcurentTests(ctx context.Context, MaxReq int, Batches int){
	URL := localutils.CheckVarExistsAndReturn("MONITORING_URL", false)
	c := controller.NewLoadController(MaxReq, Batches, URL)
	framework := ctx.Value("framework").(*framework.Framework)
	context := ctx
	ConcurentCtx = NewConcurent(framework, context, c)
	
}

func ConfigureInfiniteConcurentTests(ctx context.Context, RPS int, TimeoutSecs int){
	URL := localutils.CheckVarExistsAndReturn("MONITORING_URL", false)
	Timeout := localutils.ToDuration(TimeoutSecs)
	c := controller.NewInfiniteLoadController(Timeout, RPS, URL)
	framework := ctx.Value("framework").(*framework.Framework)
	context := ctx
	ConcurentCtx = NewConcurent(framework, context, c)
	
}
func ConfigureSpikeConcurentTests(ctx context.Context, RPS int, TimeoutSecs int){
	URL := localutils.CheckVarExistsAndReturn("MONITORING_URL", false)
	Timeout := localutils.ToDuration(TimeoutSecs)
	c := controller.NewSpikeLoadController(Timeout, RPS , 0.5, URL)
	framework := ctx.Value("framework").(*framework.Framework)
	context := ctx
	ConcurentCtx = NewConcurent(framework, context, c)
	
}

func StartBatchConcurentTests(ctx context.Context, cmp string){
	c := ConcurentCtx.Controller
	ConcurentCtx.GlobalCtx = context.WithValue(ConcurentCtx.GlobalCtx, "devfile", cmp)
	c.ConcurentlyExecute(SimpleTest)
}

func StartInfiniteConcurentTests(ctx context.Context, cmp string){
	c := ConcurentCtx.Controller
	ConcurentCtx.GlobalCtx = context.WithValue(ConcurentCtx.GlobalCtx, "devfile", cmp)
	c.ConcurentlyExecuteInfinite(ConcurrentTestRunner)
}

func StartSpikeConcurentTests(ctx context.Context, cmp string){
	c := ConcurentCtx.Controller
	ConcurentCtx.GlobalCtx = context.WithValue(ConcurentCtx.GlobalCtx, "devfile", cmp)
	c.CuncurentlyExecuteSpike(SimpleTest)
}

func CreateConcurentUser(framework *framework.Framework, user string) error {
	startTime := time.Now()
	err := users.Create(framework.CommonController.KubeRest(), user, 
		constants.HostOperatorNamespace, constants.MemberOperatorNamespace)
	if err != nil {
		return err
	}
	if err := wait.ForNamespace(framework.CommonController.KubeRest(), user); err != nil {
		klog.Fatalf("failed to find namespace '%s'", user)
		return err
	}
	UserCreationTime := time.Since(startTime)
	klog.Infof("User Created in:", UserCreationTime)
	return nil
}

func CreateAppstudioApp(Ctx context.Context, framework *framework.Framework, user string) (context.Context, error) {
	ApplicationName := fmt.Sprintf("%s-app", user)
	_, errors := framework.CommonController.CreateRegistryAuthSecret(
		"redhat-appstudio-registry-pull-secret",
		user,
		utils.GetDockerConfigJson(),
	)
	if errors != nil {
		klog.Infof("Problem Creating the secret: %v", errors)
		return Ctx, errors
	}
	app, err := framework.HasController.CreateHasApplication(ApplicationName, user)
	if err != nil {
		klog.Info("Problem Creating the Application: %v", err)
		return Ctx, err
	}
	if err := utils.WaitUntil(framework.CommonController.ApplicationGitopsRepoExists(app.Status.Devfile), 30*time.Second); err != nil {
		klog.Infof("timed out waiting for application gitops repo to be created: %v", err)
		return Ctx, err
	}
	
	return context.WithValue(Ctx, "app", app.Name), nil
}

func CreateAppstudioComponents(Ctx context.Context, framework *framework.Framework, user string, component string) (context.Context, error) {
	gitSource := parseDevfileSource(component)
	ComponentName := fmt.Sprintf("%s-component", user)
	app := Ctx.Value("app").(string)
	_, err := framework.HasController.CreateComponentDetectionQuery(ComponentName, user, gitSource, "", false)
	if err != nil {
		klog.Infof("error: %v", err)
		return Ctx, err
	}
	time.Sleep(3 * time.Second)
	compDetected := appservice.ComponentDetectionDescription{}
	cdqD, err := framework.HasController.GetComponentDetectionQuery(ComponentName, user)
	if err != nil {
		klog.Infof("error! cannot get cdq")
		return Ctx, err
	}
	for _, compDetected = range cdqD.Status.ComponentDetected {
		if !compDetected.DevfileFound {
			klog.Infof("error! Devfile not found")
			return Ctx, err
		}
		
	}
	cmp, errs := framework.HasController.CreateComponentFromStub(compDetected, ComponentName, user, "", app)
	if errs != nil {
		klog.Infof("error! cannot create component %v", err)
		return Ctx, errs
	}

	CreatedCmp, err := framework.HasController.GetHasComponent(cmp.Name, user)
	if err != nil {
		klog.Fatalf("error! cannot get component %v", err)
		return Ctx, err
	}

	return context.WithValue(Ctx, "component", CreatedCmp.Name), nil
	
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



