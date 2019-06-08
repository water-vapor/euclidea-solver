package main

import (
	"fmt"
	"github.com/water-vapor/euclidea-solver/pkg/solver"
	"github.com/water-vapor/euclidea-solver/problem"
	"os"
	"strconv"
	"time"
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

func parsePArgs(cmdOption string) int {
	// process parallel arguments
	// parallel level = 0 -> sequential
	// use level 1 as default
	var parallelLevel int
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
	return parallelLevel
}

func parseInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		printUsageAndExit()
	}
	return result
}

// handle command line arguments
func parseArgs(args []string) (int, int, int, string) {
	argumentLength := len(args)
	var cmdOption, goal string
	var parallelLevel int
	var chapter, number int
	switch argumentLength {
	case 1:
		fmt.Println("No argument!")
		printUsageAndExit()
	case 3:
		chapter = parseInt(args[1])
		number = parseInt(args[2])
		parallelLevel = 0
		goal = "E"
	case 4:
		if args[1][:2] == "-p" {
			cmdOption = args[1]
			chapter = parseInt(args[2])
			number = parseInt(args[3])
			parallelLevel = parsePArgs(cmdOption)
			goal = "E"
		} else {
			chapter = parseInt(args[1])
			number = parseInt(args[2])
			goal = args[3]
			parallelLevel = 0
		}
	case 5:
		cmdOption = args[1]
		chapter = parseInt(args[2])
		number = parseInt(args[3])
		goal = args[4]
		parallelLevel = parsePArgs(cmdOption)
	default:
		fmt.Println("Incorrect number of arguments!")
		printUsageAndExit()
	}
	return parallelLevel, chapter, number, goal
}

func main() {

	parallelLevel, chapter, number, goal := parseArgs(os.Args)

	st := problem.GetProblemByID(chapter, number, goal)
	goalSequence := st.GetSequenceByGoal()
	if parallelLevel >= len(goalSequence) {
		fmt.Println("Parallel level too deep for this problem, using ", len(goalSequence)-1)
		parallelLevel = len(goalSequence) - 1
	}

	start := time.Now()

	ctx := solver.NewParallelContext(parallelLevel)

	solver.Solve(st.Board, goalSequence, 0,st,ctx)

	// if parallel is used, wait
	if parallelLevel != 0 {
		ctx.Wg.Wait()
	}
	// output search result in console
	select {
	case <-ctx.Success:
		fmt.Println("Solution found!")
	default:
		fmt.Println("Solution not found.")
	}

	end := time.Since(start)
	fmt.Println("Took", end)
}
