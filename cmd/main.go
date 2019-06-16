package main

import (
	"fmt"
	"github.com/water-vapor/euclidea-solver/pkg/solver"
	"github.com/water-vapor/euclidea-solver/problem"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func printUsageAndExit() {
	fmt.Println("Usage: euclidea-solver [-p[=parallel_level[,thread_limit]]] <chapter_number> <problem_number> [problem_version]")
	fmt.Println("\t<chapter_number> = integer, ID of chapter.")
	fmt.Println("\t<problem_number> = integer, ID of problem.")
	fmt.Println("\t-p[=parallel_level] = An optional flag to run the editor in its parallel version.")
	fmt.Println("\t\tYou also have the option of specifying the threading level.")
	fmt.Println("\t[parallel_level] = distribute work to parallel workers at tree level.")
	fmt.Println("\t[thread_limit] = max number of go routines running at any given time.")
	fmt.Println("\t[problem_version] = the goal of the problem, typically L or E. Default is E.")
	os.Exit(0)
}

func parseInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		printUsageAndExit()
	}
	return result
}

func parsePArgs(cmdOption string) (int, int) {
	// process parallel arguments
	// parallel level = 0 -> sequential
	// use level 1 as default
	var parallelLevel, threadLimit int
	if cmdOption == "-p" {
		parallelLevel = 1
	} else if cmdOption[:3] == "-p=" {
		optionsStr := strings.Split(cmdOption[3:], ",")
		optionsInt := make([]int, len(optionsStr))
		for idx, elem := range optionsStr {
			optionsInt[idx] = parseInt(elem)
		}
		switch len(optionsInt) {
		case 1:
			parallelLevel = optionsInt[0]
			threadLimit = runtime.NumCPU()
		case 2:
			parallelLevel = optionsInt[0]
			threadLimit = optionsInt[1]
		default:
			fmt.Println("Invalid argument!")
			printUsageAndExit()
		}
		if parallelLevel <= 0 {
			fmt.Println("Number of parallel_level too small, resetting to 1.")
			parallelLevel = 1
		}
		if threadLimit <= 0 {
			fmt.Println("Number of thread limit too small, resetting to 1.")
			threadLimit = 1
		}
	} else {
		fmt.Println("Invalid argument!")
		printUsageAndExit()
	}
	return parallelLevel, threadLimit
}

// handle command line arguments
func parseArgs(args []string) (int, int, int, int, string) {
	argumentLength := len(args)
	var cmdOption, goal string
	var parallelLevel, threadLimit int
	var chapter, number int
	switch argumentLength {
	case 1:
		fmt.Println("No argument!")
		printUsageAndExit()
	case 3:
		chapter = parseInt(args[1])
		number = parseInt(args[2])
		parallelLevel = 0
		threadLimit = 1
		goal = "E"
	case 4:
		if args[1][:2] == "-p" {
			cmdOption = args[1]
			chapter = parseInt(args[2])
			number = parseInt(args[3])
			parallelLevel, threadLimit = parsePArgs(cmdOption)
			goal = "E"
		} else {
			chapter = parseInt(args[1])
			number = parseInt(args[2])
			goal = args[3]
			parallelLevel = 0
			threadLimit = 1
		}
	case 5:
		cmdOption = args[1]
		chapter = parseInt(args[2])
		number = parseInt(args[3])
		goal = args[4]
		parallelLevel, threadLimit = parsePArgs(cmdOption)
	default:
		fmt.Println("Incorrect number of arguments!")
		printUsageAndExit()
	}
	return parallelLevel, threadLimit, chapter, number, goal
}

func main() {

	parallelLevel, threadLimit, chapter, number, goal := parseArgs(os.Args)

	st := problem.GetProblemByID(chapter, number, goal)
	goalSequence := st.GetSequenceByGoal()
	if parallelLevel >= len(goalSequence) {
		fmt.Println("Parallel level too deep for this problem, using ", len(goalSequence)-1)
		parallelLevel = len(goalSequence) - 1
	}

	start := time.Now()

	ctx := solver.NewParallelContext(parallelLevel, threadLimit)

	solver.Solve(st.Board, goalSequence, 0, st, ctx)

	end := time.Since(start)

	fmt.Println("Took", end)
}
