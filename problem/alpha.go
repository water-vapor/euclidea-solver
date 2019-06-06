package problem

import (
	"github.com/water-vapor/euclidea-solver/pkg/geom"
	"math"
)

// Problem 1: Angel Of 60 Degree
func angelOf60Degree() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(1, math.Sqrt(3))

	problem.AddPoint(pt1)
	problem.HalfLines.Add(geom.NewHalfLineFromDirection(pt1, geom.NewVector2D(1, 0)))

	target := geom.NewGoal()
	target.Lines.Add(geom.NewLineFromTwoPoints(pt1, pt2))
	return NewStatement(problem, target, "OOI", "1.1 Angel Of 60 Degree")
}

//Problem 2: Perpendicular Bisector
func perpendicularBisector() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(-1, 0)
	pt2 := geom.NewPoint(1, 0)
	s := geom.NewSegment(pt1, pt2)
	l := s.Bisector()

	problem.AddSegment(s)

	target := geom.NewGoal()
	target.Lines.Add(l)
	return NewStatement(problem, target, "OOI", "1.2 Perpendicular Bisector")
}

//Problem 3: Midpoint
func midpoint() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(-1, 0)
	pt2 := geom.NewPoint(1, 0)
	problem.AddPoint(pt1)
	problem.AddPoint(pt2)

	pt3 := geom.NewPoint(0, 0)
	target := geom.NewGoal()
	target.Points.Add(pt3)
	return NewStatement(problem, target, "I+", "1.3 Midpoint")
}

//Problem 4: Circle in Square
func circleInSquare() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(-1, -1)
	pt2 := geom.NewPoint(-1, 1)
	pt3 := geom.NewPoint(1, 1)
	pt4 := geom.NewPoint(1, -1)
	s1 := geom.NewSegment(pt1, pt2)
	s2 := geom.NewSegment(pt2, pt3)
	s3 := geom.NewSegment(pt3, pt4)
	s4 := geom.NewSegment(pt4, pt1)
	problem.AddPoint(pt1)
	problem.AddPoint(pt2)
	problem.AddPoint(pt3)
	problem.AddPoint(pt4)
	problem.AddSegment(s1)
	problem.AddSegment(s2)
	problem.AddSegment(s3)
	problem.AddSegment(s4)

	target := geom.NewGoal()
	target.Circles.Add(geom.NewCircleByRadius(geom.NewPoint(0, 0), 1))
	return NewStatement(problem, target, "I+O", "1.4 Circle in Square")
}

//Problem 5: Rhombus in Rectangle
func rhombusInRectangle() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(3, 0)
	pt3 := geom.NewPoint(3, math.Sqrt(3))
	pt4 := geom.NewPoint(0, math.Sqrt(3))
	s1 := geom.NewSegment(pt1, pt2)
	s2 := geom.NewSegment(pt2, pt3)
	s3 := geom.NewSegment(pt3, pt4)
	s4 := geom.NewSegment(pt4, pt1)
	problem.AddPoint(pt1)
	problem.AddPoint(pt2)
	problem.AddPoint(pt3)
	problem.AddPoint(pt4)
	problem.AddSegment(s1)
	problem.AddSegment(s2)
	problem.AddSegment(s3)
	problem.AddSegment(s4)

	target := geom.NewGoal()
	pt5 := geom.NewPoint(2, 0)
	pt6 := geom.NewPoint(1, math.Sqrt(3))
	l1 := geom.NewLineFromTwoPoints(pt1, pt6)
	l2 := geom.NewLineFromTwoPoints(pt5, pt3)
	target.Lines.Add(l1)
	target.Lines.Add(l2)
	return NewStatement(problem, target, "+II", "1.5 Rhombus in Rectangle")
}

//Problem 6: Circle Center
func circleCenter() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, 0)
	c := geom.NewCircleByRadius(pt1, 2)

	problem.AddCircle(c)

	target := geom.NewGoal()
	target.Points.Add(pt1)
	return NewStatement(problem, target, "OOOII", "1.6 Circle Center")
}

//Problem 7: Inscribed Square
func inscribedSquare() *Statement {
	// the last two unnecessary lines are removed
	problem := geom.NewBoard()
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
	return NewStatement(problem, target, "OOIII", "1.7 Inscribed Square")
}
