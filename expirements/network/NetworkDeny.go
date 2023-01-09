package network

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
)

//MainChaosFunction
//This Function executes/injects chaos into the system
	func NetworkDeny(ctx context.Context) (error) {
		fmt.Println("replace with your logic")
		PolicyYaml := ctx.Value("PolicyFile").(string)
		// Add your code here
		ns := ctx.Value("chaos_ns").(string)
		cmd := exec.Command("kubectl", "apply", "-f", PolicyYaml, "-n", ns)
    	stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(stdout))
		return nil
	}
	
	//RunnerFunction
	//This is the Runner Function which runs the above functions in a collabrative way
	func NetworkDenyExperiment(ctx context.Context){
		err := NetworkDeny(ctx)
		if err != nil{
			errors.New("Something bad happened!!")
		}
	}