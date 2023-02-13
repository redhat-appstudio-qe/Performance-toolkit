
# Stone Soup Load Tests 

Load Tests utilizes a combination of goDog framework and the [concurency-controller](https://github.com/redhat-appstudio-qe/concurency-controller) module to run different type of concurent tests on stone soup 

### Directory Structure 
```
├── controller
│ └── controller.go
├── features/
│ └── load.feature
├── steps/
│ └── steps.go
├── clear.sh
├── load_test.go
├── load-template.env
├── run.sh
└── stuck.sh
```


#### Controller 

Contains all the control functions hooks and step definations that goDog utilizes to run the load tests 

#### features 

this folder contains all the feature files of gherkin files what goDog runs , this folder can contain multiple feature files containing multiple scenarios and scenario outlines to execute the load tests 

#### Steps 

This is the directory you can utilize to store your function implimentations that together combined will work as a scenario also note that steps are reusable throught the test/scenarios

To understand more please look at [godog](https://github.com/cucumber/godog)

## Run Load Tests 

#### Requirements 
- StoneSoup/Appstudio installed on a cluster using preview mode with sandbox in end to end mode  
- Should be logged in to the cluster using oc 
- (Optional) Deploy perfromance monitoring stack and get the Route for ingester instance [README](https://github.com/redhat-appstudio-qe/perf-monitoring#deploy-on-openshift)




```bash
  git clone https://github.com/redhat-appstudio-qe/performance-toolkit.git 
```

Go to the load-tests directory

```bash
  cd load-tests
```

Create an `.env` file using the template `load-template.env` 

```bash
  cp load-template.env .env
```
- And fill up all the required variables

To start the test run 

```bash
  ./run.sh
```

#### Clean up 

To clear users created 

```bash 
make -C /tmp/toolchain-e2e clean-users
```

To Clear the Repositories created in github org 

```bash 
GITHUB_TOKEN={TOKEN} ./clear.sh
```
