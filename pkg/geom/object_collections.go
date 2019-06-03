package geom

import (
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
)

// A type that stores the resulting intersections of two geometry objects
type Intersection struct {
	SolutionNumber int
	Solutions      []*Point
}

func NewIntersection(pts ...*Point) *Intersection {
	return &Intersection{len(pts), pts}
}

// A type that holds the goal of a problem, it can only include Points, Lines and Circles
type Goal struct {
	Points  *hashset.HashSet
	Lines   *hashset.HashSet
	Circles *hashset.HashSet
}

func NewGoal() *Goal {
	return &Goal{hashset.NewHashSet(), hashset.NewHashSet(), hashset.NewHashSet()}
}
