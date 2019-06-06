package hashset

import (
	"testing"
)

type point struct {
	Serializable
	x, y float64
}

func newTestPoint(x, y float64) *point {
	return &point{x: x, y: y}
}

func (p *point) Serialize() interface{} {
	return 12345*p.x + p.y
}

func TestHashSet_Add(t *testing.T) {
	a := NewHashSet()
	pt1 := newTestPoint(1, 2)
	pt2 := newTestPoint(3, 4)
	pt3 := newTestPoint(3, 4)
	a.Add(pt1)
	a.Add(pt2)
	a.Add(pt3)
	if len(a.m) != 2 {
		t.Errorf("Add: got %d instead of 2", len(a.m))
	}
}

func TestHashSet_Remove(t *testing.T) {
	a := NewHashSet()
	pt1 := newTestPoint(1, 2)
	pt2 := newTestPoint(3, 4)
	pt3 := newTestPoint(3, 4)
	a.Add(pt1)
	a.Add(pt2)
	a.Add(pt3)
	a.Remove(pt1)
	if len(a.m) != 1 {
		t.Errorf("Add: got %d instead of 1", len(a.m))
	}
	a.Remove(pt3)
	if len(a.m) != 0 {
		t.Errorf("Add: got %d instead of 0", len(a.m))
	}
}

func TestHashSet_Clone(t *testing.T) {
	a := NewHashSet()
	pt1 := newTestPoint(1, 2)
	pt2 := newTestPoint(3, 4)
	pt3 := newTestPoint(3, 4)
	a.Add(pt2)
	a.Add(pt3)
	b := a.Clone()
	a.Add(pt1)
	if len(a.m) != 2 && len(b.m) != 1 {
		t.Error("Clone error")
	}
}

func TestHashSet_Dict(t *testing.T) {
	a := NewHashSet()
	pt1 := newTestPoint(1, 2)
	pt2 := newTestPoint(3, 4)
	pt3 := newTestPoint(3, 4)
	a.Add(pt2)
	a.Add(pt3)
	a.Add(pt1)
	for _, elem := range a.Dict() {
		e := elem.(*point)
		if e.x != 1 && e.x != 3 {
			t.Error("iterating error")
		}
	}
}

func TestHashSet_Values(t *testing.T) {
	a := NewHashSet()
	pt1 := newTestPoint(1, 2)
	pt2 := newTestPoint(3, 4)
	a.Add(pt2)
	a.Add(pt1)
	if len(a.Values()) != 2 {
		t.Error("values() error")
	}
}
