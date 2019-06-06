package problem

import "github.com/water-vapor/euclidea-solver/pkg/geom"

// Statement is a type to describe a problem. Geometry objects in Board is given, objects in Target is the goal.
// Sequence is the given official hint, name is the problem name
// O: Circle
// I: Line
// .: Intersection (should not appear!)
// A: Angle bisector from three points
// +: Segment bisector
// L: Tangent line of a line through a point
// Z: Parallel line of a line through a point
// Q: Copy a segment as radius and create a circle
type Statement struct {
	Board    *geom.Board
	Target   *geom.Goal
	Sequence string
	Name     string
}

// NewStatement returns a new problem statement
func NewStatement(board *geom.Board, target *geom.Goal, sequence string, name string) *Statement {
	return &Statement{board, target, sequence, name}
}

// GetProblemByID returns the construction of a problem with given ID
func GetProblemByID(chapter, number int) *Statement {
	problemID := 100*chapter + number
	switch problemID {
	case 101:
		return angelOf60Degree()
	case 102:
		return perpendicularBisector()
	case 103:
		return midpoint()
	case 104:
		return circleInSquare()
	case 105:
		return rhombusInRectangle()
	case 106:
		return circleCenter()
	case 107:
		return inscribedSquare()
	case 1012:
		return centerOfRotation()
	case 1508:
		return lineCircleIntersection2()
	case 1510:
		return angelOf3Degree()
	default:
		panic("Invalid Problem ID, or Problem not implemented.")
	}
}
