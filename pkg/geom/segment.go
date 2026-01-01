package geom

import (
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
	"math"
	"math/rand"
)

// Segment is uniquely determined by its sorted endpoints
type Segment struct {
	hashset.Serializable
	point1, point2 *Point
	hash           int64 // cached hash value
}

// computeSegmentHash pre-computes the hash for a segment
func computeSegmentHash(pt1, pt2 *Point) int64 {
	cx1 := int64(math.Round(pt1.x * configs.HashPrecision))
	cy1 := int64(math.Round(pt1.y * configs.HashPrecision))
	cx2 := int64(math.Round(pt2.x * configs.HashPrecision))
	cy2 := int64(math.Round(pt2.y * configs.HashPrecision))
	return ((cx1*configs.Prime+cy1)*configs.Prime+cx2)*configs.Prime + cy2
}

// NewSegment creates a segment from two points
func NewSegment(pt1, pt2 *Point) *Segment {
	pt1First := false
	if pt1.Equal(pt2) {
		panic("Two points of segment is same!")
	}
	// x equal
	if math.Abs(pt1.x-pt2.x) < configs.Tolerance {
		if pt1.y < pt2.y {
			pt1First = true
		}
	}
	if pt1.x < pt2.x {
		pt1First = true
	}
	if pt1First {
		return &Segment{point1: pt1, point2: pt2, hash: computeSegmentHash(pt1, pt2)}
	}
	return &Segment{point1: pt2, point2: pt1, hash: computeSegmentHash(pt2, pt1)}

}

// NewSegmentFromDirection creates a segment with one point, a direction and length
func NewSegmentFromDirection(start *Point, direction *Vector2D, length float64) *Segment {
	direction.SetLength(length)
	pt2 := NewPoint(start.x+direction.x, start.y+direction.y)
	return NewSegment(start, pt2)
}

// GetEndPoints returns both end points of a segment
func (s *Segment) GetEndPoints() (*Point, *Point) {
	return s.point1, s.point2
}

// Serialize returns the cached hash of the segment
func (s *Segment) Serialize() interface{} {
	return s.hash
}

// PointInRange checks whether a point is in the coordinates range of the segment
func (s *Segment) PointInRange(pt *Point) bool {
	// range based test
	if pt.x < s.point1.x-configs.Tolerance || pt.x > s.point2.x+configs.Tolerance {
		return false
	}
	// test on y range only if line is vertical
	if math.Abs(s.point1.x-s.point2.x) < configs.Tolerance {
		if pt.y < s.point1.y-configs.Tolerance || pt.y > s.point2.y+configs.Tolerance {
			return false
		}
	}
	return true
}

// ContainsPoint checks if a point is on the segment
func (s *Segment) ContainsPoint(pt *Point) bool {
	if !s.PointInRange(pt) {
		return false
	}
	return NewLineFromSegment(s).ContainsPoint(pt)
}

// IntersectLine returns intersections with a line
func (s *Segment) IntersectLine(l *Line) *Intersection {
	return l.IntersectSegment(s)
}

// IntersectHalfLine returns intersections with a half line
func (s *Segment) IntersectHalfLine(h *HalfLine) *Intersection {
	return h.IntersectSegment(s)
}

// IntersectSegment returns intersections with a segment
func (s *Segment) IntersectSegment(s2 *Segment) *Intersection {
	// intersect as if it is a line
	intersection := s.IntersectLine(NewLineFromSegment(s2))
	// parallel, no solution, just return
	if intersection.SolutionNumber == 0 {
		return intersection
	}
	// checks whether the intersection is on s2, since we've extended it
	pt := intersection.Solutions[0]
	if pt.OnSegment(s2) {
		return intersection
	}
	return NewIntersection()
}

// IntersectCircle returns intersections with a circle
func (s *Segment) IntersectCircle(c *Circle) *Intersection {
	return c.IntersectSegment(s)
}

// Length returns the length of the segment
func (s *Segment) Length() float64 {
	dx := s.point2.x - s.point1.x
	dy := s.point2.y - s.point1.y
	return math.Sqrt(dx*dx + dy*dy)
}

// Bisector returns a line as the bisector of the segment
func (s *Segment) Bisector() *Line {
	pt := NewPoint((s.point1.x+s.point2.x)/2, (s.point1.y+s.point2.y)/2)
	v := NewVector2DFromTwoPoints(s.point1, s.point2).NormalVector()
	return NewLineFromDirection(pt, v)
}

// GetRandomPoint returns a random point on the segment
func (s *Segment) GetRandomPoint() *Point {
	v := NewVector2DFromTwoPoints(s.point1, s.point2)
	v.SetLength(rand.Float64() * v.Length())
	return NewPoint(s.point1.x+v.x, s.point1.y+v.y)
}
