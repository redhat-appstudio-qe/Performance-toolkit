package steps

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/codeready-toolchain/toolchain-e2e/setup/users"
	"github.com/codeready-toolchain/toolchain-e2e/setup/wait"
	"github.com/google/uuid"
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
	appName string
}

type BatchConcurent struct {
	Concurent
	Controller *controller.BatchLoadType
}

type InfiniteConcurent struct {
	Concurent
	Controller *controller.InfiniteLoadType
}
type SpikeConcurent struct {
	Concurent
	Controller *controller.SpikeLoadType
}

func NewBatchConcurent(Framework *framework.Framework, GlobalCtx context.Context, controller *controller.BatchLoadType) (*BatchConcurent) {
	Concurrent := Concurent{Framework: Framework, GlobalCtx: GlobalCtx}
	return &BatchConcurent{Concurent: Concurrent, Controller: controller}
}
func NewInfiniteConcurent(Framework *framework.Framework, GlobalCtx context.Context, controller *controller.InfiniteLoadType) (*InfiniteConcurent) {
	Concurrent := Concurent{Framework: Framework, GlobalCtx: GlobalCtx}
	return &InfiniteConcurent{Concurent: Concurrent, Controller: controller}
}
func NewSpikeConcurent(Framework *framework.Framework, GlobalCtx context.Context, controller *controller.SpikeLoadType) (*SpikeConcurent) {
	Concurrent := Concurent{Framework: Framework, GlobalCtx: GlobalCtx}
	return &SpikeConcurent{Concurent: Concurrent, Controller: controller}
}

var (
	AverageUserCreationTime time.Duration
	ConcurentBatchCtx *BatchConcurent
	ConcurentInfiniteCtx *InfiniteConcurent
	ConcurentSpikeCtx *SpikeConcurent
)


func ConcurrentBatchTestRunner() error {
	user := localutils.RandomString(config.USERNAME_PREFIX)
	framework := ConcurentBatchCtx.Framework
	devfile := ConcurentBatchCtx.GlobalCtx.Value("devfile").(string)
	return ConcurentTest(framework, user, devfile)
}

func ConcurrentInfiniteTestRunner() error {
	user := localutils.RandomString(config.USERNAME_PREFIX)
	framework := ConcurentInfiniteCtx.Framework
	devfile := ConcurentInfiniteCtx.GlobalCtx.Value("devfile").(string)
	return ConcurentTest(framework, user, devfile)
}

func ConcurrentSpikeTestRunner() error {
	user := localutils.RandomString(config.USERNAME_PREFIX)
	framework := ConcurentSpikeCtx.Framework
	devfile := ConcurentSpikeCtx.GlobalCtx.Value("devfile").(string)
	return ConcurentTest(framework, user, devfile)
}

func ConcurentTest(framework *framework.Framework, user string, devfile string) error {
	namespace := user + "-tenant"
	err := CreateConcurentUser(framework, user, namespace)
	if err != nil {
		return err
	}
	appName, err := CreateAppstudioApp(framework, namespace)
	if err != nil {
		return err
	}
	_, err = CreateEnvironments(namespace, framework)
	if err != nil {
		return err
	}
	componentName, err := CreateAppstudioComponent(appName, framework, namespace, devfile)
	if err != nil {
		return err
	}
	klog.Infof("Component Created Successfully :", componentName)
	return nil
}

func IsPhysicalSystemRunning(ctx context.Context) (context.Context, error) {
  return ctx,nil
}


func ConfigureBatchConcurentTests(ctx context.Context, MaxReq int, Batches int){
	URL := localutils.CheckVarExistsAndReturn("MONITORING_URL", false)
	c := controller.NewBatchController(MaxReq, Batches, URL)
	framework := ctx.Value("framework").(*framework.Framework)
	context := ctx
	ConcurentBatchCtx = NewBatchConcurent(framework, context, c)
	
}

func ConfigureInfiniteConcurentTests(ctx context.Context, RPS int, TimeoutSecs int){
	URL := localutils.CheckVarExistsAndReturn("MONITORING_URL", false)
	Timeout := localutils.ToDuration(TimeoutSecs)
	c := controller.NewInfiniteController(RPS,Timeout, URL)
	framework := ctx.Value("framework").(*framework.Framework)
	context := ctx
	
	ConcurentInfiniteCtx = NewInfiniteConcurent(framework, context, c)
	
}
func ConfigureSpikeConcurentTests(ctx context.Context, RPS int, TimeoutSecs int){
	URL := localutils.CheckVarExistsAndReturn("MONITORING_URL", false)
	Timeout := localutils.ToDuration(TimeoutSecs)
	c := controller.NewSpikeController(RPS , Timeout,0.5, URL)
	framework := ctx.Value("framework").(*framework.Framework)
	context := ctx
	ConcurentSpikeCtx = NewSpikeConcurent(framework, context, c)
	
}

func StartBatchConcurentTests(ctx context.Context, cmp string){
	c := ConcurentBatchCtx.Controller
	ConcurentBatchCtx.GlobalCtx = context.WithValue(ConcurentBatchCtx.GlobalCtx, "devfile", cmp)
	c.ConcurentlyExecute(ConcurrentBatchTestRunner)
}

