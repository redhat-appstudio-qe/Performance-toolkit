package infrastructure

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
)

//MainChaosFunction
//This Function executes/injects chaos into the system
	func NodeDrain(ctx context.Context) (error) {
		TargetNode := ctx.Value("NodeToBeDrained").(string)
		fmt.Println("Executing Node Drain")
		cmd := exec.Command("kubectl", "drain", TargetNode, "--ignore-daemonsets", "--delete-emptydir-data", "--force")
		//cmd := exec.Command("kubectl", "drain", TargetNode)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return err
		}
		fmt.Println("Result: " + out.String())
		//fmt.Println(cmd.Args)
		/* var out, stdErr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stdErr */
    	// stdout, err := cmd.Output()
		/* clea */
		//fmt.Println(string(output))
		// fmt.Println(string(stdout))
		fmt.Println(`Node Drained Successfully:`, TargetNode)
		return nil
	}
	
	//RunnerFunction
	//This is the Runner Function which runs the above functions in a collabrative way
	func NodeDrainExperiment(ctx context.Context){
		err := NodeDrain(ctx)
		if err != nil{
			errors.New("Something bad happened!!")
		}
	}