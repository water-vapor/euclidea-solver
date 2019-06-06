package problem

import (
	"github.com/water-vapor/euclidea-solver/pkg/geom"
)

// Problem 12: Center of Rotation
func centerOfRotation() *Statement {
	pt1 := geom.NewPoint(-5, 0)
	pt2 := geom.NewPoint(0, 5)
	pt3 := geom.NewPoint(-2, 6)
	v := geom.NewVector2D(1, -1.234)
	s1 := geom.NewSegment(pt1, pt2)
	s2 := geom.NewSegmentFromDirection(pt3, v, s1.Length())
	_, pt4 := s2.GetEndPoints()

	bisect1 := geom.NewSegment(pt1, pt3).Bisector()
	bisect2 := geom.NewSegment(pt2, pt4).Bisector()
	result := bisect1.IntersectLine(bisect2).Solutions[0]

	board := geom.NewBoard()
	board.AddSegment(s1)
	board.AddSegment(s2)

	target := geom.NewGoal()
	target.Points.Add(result)

	return NewStatement(board, target, "OOIOI", "10.12 Center Of Rotation")
}
