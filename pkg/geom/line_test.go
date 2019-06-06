package geom

import (
	"testing"
)

func TestNewLineFromCoefficients(t *testing.T) {
	l := NewLineFromCoefficients(1, 1, 1)
	pt1 := NewPoint(-1, 0)
	pt2 := NewPoint(0, -1)
	if !(l.ContainsPoint(pt1) && l.ContainsPoint(pt2)) {
		t.Error("Fail")
	}
}

func TestNewLineFromTwoPoints(t *testing.T) {
	l := NewLineFromTwoPoints(NewPoint(1, 0), NewPoint(0, 1))
	if !(l.ContainsPoint(NewPoint(2, -1)) && l.ContainsPoint(NewPoint(-2, 3))) {
		t.Errorf("Fail, %f x+%f y+%f==0", l.a, l.b, l.c)
	}
}

func TestNewLineAsAngleBisector(t *testing.T) {
	l1 := NewLineAsAngleBisector(NewPoint(1, 0), NewPoint(0, 0), NewPoint(0, 1))
	l2 := NewLineAsAngleBisector(NewPoint(-9, 0), NewPoint(0, 0), NewPoint(0, -8))
	if l1.Serialize() != l2.Serialize() || l1.Serialize() != NewLineFromTwoPoints(NewPoint(1, 1), NewPoint(0, 0)).Serialize() {
		t.Error("angle bisector error")
	}
	l3 := NewLineAsAngleBisector(NewPoint(-2, 2), NewPoint(0, 1), NewPoint(4, 3))
	if l3.Serialize() != NewLineFromTwoPoints(NewPoint(0, 0), NewPoint(0, 1)).Serialize() {
		t.Error("angle bisector error")
	}

}

func TestLine_IntersectLine(t *testing.T) {
	l1 := NewLineFromTwoPoints(NewPoint(1, 0), NewPoint(1.00000001, 1234))
	l2 := NewLineFromTwoPoints(NewPoint(1, 0), NewPoint(-123, 456))
	inters := l1.IntersectLine(l2)
	pt := inters.Solutions[0]
	if pt.Serialize() != NewPoint(1, 0).Serialize() {
		t.Errorf("intersection point wrong, got (%f, %f)", pt.x, pt.y)
	}
}

func TestLine_IntersectHalfLine(t *testing.T) {
	l1 := NewLineFromTwoPoints(NewPoint(1, 0), NewPoint(1.00000001, 1234))
	l2 := NewHalfLineFromTwoPoints(NewPoint(0.9999, 0), NewPoint(-123, 456))
	inters := l1.IntersectHalfLine(l2)
	if inters.SolutionNumber != 0 {
		t.Error("solution number wrong")
	}
}

func TestLine_GetRandomPoint(t *testing.T) {
	l1 := NewLineFromCoefficients(1, 2, 3)
	l2 := NewLineFromCoefficients(10, 0, 4)
	l3 := NewLineFromCoefficients(0, -4, -2)
	pt1 := l1.GetRandomPoint()
	pt2 := l2.GetRandomPoint()
	pt3 := l3.GetRandomPoint()
	if !pt1.OnLine(l1) {
		t.Error("Point should on line")
	}
	if !pt2.OnLine(l2) {
		t.Error("Point should on line")
	}
	if !pt3.OnLine(l3) {
		t.Error("Point should on line")
	}
}

func TestLine_GetParallelLineWithPoint(t *testing.T) {
	l1 := NewLineFromTwoPoints(NewPoint(0, 0), NewPoint(1, 0))
	t1 := NewLineFromTwoPoints(NewPoint(0, 1), NewPoint(1, 1))
	p1 := l1.GetParallelLineWithPoint(NewPoint(-1, 1))
	if t1.Serialize() != p1.Serialize() {
		t.Error("y=1 should be parallel with y=0")
	}
	l2 := NewLineFromTwoPoints(NewPoint(0, 0), NewPoint(0, 1))
	t2 := NewLineFromTwoPoints(NewPoint(1, 1), NewPoint(1, 0))
	p2 := l2.GetParallelLineWithPoint(NewPoint(1, -1))
	if t2.Serialize() != p2.Serialize() {
		t.Error("x=1 should be parallel with x=0")
	}
}

func TestLine_GetTangentLineWithPoint(t *testing.T) {
	l1 := NewLineFromTwoPoints(NewPoint(0, 0), NewPoint(1, 0))
	t1 := NewLineFromTwoPoints(NewPoint(0, 0), NewPoint(0, 1))
	p1 := l1.GetTangentLineWithPoint(NewPoint(0, 0))
	if t1.Serialize() != p1.Serialize() {
		t.Error("x=0 should be tangent with y=0")
	}
	p2 := l1.GetTangentLineWithPoint(NewPoint(0, -1))
	if t1.Serialize() != p2.Serialize() {
		t.Error("x=0 should be tangent with y=0")
	}
}
