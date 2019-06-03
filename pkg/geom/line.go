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

// A line is represented as ax+by+c==0, it is unique when the max coefficient is
// normalized to a fixed number, and a being non-negative.
type Line struct {
	hashset.Serializable
	a, b, c float64
}

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

func NewLineFromTwoPoints(pt1, pt2 *Point) *Line {
	a := pt2.y - pt1.y
	b := pt1.x - pt2.x
	c := pt2.x*pt1.y - pt1.x*pt2.y
	return NewLineFromCoefficients(a, b, c)
}

func NewLineFromSegment(s *Segment) *Line {
	return NewLineFromTwoPoints(s.point1, s.point2)
}

func NewLineFromDirection(pt *Point, v *Vector2D) *Line {
	vn := v.NormalVector()
	return NewLineFromCoefficients(vn.x, vn.y, -vn.x*pt.x-vn.y*pt.y)
}

func NewLineFromHalfLine(h *HalfLine) *Line {
	return NewLineFromDirection(h.point, h.direction)
}

func (l *Line) Serialize() interface{} {
	ca := int64(math.Round(l.a))
	cb := int64(math.Round(l.b))
	cc := int64(math.Round(l.c))
	return (ca*configs.Prime+cb)*configs.Prime + cc
}

func (l *Line) ContainsPoint(pt *Point) bool {
	return math.Abs(l.a*pt.x+l.b*pt.y+l.c) < lineThres
}

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

func (l *Line) GetNormalVector() *Vector2D {
	return NewVector2D(l.a, l.b)
}

func (l *Line) IntersectCircle(c *Circle) *Intersection {
	return c.IntersectLine(l)
}

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
