package infrastructure

	import (
		"context"
		"errors"
		"fmt"
	)
	
	//MainChaosFunction
	//This Function executes/injects chaos into the system
	func Nodespike(ctx context.Context) (error) {
		fmt.Println("replace with your logic")
		// Add your code here
		return nil
	}
	
	//RunnerFunction
	//This is the Runner Function which runs the above functions in a collabrative way
	func NodespikeExperiment(ctx context.Context){
		err := Nodespike(ctx)
		if err != nil{
			errors.New("Something bad happened!!")
		}
	}