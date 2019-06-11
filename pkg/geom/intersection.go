package geom

// Intersection is a type that stores the resulting intersections of two geometry objects
type Intersection struct {
	SolutionNumber int
	Solutions      []*Point
}

// NewIntersection return a new Intersection object
func NewIntersection(pts ...*Point) *Intersection {
	return &Intersection{len(pts), pts}
}
