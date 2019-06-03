package solver

import (
	"fmt"
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/geom"
	"strconv"
	"sync"
	"time"
)

func Solve(board *geom.Board, goal *geom.Goal, sequence string, name string, recursionLevel int,
	success chan interface{}, wg *sync.WaitGroup, parallelLevel int) {
	// Checks if required objects have been found.
	// This happens every step if early stopping is enabled,
	// or the search sequence is exhausted.
	useParallel := parallelLevel != 0
	count := 0
	// signal wait groups on the level of go routines called
	if useParallel && recursionLevel == parallelLevel {
		defer wg.Done()
	}
	if configs.EarlyStop || len(sequence) == 0 {
		solved := true
		for _, elem := range goal.Points.Dict() {
			pt := elem.(*geom.Point)
			if !board.Points.Contains(pt) {
				solved = false
			}
		}
		for _, elem := range goal.Circles.Dict() {
			c := elem.(*geom.Circle)
			if !board.Circles.Contains(c) {
				solved = false
			}
		}
		for _, elem := range goal.Lines.Dict() {
			l := elem.(*geom.Line)
			if !board.Lines.Contains(l) {
				solved = false
			}
		}
		if solved {
			_ = board.GeneratePlot(name + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".png")
			// close the success channel to indicate success, all other routines should terminate
			close(success)
			return
		}
		if len(sequence) == 0 {
			//_ = board.GeneratePlot("tmp_" + strconv.FormatInt(rand.Int63(), 10) + ".png")
			return
		}
	}
	// performs search
	switch sequence[0] {
	case 'O':
		validPointPairCount := 0
		// build circle based on existing points
		for _, elem1 := range board.Points.Dict() {
			for _, elem2 := range board.Points.Dict() {
				pt1 := elem1.(*geom.Point)
				pt2 := elem2.(*geom.Point)
				// same point object
				if pt1 == pt2 {
					continue
				}
				c := geom.NewCircleByPoint(pt1, pt2)
				// circle already exists in set
				if board.Circles.Contains(c) {
					continue
				}
				validPointPairCount++
				newBoard := board.Clone()
				newBoard.AddCircle(c)

				if useParallel && recursionLevel == parallelLevel-1 {
					wg.Add(1)
					go Solve(newBoard, goal, sequence[1:], name, recursionLevel+1, success, wg, parallelLevel)
					count++
				} else {
					Solve(newBoard, goal, sequence[1:], name, recursionLevel+1, success, wg, parallelLevel)
				}
				if (useParallel && recursionLevel >= parallelLevel) || !useParallel {
					// return on success
					select {
					case <-success:
						return
					default:
					}
				}

			}
		}
		// no valid points, add a point on an object and do not increase level
		if validPointPairCount == 0 {
			pts := board.GenerateRandomPoints()
			for _, pt := range pts {
				newBoard := board.Clone()
				newBoard.AddPoint(pt)
				// proceed silently
				Solve(newBoard, goal, sequence, name, recursionLevel, success, wg, parallelLevel)
				if (useParallel && recursionLevel >= parallelLevel) || !useParallel {
					// return on success
					select {
					case <-success:
						return
					default:
					}
				}
			}
		}
	case 'I':
		validPointPairCount := 0
		for _, elem1 := range board.Points.Dict() {
			for _, elem2 := range board.Points.Dict() {
				pt1 := elem1.(*geom.Point)
				pt2 := elem2.(*geom.Point)
				// same point object
				if pt1 == pt2 {
					continue
				}
				l := geom.NewLineFromTwoPoints(pt1, pt2)

				// circle already exists in set
				if board.Lines.Contains(l) {
					continue
				}
				validPointPairCount++
				newBoard := board.Clone()
				newBoard.AddLine(l)
				if useParallel && recursionLevel == parallelLevel-1 {
					wg.Add(1)
					go Solve(newBoard, goal, sequence[1:], name, recursionLevel+1, success, wg, parallelLevel)
					count++
				} else {
					Solve(newBoard, goal, sequence[1:], name, recursionLevel+1, success, wg, parallelLevel)
				}
				if (useParallel && recursionLevel >= parallelLevel) || !useParallel {
					// return on success
					select {
					case <-success:
						return
					default:
					}
				}
			}
		}
		// no valid points
		if validPointPairCount == 0 {
			pts := board.GenerateRandomPoints()
			for _, pt := range pts {
				newBoard := board.Clone()
				newBoard.AddPoint(pt)
				// proceed silently
				Solve(newBoard, goal, sequence, name, recursionLevel, success, wg, parallelLevel)
				if (useParallel && recursionLevel >= parallelLevel) || !useParallel {
					// return on success
					select {
					case <-success:
						return
					default:
					}
				}
			}
		}
	default:
		panic("operation " + sequence[0:1] + " is not implemented")
	}
	if useParallel && recursionLevel == parallelLevel-1 {
		fmt.Println(count, "go routines deployed.")
	}
	return
}
