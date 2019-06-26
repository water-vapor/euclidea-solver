package main

import (
	"flag"
	"fmt"
	"github.com/water-vapor/euclidea-solver/pkg/solver"
	"github.com/water-vapor/euclidea-solver/problem"
	"os"
	"strconv"
	"time"
)

func printUsage() {
	fmt.Printf("Usage: %s [-p] [-l level] [-t thread] [-v version] chapter_number problem_number\n", os.Args[0])
	flag.PrintDefaults()
}

func parseInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		printUsage()
		os.Exit(1)
	}
	return result
}

func main() {

	flag.Usage = printUsage

	useParallelPtr := flag.Bool("p", false, "Whether to enable parallel searching.")
	parallelLevelPtr := flag.Int("l", 1, "Distribute work to parallel workers at tree level.")
	threadLimitPtr := flag.Int("t", 100, "Max number of go routines running at any given time.")
	problemVersionPtr := flag.String("v", "E", "The goal of the problem, typically L or E. Default is E.")

	flag.Parse()

	positionalArgs := flag.Args()

	if len(positionalArgs) != 2 {
		fmt.Println("Incorrect number of positional arguments")
		printUsage()
		os.Exit(1)
	}

	chapter, number := parseInt(positionalArgs[0]), parseInt(positionalArgs[1])

	parallelLevel, threadLimit, goal := *parallelLevelPtr, *threadLimitPtr, *problemVersionPtr
	if !*useParallelPtr {
		parallelLevel = 0
	}

	st := problem.GetProblemByID(chapter, number, goal)
	goalSequence := st.GetSequenceByGoal()
	if parallelLevel >= len(goalSequence) {
		fmt.Println("Parallel level too deep for this problem, using ", len(goalSequence)-1)
		parallelLevel = len(goalSequence) - 1
	}

	start := time.Now()

	ctx := solver.NewParallelContext(parallelLevel, threadLimit)

	solver.Solve(st.Board, goalSequence, 0, st, ctx)

	fmt.Println("Number of boards searched:", ctx.GetSearchCount())

	end := time.Since(start)

	fmt.Println("Took", end)
}
