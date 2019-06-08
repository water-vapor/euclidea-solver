// Package problem provides a problem description interface and all problems
package problem

import (
	"fmt"
	"github.com/water-vapor/euclidea-solver/pkg/geom"
)

// Statement is a type to describe a problem. Geometry objects in Board is given, objects in Target is the goal.
// Sequences is the given official hint, name is the problem name
// O: Circle
// I: Line
// .: Intersection (should not appear!)
// A: Angle bisector from three points
// +: Segment bisector
// L: Tangent line of a line through a point
// Z: Parallel line of a line through a point
// Q: Copy a segment as radius and create a circle
type Statement struct {
	Board     *geom.Board
	Target    *geom.Target
	Sequences map[string]string
	Name      string
	Goal      string
}

// NewStatement returns a new problem statement
func NewStatement(board *geom.Board, target *geom.Target, sequences map[string]string, name string) *Statement {
	return &Statement{board, target, sequences, name, ""}
}

// GetSequenceByGoal returns a hint sequence by goal name
func (s *Statement) GetSequenceByGoal() string {
	result, exists := s.Sequences[s.Goal]
	if exists {
		return result
	}
	fmt.Println("Invalid Goal specification:", s.Goal, ", using default E.")
	return s.Sequences["E"]
}

// GetProblemByID returns the construction of a problem with given ID
func GetProblemByID(chapter, number int, goal string) *Statement {
	problemID := 100*chapter + number
	var statement *Statement
	switch problemID {
	case 101:
		statement = angelOf60Degree()
	case 102:
		statement = perpendicularBisector()
	case 103:
		statement = midpoint()
	case 104:
		statement = circleInSquare()
	case 105:
		statement = rhombusInRectangle()
	case 106:
		statement = circleCenter()
	case 107:
		statement = inscribedSquare()
	case 201:
		statement = angleBisector()
	case 202:
		statement = intersectionOfAngleBisectors()
	case 203:
		statement = angleOf30Degree()
	case 204:
		statement = doubleAngle()
	case 205:
		statement = cutRectangle()
	case 206:
		statement = dropAPerpendicular()
	case 207:
		statement = erectAPerpendicular()
	case 208:
		statement = tangentToCircleAtPoint()
	case 209:
		statement = circleTangentToLine()
	case 210:
		statement = circleInRhombus()
	case 1012:
		statement = centerOfRotation()
	case 1508:
		statement = lineCircleIntersection2()
	case 1510:
		statement = angelOf3Degree()
	default:
		panic("Invalid Problem ID, or Problem not implemented.")
	}
	statement.Goal = goal
	return statement
}
