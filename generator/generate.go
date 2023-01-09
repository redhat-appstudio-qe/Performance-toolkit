package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	utils "github.com/redhat-appstudio-qe/performance-toolkit/utils"
)

const (
	PWD string = "/Users/sawood/work/perf-tests"
)
func SelectedType(d int) string {
	switch {
		case d == 1:
			return "application"
		case d == 2:
			return "infrastructure"
		case d == 3:
			return "network"
		default: 
			return "Invalid"
}
}

func OpenFileForWrite(name string, data string){
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.WriteString(data); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}

func ReplaceFileContents(name string, data string){
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
	log.Fatal(err)
    }
}

func readfiletoappend(Name string, module string, Expname string){
    fileName := "expirements/featurelist.go"
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    fmt.Printf("\nFile Name: %s", fileName)
    fmt.Printf("\nSize: %d bytes", len(data))
    //fmt.Printf("\nData: %s", data)
	Beforechange := strings.TrimSuffix(string(data), "}")
	ReplaceFileContents(fileName, Beforechange)
	AfterChange := utils.GetAppendFeatureTemplate(Name, module, Expname)
	OpenFileForWrite(fileName, AfterChange)
	fmt.Println(".......................................")
	fmt.Printf("Added Expirement to feature list")
	fmt.Println(".......................................")

}

func createHelperFunctions(Name string){
	probe := utils.GetProbeTemplate(Name)
	ProbePath := "probes/probes.go"
	OpenFileForWrite(ProbePath, probe)
	before := utils.GetBeforeTemplate(Name)
	beforeAfterPath := "common/common.go"
	OpenFileForWrite(beforeAfterPath, before)
	fmt.Println("waiting for the details to populate...")
	time.Sleep(3 * time.Second)
	after := utils.GetAfterTemplate(Name)
	OpenFileForWrite(beforeAfterPath, after)
	fmt.Println("Done!")
	fmt.Println(".......................................")
	fmt.Println("Generated Probe Located in :", ProbePath)
	fmt.Println("Generated Before After Functions Located in :", beforeAfterPath)
	fmt.Println(".......................................")

}

func getPresentExp(module string) []fs.FileInfo {
	files, err := ioutil.ReadDir("expirements/"+module)
    if err != nil {
        log.Fatal(err)
    }
	return files
}


func main(){

	var name string = "";
	var option int = 0;
	var Etype int = 0;
	var OldExp int = 0;

	fmt.Print("Enter Expirement Name(Avoid Spaces or special chars): ")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter Details of Expirement:", name)
	fmt.Println("[1] This is a Fresh Expirement")
	fmt.Println("[2] New Scenario Based on Expirents Present")
	fmt.Print("Choose an option:")
	fmt.Scanf("%d", &option)
	fmt.Println("[1] Application Based")
	fmt.Println("[2] Infrastructure Based")
	fmt.Println("[3] Network Based")
	fmt.Print("Choose an option:")
	fmt.Scanf("%d", &Etype)
	subPath := "expirements/" + SelectedType(Etype) + "/"
	result := strings.ReplaceAll(name, " ", "")
	MainFunction := strings.Title(result)
	if option == 1{
		fmt.Println("you have selected a fresh expirement of type: ", SelectedType(Etype))
		fmt.Println("generating experiment file")
		Selectedmodule := SelectedType(Etype)
		if Selectedmodule == "Invalid"{
			log.Fatal("Invalid type selected")
			os.Exit(1)	
		}
		
		f, err := os.Create(subPath + result + ".go")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = f.WriteString(utils.GetExperimentTemplate(Selectedmodule, MainFunction))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(".......................................")
		fmt.Println("Generated Expirement File :", f.Name())
		fmt.Println(".......................................")
		fmt.Println("Generating Probes Functions For the file:", f.Name())
		createHelperFunctions(MainFunction)
		readfiletoappend(MainFunction, Selectedmodule, MainFunction)
		
		
		
	} else if option == 2{
		Selectedmodule := SelectedType(Etype)
		if Selectedmodule == "Invalid"{
			log.Fatal("Invalid type selected")
			os.Exit(1)	
		}
		files := getPresentExp(Selectedmodule)
		fmt.Println("Expirements Present in module are:")
		for i :=0; i< len(files); i++ {
			strtemp := fmt.Sprintf("[%d] %s",i, files[i].Name())
			fmt.Println(strtemp)
		}
		fmt.Print("Choose an option:")
		fmt.Scanf("%d", &OldExp)
		SelectedExp := files[OldExp].Name()
		SelectedExp = strings.Title(SelectedExp)
		SelectedExp = SelectedExp[:len(SelectedExp)-3]
		fmt.Println("Selected experiment is:", SelectedExp)
		fmt.Println(".......................................")
		fmt.Println("Generating Probes Functions For the file:", SelectedExp)
		createHelperFunctions(MainFunction)
		readfiletoappend(SelectedExp, Selectedmodule, MainFunction)

	} else {
		log.Fatal("Invalid input!")
	}
}

