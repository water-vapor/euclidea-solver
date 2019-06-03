package problem

import (
	"github.com/water-vapor/euclidea-solver/pkg/geom"
	"math"
)

// Problem 1: Angel Of 60 Degree
func angelOf60Degree() *Statement {
	problem := geom.NewGeomBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(1, math.Sqrt(3))

	problem.AddPoint(pt1)
	problem.HalfLines.Add(geom.NewHalfLineFromDirection(pt1, geom.NewVector2D(1, 0)))

	target := geom.NewGoal()
	target.Lines.Add(geom.NewLineFromTwoPoints(pt1, pt2))
	return NewStatement(problem, target, "OOI", "Angel Of 60 Degree")
}

//Problem 2: Perpendicular Bisector
func perpendicularBisector() *Statement {
	problem := geom.NewGeomBoard()
	pt1 := geom.NewPoint(-1, 0)
	pt2 := geom.NewPoint(1, 0)
	s := geom.NewSegment(pt1, pt2)
	l := s.Bisector()

	problem.AddSegment(s)

	target := geom.NewGoal()
	target.Lines.Add(l)
	return NewStatement(problem, target, "OOI", "Perpendicular Bisector")
}

//Problem 6: Circle Center
func circleCenter() *Statement {
	problem := geom.NewGeomBoard()
	pt1 := geom.NewPoint(0, 0)
	c := geom.NewCircleByRadius(pt1, 2)

	problem.AddCircle(c)

	target := geom.NewGoal()
	target.Points.Add(pt1)
	return NewStatement(problem, target, "OOOII", "Circle Center")
}

//Problem 7: Inscribed Square
func inscribedSquare() *Statement {
	// the last two unnecessary lines are removed
	problem := geom.NewGeomBoard()
	pt1 := geom.NewPoint(0, 0)
	c := geom.NewCircleByRadius(pt1, 2)

	problem.AddCircle(c)
	problem.AddPoint(pt1)
	problem.AddPoint(geom.NewPoint(0, 2))

	target := geom.NewGoal()
	pt2 := geom.NewPoint(2, 0)
	pt3 := geom.NewPoint(-2, 0)
	pt4 := geom.NewPoint(0, -2)
	target.Points.Add(pt2)
	target.Points.Add(pt3)
	target.Points.Add(pt4)
	target.Lines.Add(geom.NewLineFromTwoPoints(pt2, pt4))
	target.Lines.Add(geom.NewLineFromTwoPoints(pt3, pt4))
	return NewStatement(problem, target, "OOIII", "Inscribed Square")
}
