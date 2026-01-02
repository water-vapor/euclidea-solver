// Package solver implements the DFS search algorithm for geometric construction searching
package solver

import (
	"fmt"
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/geom"
	"github.com/water-vapor/euclidea-solver/problem"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// Semaphore is a synchronous primitive that limits the number of go routines in buffer
type Semaphore struct {
	c chan struct{}
}

// NewSemaphore creates a new semaphore object
func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{make(chan struct{}, limit)}
}

// Down should be called when adding an object to buffer
func (s *Semaphore) Down() {
	s.c <- struct{}{}
}

// Up should be called when removing an object to buffer
func (s *Semaphore) Up() {
	<-s.c
}

// ParallelContext holds synchronous primitives and information used by solver
type ParallelContext struct {
	parallelLevel int
	wg            *sync.WaitGroup
	success       chan interface{}
	sema          *Semaphore
	searchCount   int64
	closeOnce     sync.Once // Ensures success channel is closed exactly once
}

// NewParallelContext creates a new ParallelContext
func NewParallelContext(level int, threadLimit int) *ParallelContext {
	success := make(chan interface{})
	sema := NewSemaphore(threadLimit)
	var wg sync.WaitGroup
	return &ParallelContext{parallelLevel: level, wg: &wg, success: success, sema: sema, searchCount: 0}
}

// GetSearchCount outputs the number of full searches
func (ctx *ParallelContext) GetSearchCount() int64 {
	return ctx.searchCount
}

