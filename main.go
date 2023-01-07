package main

import (
	"context"

	"github.com/redhat-appstudio-qe/performance-toolkit/state"
	"fmt"

)

func main(){

	fmt.Println("Start Chaos Test!")
	ctx := context.Background();


	state.IsSystemRunning(ctx);

	ctx, err := state.SteadyState(ctx);

	if err != nil {
		fmt.Println("error occured!");
	}

	state.Chaos(ctx);

	
}