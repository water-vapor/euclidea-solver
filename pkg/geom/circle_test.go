package geom

import (
	"math"
	"testing"
)

func TestCircle_IntersectCircle(t *testing.T) {
	c1 := NewCircleByRadius(NewPoint(0, 0), 1)
	c2 := NewCircleByPoint(NewPoint(3, 0), NewPoint(3, 2))
	inters := c1.IntersectCircle(c2)
	if inters.SolutionNumber != 1 {
		t.Error("number of solution wrong")
	}
	pt := inters.Solutions[0]
	if pt.Serialize() != NewPoint(1, 0).Serialize() {
		t.Errorf("intersection point wrong, should be (1,0), got (%f, %f)", pt.x, pt.y)
	}
	c3 := NewCircleByPoint(NewPoint(8, 0), NewPoint(3, 0))
	inters2 := c1.IntersectCircle(c3)
	if inters2.SolutionNumber != 0 {
		t.Error("number of solution wrong")
	}
	c4 := NewCircleByPoint(NewPoint(1234.5678, 0), NewPoint(0, 1))
	inters3 := c1.IntersectCircle(c4)
	if inters3.SolutionNumber != 2 {
		t.Error("number of solution wrong")
	}
	pt1 := inters3.Solutions[0]
	pt2 := inters3.Solutions[1]
	if pt2.Serialize() != NewPoint(0, 1).Serialize() {
		t.Errorf("intersection point wrong, should be (0,1), got (%f, %f)", pt1.x, pt1.y)
	}
	if pt1.Serialize() != NewPoint(0, -1).Serialize() {
		t.Errorf("intersection point wrong, should be (0,-1), got (%f, %f)", pt2.x, pt2.y)
	}
}

func TestCircle_IntersectLine(t *testing.T) {
	c1 := NewCircleByRadius(NewPoint(0, 0), 2)
	l1 := NewLineFromTwoPoints(NewPoint(8, 2), NewPoint(10, 2))
	inters1 := c1.IntersectLine(l1)
	if inters1.SolutionNumber != 1 {
		t.Error("number of solution wrong")
	}
	pt := inters1.Solutions[0]
	if pt.Serialize() != NewPoint(0, 2).Serialize() {
		t.Errorf("intersection point wrong, should be (0,1), got (%f, %f)", pt.x, pt.y)
	}
	pt1 := NewPoint(math.Sqrt(3), 1)
	pt2 := NewPoint(2, 0)
	l2 := NewLineFromTwoPoints(pt1, pt2)
	inters2 := c1.IntersectLine(l2)
	if inters2.SolutionNumber != 2 {
		t.Error("number of solution wrong")
	}
	pti1 := inters2.Solutions[0]
	pti2 := inters2.Solutions[1]
	if pti1.Serialize() != pt1.Serialize() {
		t.Errorf("intersection point wrong, should be (Sqrt(3),1), got (%f, %f)", pti1.x, pti1.y)
	}
	if pti2.Serialize() != pt2.Serialize() {
		t.Errorf("intersection point wrong, should be (2,0), got (%f, %f)", pti2.x, pti2.y)
	}
}

func TestCircle_IntersectHalfLine(t *testing.T) {
	c1 := NewCircleByRadius(NewPoint(0, 0), 2)
	l1 := NewHalfLineFromTwoPoints(NewPoint(0, 2), NewPoint(10, 2))
	inters1 := c1.IntersectHalfLine(l1)
	if inters1.SolutionNumber != 1 {
		t.Error("number of solution wrong")
	}
	pt := inters1.Solutions[0]
	if pt.Serialize() != NewPoint(0, 2).Serialize() {
		t.Errorf("intersection point wrong, should be (0,1), got (%f, %f)", pt.x, pt.y)
	}
	pt1 := NewPoint(math.Sqrt(3), 1)
	pt2 := NewPoint(1, 0)
	l2 := NewHalfLineFromTwoPoints(pt2, pt1)
	inters2 := c1.IntersectHalfLine(l2)
	if inters2.SolutionNumber != 1 {
		t.Error("number of solution wrong")
	}
	pti1 := inters2.Solutions[0]
	if pti1.Serialize() != pt1.Serialize() {
		t.Errorf("intersection point wrong, should be (Sqrt(3),1), got (%f, %f)", pti1.x, pti1.y)
	}
}

func TestCircle_IntersectSegment(t *testing.T) {
	c1 := NewCircleByRadius(NewPoint(0, 0), 2)
	s1 := NewSegment(NewPoint(0, 1), NewPoint(0, 3))
	s2 := NewSegment(NewPoint(1, 0), NewPoint(0, 1))
	inters1 := c1.IntersectSegment(s1)
	inters2 := c1.IntersectSegment(s2)
	if inters1.SolutionNumber != 1 {
		t.Errorf("number of solution wrong, got %d", inters1.SolutionNumber)
	}
	if inters2.SolutionNumber != 0 {
		t.Errorf("number of solution wrong, got %d", inters2.SolutionNumber)
	}
}

func TestCircle_GetRandomPoint(t *testing.T) {
	c := NewCircleByRadius(NewPoint(12.34, 56.78), 90.12)
	pt := c.GetRandomPoint()
	if !pt.OnCircle(c) {
		t.Error("Point should on circle")
	}
}
