package problem

import "github.com/water-vapor/euclidea-solver/pkg/geom"

type Statement struct {
	Board    *geom.Board
	Target   *geom.Goal
	Sequence string
	Name     string
}

func NewStatement(board *geom.Board, target *geom.Goal, sequence string, name string) *Statement {
	return &Statement{board, target, sequence, name}
}

func GetProblemByID(chapter, number int) *Statement {
	problemID := 100*chapter + number
	switch problemID {
	case 101:
		return angelOf60Degree()
	case 1012:
		return centerOfRotation()
	default:
		panic("Invalid Problem ID, or Problem not implemented.")
	}
}