// Solve implements the DFS search algorithm
func Solve(board *geom.Board, sequence string, recursionLevel int,
	st *problem.Statement, ctx *ParallelContext) {
	// Checks if required objects have been found.
	// This happens every step if early stopping is enabled,
	// or the search sequence is exhausted.
	useParallel := ctx.parallelLevel != 0
	shouldLaunchWorkers := useParallel && recursionLevel == ctx.parallelLevel-1
	count := 0
	var newBoard *geom.Board
	var newSequence string

	getNewBoard := func(b *geom.Board) *geom.Board {
		if shouldLaunchWorkers {
			return b.Clone()
		}
		return b
	}

	// Check if solution was found (used at all levels for early termination)
	isSolved := func() bool {
		select {
		case <-ctx.success:
			return true
		default:
			return false
		}
	}

	recursion := func(shouldBacktrack bool) {
		if shouldLaunchWorkers {
			// Check before spawning to avoid wasted work after solution found
			if isSolved() {
				return
			}
			ctx.wg.Add(1)
			ctx.sema.Down()
			go Solve(newBoard, newSequence, recursionLevel+1, st, ctx)
			count++
		} else {
			Solve(newBoard, newSequence, recursionLevel+1, st, ctx)
			if shouldBacktrack {
				newBoard.RemoveLastGeometryObject()
			}
		}
	}

	shouldReturn := func() bool {
		return isSolved()
	}

	// signal wait groups on the level of go routines called
	if useParallel && recursionLevel == ctx.parallelLevel {
		defer ctx.wg.Done()
		defer ctx.sema.Up()
	}

	// At level 0, ensure we wait for workers and output result
	if recursionLevel == 0 && useParallel {
		// Defers run in LIFO order, so register output first, then wait
		defer func() {
			select {
			case <-ctx.success:
				fmt.Println("Solution found!")
			default:
				fmt.Println("Solution not found.")
			}
		}()
		// This runs BEFORE the output (LIFO) - ensures workers complete before reading searchCount
		defer ctx.wg.Wait()
	} else if recursionLevel == 0 {
		// Non-parallel mode: just output result
		defer func() {
			select {
			case <-ctx.success:
				fmt.Println("Solution found!")
			default:
				fmt.Println("Solution not found.")
			}
		}()
	}

	// terminate subsequent calls after success
	if shouldReturn() {
		return
	}

	if configs.EarlyStop || len(sequence) == 0 {
		solved := true
		for _, elem := range st.Target.Points.Dict() {
			pt := elem.(*geom.Point)
			if !board.Points.Contains(pt) {
				solved = false
				break
			}
		}
		if solved {
			for _, elem := range st.Target.Circles.Dict() {
				c := elem.(*geom.Circle)
				if !board.Circles.Contains(c) {
					solved = false
					break
				}
			}
		}
		if solved {
			for _, elem := range st.Target.Lines.Dict() {
				l := elem.(*geom.Line)
				if !board.Lines.Contains(l) {
					solved = false
					break
				}
			}
		}

		// statistics
		atomic.AddInt64(&ctx.searchCount, 1)

		if solved {
			ctx.closeOnce.Do(func() {
				_ = board.GeneratePlot(st.Name + "_" + st.Goal + "_" + strconv.FormatInt(time.Now().Unix(), 10))
				close(ctx.success)
			})
			return
		}

		if len(sequence) == 0 {
			//_ = board.GeneratePlot("tmp_" + strconv.FormatInt(rand.Int63(), 10))
			return
		}
	}

	// no enough points
	if len(board.Points.Dict()) < 2 {
		pts := board.GenerateRandomPoints()
		for _, pt := range pts {
			newBoard = board.Clone()
			newBoard.AddPoint(pt)
			// proceed without decreasing sequence
			newSequence = sequence
			recursion(false)
			if shouldReturn() {
				return
			}
		}
	}

	// performs search
	switch sequence[0] {
	case 'O':
		// build circle based on existing points
		// Use Values() to snapshot points - iterating Dict() while AddCircleTrace
		// modifies the map is undefined behavior in Go
		pointSet := board.Points.Values()
		for _, elem1 := range pointSet {
			for _, elem2 := range pointSet {
				pt1 := elem1.(*geom.Point)
				pt2 := elem2.(*geom.Point)
				// same point object
				if pt1.Equal(pt2) {
					continue
				}
				c := geom.NewCircleByPoint(pt1, pt2)
				// circle already exists in set
				if board.Circles.Contains(c) {
					continue
				}
				newBoard = getNewBoard(board)
				newBoard.AddCircleTrace(c)
				newSequence = sequence[1:]
				recursion(true)
				if shouldReturn() {
					return
				}
			}
		}
	case 'I':
		pointSet := board.Points.Values()
		for index1, elem1 := range pointSet {
			for index2, elem2 := range pointSet {
				// two points must be ordered to ensure no duplicated lines are created
				if index1 >= index2 {
					continue
				}
				pt1 := elem1.(*geom.Point)
				pt2 := elem2.(*geom.Point)
				l := geom.NewLineFromTwoPoints(pt1, pt2)

				// line already exists in set
				if board.Lines.Contains(l) {
					continue
				}
				newBoard = getNewBoard(board)
				newBoard.AddLineTrace(l)
				newSequence = sequence[1:]
				recursion(true)

				if shouldReturn() {
					return
				}
			}
		}
	case 'A':
		pointSet := board.Points.Values()
		for index2, elem2 := range pointSet {
			for index1, elem1 := range pointSet {
				for index3, elem3 := range pointSet {
					// points on side must be ordered;
					// non of them should be same
					if index1 >= index3 || index1 == index2 || index2 == index3 {
						continue
					}
					pt1 := elem1.(*geom.Point)
					pt2 := elem2.(*geom.Point)
					pt3 := elem3.(*geom.Point)
					l := geom.NewLineAsAngleBisector(pt1, pt2, pt3)
					// line already exists in set
					if board.Lines.Contains(l) {
						continue
					}
					newBoard = getNewBoard(board)
					newBoard.AddLineTrace(l)
					newSequence = sequence[1:]
					recursion(true)
					if shouldReturn() {
						return
					}
				}
			}
		}
	case '+':
		pointSet := board.Points.Values()
		for index1, elem1 := range pointSet {
			for index2, elem2 := range pointSet {
				// end points of segment must be ordered
				if index1 >= index2 {
					continue
				}
				pt1 := elem1.(*geom.Point)
				pt2 := elem2.(*geom.Point)
				l := geom.NewSegment(pt1, pt2).Bisector()
				// line already exists in set
				if board.Lines.Contains(l) {
					continue
				}
				newBoard = getNewBoard(board)
				newBoard.AddLineTrace(l)
				newSequence = sequence[1:]
				recursion(true)
				if shouldReturn() {
					return
				}
			}
		}
	case 'L':
		// Use map to dedupe candidates by hash (avoid exploring identical branches)
		tangentLineSeen := make(map[interface{}]*geom.Line)
		for _, elem1 := range board.Points.Values() {
			pt := elem1.(*geom.Point)
			for _, elem2 := range board.Lines.Values() {
				l := elem2.(*geom.Line)
				tangentLine := l.GetTangentLineWithPoint(pt)
				tangentLineSeen[tangentLine.Serialize()] = tangentLine
			}
			for _, elem2 := range board.HalfLines.Values() {
				h := elem2.(*geom.HalfLine)
				l := geom.NewLineFromHalfLine(h)
				if board.Lines.Contains(l) {
					continue
				}
				tangentLine := l.GetTangentLineWithPoint(pt)
				tangentLineSeen[tangentLine.Serialize()] = tangentLine
			}
			for _, elem2 := range board.Segments.Values() {
				s := elem2.(*geom.Segment)
				l := geom.NewLineFromSegment(s)
				if board.Lines.Contains(l) {
					continue
				}
				tangentLine := l.GetTangentLineWithPoint(pt)
				tangentLineSeen[tangentLine.Serialize()] = tangentLine
			}
		}
		for _, tangentLine := range tangentLineSeen {
			// line already exists in set
			if board.Lines.Contains(tangentLine) {
				continue
			}
			newBoard = getNewBoard(board)
			newBoard.AddLineTrace(tangentLine)
			newSequence = sequence[1:]
			recursion(true)
			if shouldReturn() {
				return
			}
		}
	case 'Z':
		// Use map to dedupe candidates by hash (avoid exploring identical branches)
		parallelLineSeen := make(map[interface{}]*geom.Line)
		for _, elem1 := range board.Points.Values() {
			pt := elem1.(*geom.Point)
			for _, elem2 := range board.Lines.Values() {
				l := elem2.(*geom.Line)
				parallelLine := l.GetParallelLineWithPoint(pt)
				parallelLineSeen[parallelLine.Serialize()] = parallelLine
			}
			for _, elem2 := range board.HalfLines.Values() {
				h := elem2.(*geom.HalfLine)
				l := geom.NewLineFromHalfLine(h)
				if board.Lines.Contains(l) {
					continue
				}
				parallelLine := l.GetParallelLineWithPoint(pt)
				parallelLineSeen[parallelLine.Serialize()] = parallelLine
			}
			for _, elem2 := range board.Segments.Values() {
				s := elem2.(*geom.Segment)
				l := geom.NewLineFromSegment(s)
				if board.Lines.Contains(l) {
					continue
				}
				parallelLine := l.GetParallelLineWithPoint(pt)
				parallelLineSeen[parallelLine.Serialize()] = parallelLine
			}
		}
		for _, parallelLine := range parallelLineSeen {
			// line already exists in set
			if board.Lines.Contains(parallelLine) {
				continue
			}
			newBoard = getNewBoard(board)
			newBoard.AddLineTrace(parallelLine)
			newSequence = sequence[1:]
			recursion(true)
			if shouldReturn() {
				return
			}
		}
	case 'Q':
		pointSet := board.Points.Values()
		// segment on 1,3 create on 2
		for index2, elem2 := range pointSet {
			for index1, elem1 := range pointSet {
				for index3, elem3 := range pointSet {
					// points on segment must be ordered;
					// non of them should be same
					if index1 >= index3 || index1 == index2 || index2 == index3 {
						continue
					}
					pt1 := elem1.(*geom.Point)
					pt2 := elem2.(*geom.Point)
					pt3 := elem3.(*geom.Point)
					d := geom.NewSegment(pt1, pt3)
					c := geom.NewCircleByRadius(pt2, d.Length())
					// line already exists in set
					if board.Circles.Contains(c) {
						continue
					}
					newBoard = getNewBoard(board)
					newBoard.AddCircleTrace(c)
					newSequence = sequence[1:]
					recursion(true)
					if shouldReturn() {
						return
					}
				}
			}
		}
	default:
		panic("operation " + sequence[0:1] + " is not implemented")
	}

	if useParallel && recursionLevel == ctx.parallelLevel-1 {
		fmt.Println(count, "go routines deployed.")
	}
}
