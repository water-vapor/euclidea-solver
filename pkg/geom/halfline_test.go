package geom

import "testing"

func TestHalfLine_ContainsPoint(t *testing.T) {
	h1 := NewHalfLineFromTwoPoints(NewPoint(0, 0), NewPoint(1, 0))
	if h1.ContainsPoint(NewPoint(-1, 0)) {
		t.Error("does not contain (-1,0)")
	}
	if !h1.ContainsPoint(NewPoint(2, 0)) {
		t.Error("contains (2,0)")
	}
	h2 := NewHalfLineFromTwoPoints(NewPoint(0, 0), NewPoint(0, 1))
	if h2.ContainsPoint(NewPoint(0, -1)) {
		t.Error("does not contain (0,-1)")
	}
	if !h2.ContainsPoint(NewPoint(0, 2)) {
		t.Error("contains (0,2)")
	}
}

func TestHalfLine_GetRandomPoint(t *testing.T) {
	h1 := NewHalfLineFromTwoPoints(NewPoint(123, 456), NewPoint(789, 123))
	pt := h1.GetRandomPoint()
	if !pt.OnHalfLine(h1) {
		t.Error("Point should on halfline")
	}
}
