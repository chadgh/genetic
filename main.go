package main

import (
	"fmt"
	"os"

	"github.com/chadgh/genetic/catdog"
	"github.com/chadgh/genetic/eightqueens"
	"github.com/chadgh/genetic/queens"
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
