package geom

import "math"

// Vector2D is a 2D vector object
type Vector2D struct {
	x, y float64
}

// NewVector2D creates a 2D vector from its components
func NewVector2D(x, y float64) *Vector2D {
	return &Vector2D{x: x, y: y}
}

// NewVector2DFromTwoPoints creates a 2D vector from two points, pt1 to pt2
func NewVector2DFromTwoPoints(pt1, pt2 *Point) *Vector2D {
	return &Vector2D{x: pt2.x - pt1.x, y: pt2.y - pt1.y}
}

// Clone returns a copy of the vector, since a vector can be modified
func (v *Vector2D) Clone() *Vector2D {
	return NewVector2D(v.x, v.y)
}

// Length returns the norm of a vector
func (v *Vector2D) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Normalize sets the length of vector to 1, reserving its direction
func (v *Vector2D) Normalize() {
	v.SetLength(1)
}

// SetLength sets the length of vector to a certain value
func (v *Vector2D) SetLength(l float64) {
	coeff := l / math.Sqrt(v.x*v.x+v.y*v.y)
	v.x *= coeff
	v.y *= coeff
}

// NormalVector gets a new normal vector by rotating Pi/2 counter clockwise
func (v *Vector2D) NormalVector() *Vector2D {
	return NewVector2D(-v.y, v.x)
}
