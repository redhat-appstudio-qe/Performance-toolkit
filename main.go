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

//prettify - Done
// create a generator - flow chart done
//add 3 tests 
// create metrics 
// export to json 
// continue on error
// containerize 
// introduce logging - later