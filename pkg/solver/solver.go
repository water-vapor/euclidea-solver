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
}

// NewParallelContext creates a new ParallelContext
func NewParallelContext(level int, threadLimit int) *ParallelContext {
	success := make(chan interface{})
	sema := NewSemaphore(threadLimit)
	var wg sync.WaitGroup
	return &ParallelContext{level, &wg, success, sema, 0}
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
	count := 0
	var newBoard *geom.Board
	var newSequence string

	recursion := func() {
		if useParallel && recursionLevel == ctx.parallelLevel-1 {
			ctx.wg.Add(1)
			ctx.sema.Down()
			go Solve(newBoard, newSequence, recursionLevel+1, st, ctx)
			count++
		} else {
			Solve(newBoard, newSequence, recursionLevel+1, st, ctx)
		}
	}

	shouldReturn := func() bool {
		if (useParallel && recursionLevel >= ctx.parallelLevel) || !useParallel {
			select {
			case <-ctx.success:
				return true
			default:
			}
		}
		return false
	}

	// signal wait groups on the level of go routines called
	if useParallel && recursionLevel == ctx.parallelLevel {
		defer ctx.wg.Done()
		defer ctx.sema.Up()
	}

	// output result
	if recursionLevel == 0 {
		defer func() {
			// output final result
			if recursionLevel == 0 {
				select {
				case <-ctx.success:
					fmt.Println("Solution found!")
				default:
					fmt.Println("Solution not found.")
				}
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
			}
		}
		for _, elem := range st.Target.Circles.Dict() {
			c := elem.(*geom.Circle)
			if !board.Circles.Contains(c) {
				solved = false
			}
		}
		for _, elem := range st.Target.Lines.Dict() {
			l := elem.(*geom.Line)
			if !board.Lines.Contains(l) {
				solved = false
			}
		}

		// statistics
		atomic.AddInt64(&ctx.searchCount, 1)

		if solved {
			_ = board.GeneratePlot(st.Name + "_" + st.Goal + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".png")
			// close the success channel to indicate success, all other routines should terminate
			// return on success
			select {
			// if already closed, just return
			case <-ctx.success:
				return
			default:
				close(ctx.success)
				return
			}
		}

		if len(sequence) == 0 {
			//_ = board.GeneratePlot("tmp_" + strconv.FormatInt(rand.Int63(), 10) + ".png")
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
			recursion()
			if shouldReturn() {
				return
			}
		}
	}

	// performs search
	switch sequence[0] {
	case 'O':
		// build circle based on existing points
		for _, elem1 := range board.Points.Dict() {
			for _, elem2 := range board.Points.Dict() {
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
				newBoard = board.Clone()
				newBoard.AddCircle(c)
				newSequence = sequence[1:]
				recursion()
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
				newBoard = board.Clone()
				newBoard.AddLine(l)
				newSequence = sequence[1:]
				recursion()
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
					newBoard = board.Clone()
					newBoard.AddLine(l)
					newSequence = sequence[1:]
					recursion()
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
				newBoard = board.Clone()
				newBoard.AddLine(l)
				newSequence = sequence[1:]
				recursion()
				if shouldReturn() {
					return
				}
			}
		}
	case 'L':
		tangentLineCandidates := make([]*geom.Line, 0)
		for _, elem1 := range board.Points.Values() {
			pt := elem1.(*geom.Point)
			for _, elem2 := range board.Lines.Values() {
				l := elem2.(*geom.Line)
				tangentLine := l.GetTangentLineWithPoint(pt)
				tangentLineCandidates = append(tangentLineCandidates, tangentLine)
			}
			for _, elem2 := range board.HalfLines.Values() {
				h := elem2.(*geom.HalfLine)
				l := geom.NewLineFromHalfLine(h)
				if board.Lines.Contains(l) {
					continue
				}
				tangentLine := l.GetTangentLineWithPoint(pt)
				tangentLineCandidates = append(tangentLineCandidates, tangentLine)
			}
			for _, elem2 := range board.Segments.Values() {
				s := elem2.(*geom.Segment)
				l := geom.NewLineFromSegment(s)
				if board.Lines.Contains(l) {
					continue
				}
				tangentLine := l.GetTangentLineWithPoint(pt)
				tangentLineCandidates = append(tangentLineCandidates, tangentLine)
			}
		}
		for _, tangentLine := range tangentLineCandidates {
			// line already exists in set
			if board.Lines.Contains(tangentLine) {
				continue
			}
			newBoard = board.Clone()
			newBoard.AddLine(tangentLine)
			newSequence = sequence[1:]
			recursion()
			if shouldReturn() {
				return
			}
		}
	case 'Z':
		parallelLineCandidates := make([]*geom.Line, 0)
		for _, elem1 := range board.Points.Values() {
			pt := elem1.(*geom.Point)
			for _, elem2 := range board.Lines.Values() {
				l := elem2.(*geom.Line)
				parallelLine := l.GetParallelLineWithPoint(pt)
				parallelLineCandidates = append(parallelLineCandidates, parallelLine)
			}
			for _, elem2 := range board.HalfLines.Values() {
				h := elem2.(*geom.HalfLine)
				l := geom.NewLineFromHalfLine(h)
				if board.Lines.Contains(l) {
					continue
				}
				parallelLine := l.GetParallelLineWithPoint(pt)
				parallelLineCandidates = append(parallelLineCandidates, parallelLine)
			}
			for _, elem2 := range board.Segments.Values() {
				s := elem2.(*geom.Segment)
				l := geom.NewLineFromSegment(s)
				if board.Lines.Contains(l) {
					continue
				}
				parallelLine := l.GetParallelLineWithPoint(pt)
				parallelLineCandidates = append(parallelLineCandidates, parallelLine)
			}
		}
		for _, parallelLine := range parallelLineCandidates {
			// line already exists in set
			if board.Lines.Contains(parallelLine) {
				continue
			}
			newBoard = board.Clone()
			newBoard.AddLine(parallelLine)
			newSequence = sequence[1:]
			recursion()
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
					newBoard = board.Clone()
					newBoard.AddCircle(c)
					newSequence = sequence[1:]
					recursion()
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

	// if parallel is used, wait
	if useParallel && recursionLevel == 0 {
		ctx.wg.Wait()
	}

	return
}
