package geom

import "math"

type Vector2D struct {
	x, y float64
}

func NewVector2D(x, y float64) *Vector2D {
	return &Vector2D{x: x, y: y}
}

func NewVector2DFromTwoPoints(pt1, pt2 *Point) *Vector2D {
	return &Vector2D{x: pt2.x - pt1.x, y: pt2.y - pt1.y}
}

// a vector can be modified, so a copy constructor is provided
func (v *Vector2D) Clone() *Vector2D {
	return NewVector2D(v.x, v.y)
}

func (v *Vector2D) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vector2D) Normalize() {
	v.SetLength(1)
}

func (v *Vector2D) SetLength(l float64) {
	coeff := l / math.Sqrt(v.x*v.x+v.y*v.y)
	v.x *= coeff
	v.y *= coeff
}

// Get a new normal vector by rotating Pi/2 counter clockwise
func (v *Vector2D) NormalVector() *Vector2D {
	return NewVector2D(-v.y, v.x)
}
