package main

import (
	"fmt"
	"github.com/water-vapor/euclidea-solver/pkg/solver"
	"github.com/water-vapor/euclidea-solver/problem"
	"os"
	"strconv"
	"sync"
)

func printUsageAndExit() {
	fmt.Println("Usage: euclidea-solver [-p[=parallel level]] <chapter number> <problem number>")
	fmt.Println("\t<chapter number> = integer, ID of chapter.")
	fmt.Println("\t<problem number> = integer, ID of problem.")
	fmt.Println("\t-p[=parallel level] = An optional flag to run the editor in its parallel version.")
	fmt.Println("\t\tYou also have the option of specifying the threading level")
	fmt.Println("\t[parallel level] = distribute work to parallel workers at tree level")
	os.Exit(0)
}

// handle command line arguments
func processArgs(args []string) (int, int, int) {
	argumentLength := len(args)
	var cmdOption string
	var parallelLevel int
	var chapter, number int
	switch argumentLength {
	case 1:
		fmt.Println("No argument!")
		printUsageAndExit()
	case 3:
		chapter, _ = strconv.Atoi(args[1])
		number, _ = strconv.Atoi(args[2])
		parallelLevel = 0
	case 4:
		cmdOption = args[1]
		chapter, _ = strconv.Atoi(args[2])
		number, _ = strconv.Atoi(args[3])
		// process parallel arguments
		// parallel level = 0 -> sequential
		// use level 1 as default
		if cmdOption == "-p" {
			parallelLevel = 1
		} else if cmdOption[:3] == "-p=" {
			parallelLevel, _ = strconv.Atoi(cmdOption[3:])
			if parallelLevel <= 0 {
				fmt.Println("Number of parallel level too small, resetting to 1.")
				parallelLevel = 1
			}
		} else {
			fmt.Println("Invalid argument!")
			printUsageAndExit()
		}
	default:
		fmt.Println("Incorrect number of arguments!")
		printUsageAndExit()
	}
	return parallelLevel, chapter, number
}

func main() {

	parallelLevel, chapter, number := processArgs(os.Args)

	success := make(chan interface{})
	var wg sync.WaitGroup

	ps := problem.GetProblemByID(chapter, number)
	if parallelLevel >= len(ps.Sequence) {
		fmt.Println("Parallel level too deep for this problem, using ", len(ps.Sequence)-1)
		parallelLevel = len(ps.Sequence) - 1
	}
	solver.Solve(ps.Board, ps.Target, ps.Sequence, ps.Name, 0, success, &wg, parallelLevel)

	// if parallel is used, wait
	if parallelLevel != 0 {
		wg.Wait()
	}
	// output search result in console
	select {
	case <-success:
		fmt.Println("Solution found!")
	default:
		fmt.Println("Solution not found.")
	}
}
