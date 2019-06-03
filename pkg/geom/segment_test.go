package geom

import "testing"

func TestSegment_ContainsPoint(t *testing.T) {
	s1 := NewSegment(NewPoint(0, 0), NewPoint(1, 0))
	if s1.ContainsPoint(NewPoint(-1, 0)) {
		t.Error("does not contain (-1,0)")
	}
	if s1.ContainsPoint(NewPoint(2, 0)) {
		t.Error("does not contain (2,0)")
	}
	if !s1.ContainsPoint(NewPoint(0, 0)) {
		t.Error("contains (0,0)")
	}
	s2 := NewSegment(NewPoint(0, 0), NewPoint(0, 1))
	if s2.ContainsPoint(NewPoint(0, -1)) {
		t.Error("does not contain (0,-1)")
	}
	if s2.ContainsPoint(NewPoint(0, 2)) {
		t.Error("does not contain (0,2)")
	}
	if !s2.ContainsPoint(NewPoint(0, 0)) {
		t.Error("contains (0,0)")
	}
}

func TestSegment_Bisector(t *testing.T) {
	s1 := NewSegment(NewPoint(-1, 0), NewPoint(1, 0))
	l1 := NewLineFromTwoPoints(NewPoint(0, 0), NewPoint(0, 1))
	if s1.Bisector().Serialize() != l1.Serialize() {
		t.Error("")
	}
}

func TestSegment_GetRandomPoint(t *testing.T) {
	s1 := NewSegment(NewPoint(12, 34), NewPoint(56, 78))
	pt := s1.GetRandomPoint()
	if !pt.OnSegment(s1) {
		t.Error("Point should on segment")
	}
}
