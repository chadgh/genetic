package main

import (
	"fmt"
	"os"

	"gitlab.com/chadgh/genetic/catdog"
	"gitlab.com/chadgh/genetic/eightqueens"
	"gitlab.com/chadgh/genetic/queens"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Exiting")
		return
	}
	toRun := args[1]

	switch toRun {
	case "catdog":
		catdog.Run()
	case "queens":
		queens.Run()
	case "8":
		eightqueens.Run()
	default:
		fmt.Println("Doing nothing")
	}
}
