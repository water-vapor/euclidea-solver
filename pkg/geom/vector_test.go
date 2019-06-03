package geom

import (
	"github.com/water-vapor/euclidea-solver/configs"
	"math"
	"testing"
)

func TestVector2D_Normalize(t *testing.T) {
	v := NewVector2D(10, 0)
	v.Normalize()
	if v.x != 1 {
		t.Errorf("Vector normalize: x=1, got %f", v.x)
	}
	v = NewVector2D(-2, -1)
	v.Normalize()
	if math.Abs(v.x*v.x+v.y*v.y-1) > configs.Tolerance {
		t.Errorf("Vector normalize: %f %f %f", v.x, v.y, v.x*v.x+v.y*v.y)
	}
}

func TestVector2D_SetLength(t *testing.T) {
	v := NewVector2D(0, 5)
	v.SetLength(10)
	if v.y != 10 || v.x != 0 {
		t.Errorf("Vector normalize: (0,10), got (%f,%f)", v.x, v.y)
	}
}
