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

// Target is a type that holds the target of a problem, it can only include Points, Lines and Circles
type Target struct {
	Points  *hashset.HashSet
	Lines   *hashset.HashSet
	Circles *hashset.HashSet
}

// NewTarget return a new Target object
func NewTarget() *Target {
	return &Target{hashset.NewHashSet(), hashset.NewHashSet(), hashset.NewHashSet()}
}
