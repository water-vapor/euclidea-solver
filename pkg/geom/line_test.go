package geom

import "testing"

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
