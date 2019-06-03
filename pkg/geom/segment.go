package geom

import (
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
	"math"
	"math/rand"
)

// A segment is uniquely determined by its sorted endpoints
type Segment struct {
	hashset.Serializable
	point1, point2 *Point
}

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
		return &Segment{point1: pt1, point2: pt2}
	} else {
		return &Segment{point1: pt2, point2: pt1}
	}
}

func NewSegmentFromDirection(start *Point, direction *Vector2D, length float64) *Segment {
	direction.SetLength(length)
	pt2 := NewPoint(start.x+direction.x, start.y+direction.y)
	return NewSegment(start, pt2)
}

func (s *Segment) GetEndPoints() (*Point, *Point) {
	return s.point1, s.point2
}

func (s *Segment) Serialize() interface{} {
	cx1 := int64(math.Round(s.point1.x * configs.HashPrecision))
	cy1 := int64(math.Round(s.point1.y * configs.HashPrecision))
	cx2 := int64(math.Round(s.point2.x * configs.HashPrecision))
	cy2 := int64(math.Round(s.point2.y * configs.HashPrecision))
	return ((cx1*configs.Prime+cy1)*configs.Prime+cx2)*configs.Prime + cy2
}

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

func (s *Segment) ContainsPoint(pt *Point) bool {
	if !s.PointInRange(pt) {
		return false
	}
	return NewLineFromSegment(s).ContainsPoint(pt)
}

func (s *Segment) IntersectLine(l *Line) *Intersection {
	return l.IntersectSegment(s)
}

func (s *Segment) IntersectHalfLine(h *HalfLine) *Intersection {
	return h.IntersectSegment(s)
}

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

func (s *Segment) IntersectCircle(c *Circle) *Intersection {
	return c.IntersectSegment(s)
}

func (s *Segment) Length() float64 {
	dx := s.point2.x - s.point1.x
	dy := s.point2.y - s.point1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func (s *Segment) Bisector() *Line {
	pt := NewPoint((s.point1.x+s.point2.x)/2, (s.point1.y+s.point2.y)/2)
	v := NewVector2DFromTwoPoints(s.point1, s.point2).NormalVector()
	return NewLineFromDirection(pt, v)
}

func (s *Segment) GetRandomPoint() *Point {
	v := NewVector2DFromTwoPoints(s.point1, s.point2)
	v.SetLength(rand.Float64() * v.Length())
	return NewPoint(s.point1.x+v.x, s.point1.y+v.y)
}
