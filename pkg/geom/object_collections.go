package geom

import (
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
)

// Intersection is a type that stores the resulting intersections of two geometry objects
type Intersection struct {
	SolutionNumber int
	Solutions      []*Point
}

// NewIntersection return a new Intersection object
func NewIntersection(pts ...*Point) *Intersection {
	return &Intersection{len(pts), pts}
}

// Goal is a type that holds the goal of a problem, it can only include Points, Lines and Circles
type Goal struct {
	Points  *hashset.HashSet
	Lines   *hashset.HashSet
	Circles *hashset.HashSet
}

// NewGoal return a new Goal object
func NewGoal() *Goal {
	return &Goal{hashset.NewHashSet(), hashset.NewHashSet(), hashset.NewHashSet()}
}
