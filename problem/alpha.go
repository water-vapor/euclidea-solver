package problem

import (
	"github.com/water-vapor/euclidea-solver/pkg/geom"
	"math"
)

// Board 1: Angel Of 60 Degree
func angelOf60Degree() *Statement {
	problem := geom.NewGeomBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(1, math.Sqrt(3))

	problem.Points.Add(pt1)
	problem.HalfLines.Add(geom.NewHalfLineFromDirection(pt1, geom.NewVector2D(1, 0)))

	target := geom.NewGoal()
	target.Lines.Add(geom.NewLineFromTwoPoints(pt1, pt2))
	return NewStatement(problem, target, "OOI", "Angel Of 60 Degree")
}
