package geom

import (
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
	"math"
	"math/rand"
)

// A circle is uniquely determined by its center point and radius
type Circle struct {
	hashset.Serializable
	center *Point
	r      float64
}

func NewCircleByPoint(center, onSide *Point) *Circle {
	return &Circle{center: center, r: NewSegment(center, onSide).Length()}
}

func NewCircleByRadius(center *Point, r float64) *Circle {
	return &Circle{center: center, r: r}
}

func (c *Circle) GetCenter() *Point {
	return c.center
}

func (c *Circle) GetRadius() float64 {
	return c.r
}

func (c *Circle) Serialize() interface{} {
	cx := int64(math.Round(c.center.x * configs.HashPrecision))
	cy := int64(math.Round(c.center.y * configs.HashPrecision))
	cr := int64(math.Round(c.r * configs.HashPrecision))
	return (cx*configs.Prime+cy)*configs.Prime + cr
}

func (c *Circle) ContainsPoint(pt *Point) bool {
	return math.Abs(NewSegment(pt, c.center).Length()-c.r) < configs.Tolerance
}

func (c *Circle) IntersectLine(l *Line) *Intersection {
	distNumer := l.a*c.center.x + l.b*c.center.y + l.c
	distDenomSquare := l.a*l.a + l.b*l.b
	dist := math.Abs(distNumer) / math.Sqrt(distDenomSquare)
	// tangent circles with tolerance
	if math.Abs(dist-c.r) < configs.Tolerance {
		v := l.GetNormalVector()
		tangentLine := NewLineFromDirection(c.center, v)
		return l.IntersectLine(tangentLine)
	}
	if dist > c.r {
		return NewIntersection()
	}
	//	{{x -> -((-b^2 x1 + a (c + b y1) + Sqrt[
	//     b^2 ((a^2 + b^2) r1^2 - (c + a x1 + b y1)^2)])/(a^2 + b^2)),
	//  y -> (-b^2 (c + a x1) + a^2 b y1 +
	//    a Sqrt[b^2 ((a^2 + b^2) r1^2 - (c + a x1 + b y1)^2)])/(
	//   b (a^2 + b^2))}, {x -> (
	//   b^2 x1 - a (c + b y1) + Sqrt[
	//    b^2 ((a^2 + b^2) r1^2 - (c + a x1 + b y1)^2)])/(a^2 + b^2),
	//  y -> (-b^2 (c + a x1) + a^2 b y1 -
	//    a Sqrt[b^2 ((a^2 + b^2) r1^2 - (c + a x1 + b y1)^2)])/(
	//   b (a^2 + b^2))}}

	det := math.Sqrt(distDenomSquare*c.r*c.r - distNumer*distNumer)
	ptxc := l.b*l.b*c.center.x - l.a*(l.c+l.b*c.center.y)
	ptyc := -l.b*(l.c+l.a*c.center.x) + l.a*l.a*c.center.y
	pt1x := (ptxc - l.b*det) / distDenomSquare
	pt2x := (ptxc + l.b*det) / distDenomSquare
	pt1y := (ptyc + l.a*det) / distDenomSquare
	pt2y := (ptyc - l.a*det) / distDenomSquare
	return NewIntersection(NewPoint(pt1x, pt1y), NewPoint(pt2x, pt2y))
}

func (c *Circle) IntersectCircle(c2 *Circle) *Intersection {
	// center same, return no intersection
	if c.center.Equal(c2.center) {
		return NewIntersection()
	}
	dist := NewSegment(c.center, c2.center).Length()
	// tangent circles with tolerance
	if math.Abs(dist-c.r-c2.r) < configs.Tolerance {
		// vector from c to c2
		v := NewVector2DFromTwoPoints(c.center, c2.center)
		v.SetLength(c.r)
		pt := NewPoint(c.center.x+v.x, c.center.y+v.y)
		return NewIntersection(pt)
	}
	// separated circles
	if dist > c.r+c2.r {
		return NewIntersection()
	}
	// one circle inside another
	if math.Abs(c.r-c2.r) > dist {
		return NewIntersection()
	}
	// Implements a nice looking formula
	//https://math.stackexchange.com/a/1367732
	R2 := dist * dist
	coeff1 := (c.r*c.r - c2.r*c2.r) / R2
	coeff2 := math.Sqrt(2*(c.r*c.r+c2.r*c2.r)/R2 - coeff1*coeff1 - 1)
	pt1x := (c.center.x+c2.center.x)/2 + (c2.center.x-c.center.x)*coeff1/2 + (c2.center.y-c.center.y)*coeff2/2
	pt2x := (c.center.x+c2.center.x)/2 + (c2.center.x-c.center.x)*coeff1/2 - (c2.center.y-c.center.y)*coeff2/2
	pt1y := (c.center.y+c2.center.y)/2 + (c2.center.y-c.center.y)*coeff1/2 - (c2.center.x-c.center.x)*coeff2/2
	pt2y := (c.center.y+c2.center.y)/2 + (c2.center.y-c.center.y)*coeff1/2 + (c2.center.x-c.center.x)*coeff2/2
	return NewIntersection(NewPoint(pt1x, pt1y), NewPoint(pt2x, pt2y))
}

func (c *Circle) IntersectHalfLine(h *HalfLine) *Intersection {
	// intersect as if it is a line
	intersection := c.IntersectLine(NewLineFromHalfLine(h))
	// based on number of Solutions...
	if intersection.SolutionNumber == 0 {
		return intersection
	}
	if intersection.SolutionNumber == 1 {
		pt := intersection.Solutions[0]
		if pt.InHalfLineRange(h) {
			return intersection
		}
		return NewIntersection()
	}
	// solution number == 2
	pt1 := intersection.Solutions[0]
	pt2 := intersection.Solutions[1]
	if pt1.InHalfLineRange(h) {
		if pt2.InHalfLineRange(h) {
			return intersection
		} else {
			return NewIntersection(pt1)
		}
	} else {
		if pt2.InHalfLineRange(h) {
			return NewIntersection(pt2)
		} else {
			return NewIntersection()
		}
	}
}

func (c *Circle) IntersectSegment(s *Segment) *Intersection {
	// intersect as if it is a line
	intersection := c.IntersectLine(NewLineFromSegment(s))
	// based on number of Solutions...
	if intersection.SolutionNumber == 0 {
		return intersection
	}
	if intersection.SolutionNumber == 1 {
		pt := intersection.Solutions[0]
		if pt.InSegmentRange(s) {
			return intersection
		}
		return NewIntersection()
	}
	// solution number == 2
	pt1 := intersection.Solutions[0]
	pt2 := intersection.Solutions[1]
	if pt1.InSegmentRange(s) {
		if pt2.InSegmentRange(s) {
			return intersection
		} else {
			return NewIntersection(pt1)
		}
	} else {
		if pt2.InSegmentRange(s) {
			return NewIntersection(pt2)
		} else {
			return NewIntersection()
		}
	}
}

// Get a random point on the circle
func (c *Circle) GetRandomPoint() *Point {
	// random number from -1 to 1
	x := rand.Float64()*2 - 1
	y := math.Sqrt(1 - x*x)
	// decide sign
	if rand.Float64() < 0.5 {
		y *= -1
	}
	v := NewVector2D(x, y)
	v.SetLength(c.r)
	return NewPoint(c.center.x+v.x, c.center.y+v.y)
}
