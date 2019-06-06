package geom

import (
	"fmt"
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
	"math"
	"math/rand"
)

const (
	maxCoeff  = 10000000
	lineThres = 0.1
)

// Line is a line represented as ax+by+c==0, it is unique when the max coefficient is
// normalized to a fixed number, and a being non-negative.
type Line struct {
	hashset.Serializable
	a, b, c float64
}

// NewLineFromCoefficients creates a line from coefficients
func NewLineFromCoefficients(a, b, c float64) *Line {
	// a should be positive
	if a < 0 {
		a *= -1
		b *= -1
		c *= -1
	} else {
		if a == 0 && b < 0 {
			b *= -1
			c *= -1
		}
	}
	max := math.Max(math.Max(a, math.Abs(b)), math.Abs(c))
	coeff := maxCoeff / max
	a *= coeff
	b *= coeff
	c *= coeff
	return &Line{a: a, b: b, c: c}
}

// NewLineFromTwoPoints creates a line from two points
func NewLineFromTwoPoints(pt1, pt2 *Point) *Line {
	a := pt2.y - pt1.y
	b := pt1.x - pt2.x
	c := pt2.x*pt1.y - pt1.x*pt2.y
	return NewLineFromCoefficients(a, b, c)
}

// NewLineFromSegment creates a line from a segment
func NewLineFromSegment(s *Segment) *Line {
	return NewLineFromTwoPoints(s.point1, s.point2)
}

// NewLineFromDirection creates a line from a point and a vector as direction
func NewLineFromDirection(pt *Point, v *Vector2D) *Line {
	vn := v.NormalVector()
	return NewLineFromCoefficients(vn.x, vn.y, -vn.x*pt.x-vn.y*pt.y)
}

// NewLineFromHalfLine creates a line from a half line
func NewLineFromHalfLine(h *HalfLine) *Line {
	return NewLineFromDirection(h.point, h.direction)
}

// NewLineAsAngleBisector returns a line as angle bisector of pt1,pt2,pt3
func NewLineAsAngleBisector(pt1, pt2, pt3 *Point) *Line {
	s1 := NewSegment(pt1, pt2)
	s2 := NewSegment(pt2, pt3)
	d1 := s1.Length()
	d2 := s2.Length()
	inters1 := pt1
	if d1 > d2 {
		d1 = d2
		s1, s2 = s2, s1
		inters1 = pt3
	}
	c := NewCircleByRadius(pt2, d1)
	inters2 := c.IntersectSegment(s2).Solutions[0]
	return NewSegment(inters1, inters2).Bisector()
}

// Serialize returns the hash of a line
func (l *Line) Serialize() interface{} {
	ca := int64(math.Round(l.a))
	cb := int64(math.Round(l.b))
	cc := int64(math.Round(l.c))
	return (ca*configs.Prime+cb)*configs.Prime + cc
}

// ContainsPoint checks whether a point is on the line
func (l *Line) ContainsPoint(pt *Point) bool {
	return math.Abs(l.a*pt.x+l.b*pt.y+l.c) < lineThres
}

// IntersectLine returns intersections with another line
func (l *Line) IntersectLine(l2 *Line) *Intersection {
	// parallel
	if l.a == l2.a && l.b == l2.b {
		return NewIntersection()
	}
	//x -> -((-b2 c + b c2)/(a2 b - a b2)), y -> -((a2 c - a c2)/(
	//  a2 b - a b2))
	denom := l.a*l2.b - l2.a*l.b
	x := (-l2.b*l.c + l.b*l2.c) / denom
	y := (l2.a*l.c - l.a*l2.c) / denom
	return NewIntersection(NewPoint(x, y))
}

// IntersectHalfLine returns intersections with a half line
func (l *Line) IntersectHalfLine(h *HalfLine) *Intersection {
	// intersect as if it is a line
	intersection := l.IntersectLine(NewLineFromHalfLine(h))
	// parallel, no solution, just return
	if intersection.SolutionNumber == 0 {
		return intersection
	}
	// checks whether the intersection is on the halfline
	pt := intersection.Solutions[0]
	if pt.InHalfLineRange(h) {
		return intersection
	}
	return NewIntersection()
}

// IntersectSegment returns intersections with a segment
func (l *Line) IntersectSegment(s *Segment) *Intersection {
	// intersect as if it is a line
	intersection := l.IntersectLine(NewLineFromSegment(s))
	// parallel, no solution, just return
	if intersection.SolutionNumber == 0 {
		return intersection
	}
	// checks whether the intersection is on the segment
	pt := intersection.Solutions[0]
	if pt.InSegmentRange(s) {
		return intersection
	}
	return NewIntersection()
}

// GetNormalVector returns a normal vector of the line, length not specified
func (l *Line) GetNormalVector() *Vector2D {
	return NewVector2D(l.a, l.b)
}

// GetParallelVector returns a parallel vector of the line, length not specified
func (l *Line) GetParallelVector() *Vector2D {
	return NewVector2D(-l.b, l.a)
}

// GetTangentLineWithPoint returns a tangent line through a point
func (l *Line) GetTangentLineWithPoint(pt *Point) *Line {
	return NewLineFromDirection(pt, l.GetNormalVector())
}

// GetParallelLineWithPoint returns a parallel line through a point
func (l *Line) GetParallelLineWithPoint(pt *Point) *Line {
	return NewLineFromDirection(pt, l.GetParallelVector())
}

// IntersectCircle returns intersections with a circle
func (l *Line) IntersectCircle(c *Circle) *Intersection {
	return c.IntersectLine(l)
}

// GetRandomPoint returns a random point on the line
func (l *Line) GetRandomPoint() *Point {
	// not a vertical line
	if math.Abs(l.b) > 1 {
		x := rand.Float64()*configs.RandomPointRange*2 - configs.RandomPointRange
		y := -(l.a*x + l.c) / l.b
		if math.Abs(y) < configs.MaxPointCoord {
			return NewPoint(x, y)
		}
	} else {
		y := rand.Float64()*configs.RandomPointRange*2 - configs.RandomPointRange
		x := -(l.b*y + l.c) / l.a
		if math.Abs(x) < configs.MaxPointCoord {
			return NewPoint(x, y)
		}
	}
	fmt.Println("Line::GetRandomPoint: This line is out of range.")
	return NewPoint(-l.c/l.a, 0)

}
