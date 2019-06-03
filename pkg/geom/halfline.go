package geom

import (
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
	"math"
	"math/rand"
)

// A half line is determined uniquely by its starting point and normalized direction vector
type HalfLine struct {
	hashset.Serializable
	point     *Point
	direction *Vector2D
}

func NewHalfLineFromTwoPoints(source *Point, direction *Point) *HalfLine {
	v := NewVector2D(direction.x-source.x, direction.y-source.y)
	v.Normalize()
	return &HalfLine{point: source, direction: v}
}

func NewHalfLineFromDirection(pt *Point, direction *Vector2D) *HalfLine {
	direction.Normalize()
	return &HalfLine{point: pt, direction: direction}
}

func (h *HalfLine) GetEndPoint() *Point {
	return h.point
}

func (h *HalfLine) Serialize() interface{} {
	cx := int64(math.Round(h.direction.x * configs.HashPrecision))
	cy := int64(math.Round(h.direction.y * configs.HashPrecision))
	cx0 := int64(math.Round(h.point.x * configs.HashPrecision))
	cy0 := int64(math.Round(h.point.y * configs.HashPrecision))
	return ((cx*configs.Prime+cy)*configs.Prime+cx0)*configs.Prime + cy0
}

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

func (h *HalfLine) ContainsPoint(pt *Point) bool {
	if !h.PointInRange(pt) {
		return false
	}
	return NewLineFromHalfLine(h).ContainsPoint(pt)
}

func (h *HalfLine) IntersectLine(l *Line) *Intersection {
	return l.IntersectHalfLine(h)
}

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

func (h *HalfLine) IntersectCircle(c *Circle) *Intersection {
	return c.IntersectHalfLine(h)
}

func (h *HalfLine) GetRandomPoint() *Point {
	v := h.direction.Clone()
	v.SetLength(rand.Float64() * configs.RandomPointRange)
	return NewPoint(h.point.x+v.x, h.point.y+v.y)
}
