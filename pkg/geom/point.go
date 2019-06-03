package geom

import (
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
	"math"
)

type Point struct {
	hashset.Serializable
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (pt *Point) GetCoords() (float64, float64) {
	return pt.x, pt.y
}

func (pt *Point) Serialize() interface{} {
	ptx := int64(math.Round(pt.x * configs.HashPrecision))
	pty := int64(math.Round(pt.y * configs.HashPrecision))
	return ptx*configs.Prime + pty
}

func (pt *Point) OnLine(l *Line) bool {
	return l.ContainsPoint(pt)
}

func (pt *Point) OnHalfLine(h *HalfLine) bool {
	return h.ContainsPoint(pt)
}

func (pt *Point) InHalfLineRange(h *HalfLine) bool {
	return h.PointInRange(pt)
}

func (pt *Point) OnSegment(s *Segment) bool {
	return s.ContainsPoint(pt)
}

func (pt *Point) InSegmentRange(s *Segment) bool {
	return s.PointInRange(pt)
}

func (pt *Point) OnCircle(c *Circle) bool {
	return c.ContainsPoint(pt)
}

func (pt *Point) DistanceToLine(l *Line) float64 {
	return math.Abs(l.a*pt.x+l.b*pt.y+l.c) / math.Sqrt(l.a*l.a+l.b*l.b)
}

func (pt *Point) Equal(pt2 *Point) bool {
	return math.Abs(pt.x-pt2.x) < configs.Tolerance && math.Abs(pt.y-pt2.y) < configs.Tolerance
}
