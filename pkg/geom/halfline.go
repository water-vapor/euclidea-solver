package geom

import (
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
	"math"
	"math/rand"
)

// HalfLine is determined uniquely by its starting point and normalized direction vector
type HalfLine struct {
	hashset.Serializable
	point     *Point
	direction *Vector2D
	hash      int64 // cached hash value
}

// computeHalfLineHash pre-computes the hash for a half line
func computeHalfLineHash(pt *Point, dir *Vector2D) int64 {
	cx := int64(math.Round(dir.x * configs.HashPrecision))
	cy := int64(math.Round(dir.y * configs.HashPrecision))
	cx0 := int64(math.Round(pt.x * configs.HashPrecision))
	cy0 := int64(math.Round(pt.y * configs.HashPrecision))
	return ((cx*configs.Prime+cy)*configs.Prime+cx0)*configs.Prime + cy0
}

// NewHalfLineFromTwoPoints creates a half line from two points, with source as end point
func NewHalfLineFromTwoPoints(source *Point, direction *Point) *HalfLine {
	v := NewVector2D(direction.x-source.x, direction.y-source.y)
	v.Normalize()
	return &HalfLine{point: source, direction: v, hash: computeHalfLineHash(source, v)}
}

// NewHalfLineFromDirection creates a half line from its end point and a direction
func NewHalfLineFromDirection(pt *Point, direction *Vector2D) *HalfLine {
	direction.Normalize()
	return &HalfLine{point: pt, direction: direction, hash: computeHalfLineHash(pt, direction)}
}

// GetEndPoint returns its end point
func (h *HalfLine) GetEndPoint() *Point {
	return h.point
}

// Serialize returns the cached hash of the half line
func (h *HalfLine) Serialize() interface{} {
	return h.hash
}

// PointInRange checks if a point is in the coordinates of a half line
func (h *HalfLine) PointInRange(pt *Point) bool {
	if math.Abs(h.direction.x) < configs.Tolerance {
		// line is vertical
		if h.direction.y < 0 && pt.y-h.point.y-configs.Tolerance > 0 {
			return false
		}
		if h.direction.y > 0 && pt.y-h.point.y+configs.Tolerance < 0 {
			return false
		}
	} else {
		if h.direction.x < 0 && pt.x-h.point.x-configs.Tolerance > 0 {
			return false
		}
		if h.direction.x > 0 && pt.x-h.point.x+configs.Tolerance < 0 {
			return false
		}
	}
	return true
}

// ContainsPoint checks if a point is on the half line
func (h *HalfLine) ContainsPoint(pt *Point) bool {
	if !h.PointInRange(pt) {
		return false
	}
	return NewLineFromHalfLine(h).ContainsPoint(pt)
}

// IntersectLine returns intersections with a line
func (h *HalfLine) IntersectLine(l *Line) *Intersection {
	return l.IntersectHalfLine(h)
}

// IntersectHalfLine returns intersections with another half line
func (h *HalfLine) IntersectHalfLine(h2 *HalfLine) *Intersection {
	// intersect as if it is a line
	intersection := h.IntersectLine(NewLineFromHalfLine(h2))
	// parallel, no solution, just return
	if intersection.SolutionNumber == 0 {
		return intersection
	}
	// checks whether the intersection is on h2, since we've extended it
	pt := intersection.Solutions[0]
	if pt.OnHalfLine(h2) {
		return intersection
	}
	return NewIntersection()
}

// IntersectSegment returns intersections with a segment
func (h *HalfLine) IntersectSegment(s *Segment) *Intersection {
	// intersect as if it is a line
	intersection := h.IntersectLine(NewLineFromSegment(s))
	// parallel, no solution, just return
	if intersection.SolutionNumber == 0 {
		return intersection
	}
	// checks whether the intersection is on s, since we've extended it
	pt := intersection.Solutions[0]
	if pt.OnSegment(s) {
		return intersection
	}
	return NewIntersection()
}

// IntersectCircle returns intersections with a circle
func (h *HalfLine) IntersectCircle(c *Circle) *Intersection {
	return c.IntersectHalfLine(h)
}

// GetRandomPoint returns a random point on the half line
func (h *HalfLine) GetRandomPoint() *Point {
	v := h.direction.Clone()
	v.SetLength(rand.Float64() * configs.RandomPointRange)
	return NewPoint(h.point.x+v.x, h.point.y+v.y)
}
