package main

import (
	"fmt"
	"os"

	"gitlab.com/chadgh/genetic/queens"
	// "gitlab.com/chadgh/genetic/catdog"
	// "gitlab.com/chadgh/genetic/napsack"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Exiting")
		return
	}
	toRun := args[1]

	switch toRun {
	// case "projectestimation":
	//	projectestimation.Run()
	// case "catdog":
	// 	catdog.Run()
	// case "napsack":
	// 	napsack.Run()
	case "queens":
		queens.Run()
	default:
		fmt.Println("Doing nothing")
	}
}