func StartInfiniteConcurentTests(ctx context.Context, cmp string){
	c := ConcurentInfiniteCtx.Controller
	ConcurentInfiniteCtx.GlobalCtx = context.WithValue(ConcurentInfiniteCtx.GlobalCtx, "devfile", cmp)
	c.ConcurentlyExecuteInfinite(ConcurrentInfiniteTestRunner)
}

func StartSpikeConcurentTests(ctx context.Context, cmp string){
	c := ConcurentSpikeCtx.Controller
	ConcurentSpikeCtx.GlobalCtx = context.WithValue(ConcurentSpikeCtx.GlobalCtx, "devfile", cmp)
	c.CuncurentlyExecuteSpike(ConcurrentSpikeTestRunner)
}

func CreateConcurentUser(framework *framework.Framework, user string, namespace string) (error) {
	err := users.Create(framework.CommonController.KubeRest(), user, 
		constants.HostOperatorNamespace, constants.MemberOperatorNamespace)
	if err != nil {
		return  err
	}
	if err := wait.ForNamespace(framework.CommonController.KubeRest(), namespace); err != nil {
		return err
	}
	return  nil
}

func CreateAppstudioApp(framework *framework.Framework, namespace string) (string, error) {
	ApplicationName := fmt.Sprintf("%s-app", namespace)
	_, errors := framework.CommonController.CreateRegistryAuthSecret(
		"redhat-appstudio-registry-pull-secret",
		namespace,
		utils.GetDockerConfigJson(),
	)
	if errors != nil {
		klog.Infof("Problem Creating the secret: %v", errors)
		return "", errors
	}
	app, err := framework.HasController.CreateHasApplication(ApplicationName, namespace)
	if err != nil {
		klog.Info("Problem Creating the Application: %v", err)
		return "", err
	}
	if err := utils.WaitUntil(framework.CommonController.ApplicationGitopsRepoExists(app.Status.Devfile), 30*time.Second); err != nil {
		klog.Infof("timed out waiting for application gitops repo to be created: %v", err)
		return "", err
	}
	
	return app.Name , nil
}

func CreateEnvironments(namespace string, framework *framework.Framework) (string, error) {
	env, err := framework.IntegrationController.CreateEnvironment(namespace)
	if err != nil{
		return "", err
	}
	return env.Name, nil
}

func CreateAppstudioComponent(app string, framework *framework.Framework, namespace string, component string)(string, error){
	gitSource := parseDevfileSource(component)
	var containerIMG = fmt.Sprintf("quay.io/%s/test-images:%s", utils.GetQuayIOOrganization(), strings.Replace(uuid.New().String(), "-", "", -1))
	ComponentName := fmt.Sprintf("%s-component", namespace)
	cmp, err := framework.HasController.CreateComponent(app, ComponentName, namespace, gitSource, "", "", containerIMG, "", true)
	if err != nil {
		klog.Infof("error: %v", err)
		return "", err
	}
	return cmp.Name, nil
}

func CreateAppstudioComponentsSTUB(app string, framework *framework.Framework, namespace string, component string) (string, error) {
	gitSource := parseDevfileSource(component)
	ComponentName := fmt.Sprintf("%s-component", namespace)
	var containerIMG = fmt.Sprintf("quay.io/%s/test-images:%s", utils.GetQuayIOOrganization(), strings.Replace(uuid.New().String(), "-", "", -1))
	framework.HasController.CreateComponent(app, ComponentName, namespace, gitSource, "", "", containerIMG, "", true)
	_, err := framework.HasController.CreateComponentDetectionQuery(ComponentName,namespace, gitSource, "", false)
	if err != nil {
		klog.Infof("error: %v", err)
		return "", err
	}
	time.Sleep(3 * time.Second)
	compDetected := appservice.ComponentDetectionDescription{}
	cdqD, err := framework.HasController.GetComponentDetectionQuery(ComponentName, namespace)
	if err != nil {
		klog.Infof("error! cannot get cdq")
		return "", err
	}
	for _, compDetected = range cdqD.Status.ComponentDetected {
		if !compDetected.DevfileFound {
			klog.Infof("error! Devfile not found")
			return "", err
		}
		
	}
	cmp, errs := framework.HasController.CreateComponentFromStub(compDetected, ComponentName, namespace, "", app)
	if errs != nil {
		klog.Infof("error! cannot create component %v", err)
		return "", errs
	}

	CreatedCmp, err := framework.HasController.GetHasComponent(cmp.Name, namespace)
	if err != nil {
		klog.Infof("error! cannot get component %v", err)
		return "", err
	}

	return CreatedCmp.Name, nil
	
}

func PrintMetrics(ctx context.Context){
	closeMetrics := ctx.Value("closeMetrics").(chan struct{})
	defer close(closeMetrics)
	metricsInstance := ctx.Value("metricsInstance").(*metrics.Gatherer)
	metricsInstance.PrintResults()
}

func parseDevfileSource(DevfileSource string) string {
  if DevfileSource == "Quarkus"{
    return config.QuarkusDevfileSource
  } else if DevfileSource == "python"{
	return config.PythonDevfileSource
  } else if DevfileSource == "dotnet"{
	return config.DotnetDevfileSource
  }
  return config.NodejsDevfileSource
}






