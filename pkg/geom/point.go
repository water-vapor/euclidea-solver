package geom

import (
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
	"math"
)

// Point is a object containing its two coordinates
type Point struct {
	hashset.Serializable
	x, y float64
}

// NewPoint creates a point from coordinates
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// GetCoords returns its coordinates, should be only used for debugging
func (pt *Point) GetCoords() (float64, float64) {
	return pt.x, pt.y
}

// Serialize returns the hash of a point
func (pt *Point) Serialize() interface{} {
	ptx := int64(math.Round(pt.x * configs.HashPrecision))
	pty := int64(math.Round(pt.y * configs.HashPrecision))
	return ptx*configs.Prime + pty
}

// OnLine checks if the point is on a line
func (pt *Point) OnLine(l *Line) bool {
	return l.ContainsPoint(pt)
}

// OnHalfLine checks if the point is on a half line
func (pt *Point) OnHalfLine(h *HalfLine) bool {
	return h.ContainsPoint(pt)
}

// InHalfLineRange checks if the point in the coordinate range of a half line.
// This function should only be used when the point is guaranteed to be on the line which the half line belong to
func (pt *Point) InHalfLineRange(h *HalfLine) bool {
	return h.PointInRange(pt)
}

// OnSegment checks if the point is on a segment
func (pt *Point) OnSegment(s *Segment) bool {
	return s.ContainsPoint(pt)
}

// InSegmentRange checks if the point in the coordinate range of a segment.
// This function should only be used when the point is guaranteed to be on the line which the segment belong to
func (pt *Point) InSegmentRange(s *Segment) bool {
	return s.PointInRange(pt)
}

// OnCircle checks if the point is on a circle
func (pt *Point) OnCircle(c *Circle) bool {
	return c.ContainsPoint(pt)
}

// DistanceToLine calculates the distance from the point to a line
func (pt *Point) DistanceToLine(l *Line) float64 {
	return math.Abs(l.a*pt.x+l.b*pt.y+l.c) / math.Sqrt(l.a*l.a+l.b*l.b)
}

// Equal checks equality of two points with tolerance
func (pt *Point) Equal(pt2 *Point) bool {
	return math.Abs(pt.x-pt2.x) < configs.Tolerance && math.Abs(pt.y-pt2.y) < configs.Tolerance
}
