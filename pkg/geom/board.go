package geom

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/water-vapor/euclidea-solver/configs"
	"github.com/water-vapor/euclidea-solver/pkg/hashset"
	"math"
)

// A geometry board containing all geometry objects
type Board struct {
	Points    *hashset.HashSet
	Lines     *hashset.HashSet
	Circles   *hashset.HashSet
	HalfLines *hashset.HashSet
	Segments  *hashset.HashSet
}

func NewGeomBoard() *Board {
	return &Board{hashset.NewHashSet(), hashset.NewHashSet(), hashset.NewHashSet(), hashset.NewHashSet(), hashset.NewHashSet()}
}

func (gb *Board) Clone() *Board {
	ret := NewGeomBoard()
	ret.Points = gb.Points.Clone()
	ret.Lines = gb.Lines.Clone()
	ret.Circles = gb.Circles.Clone()
	ret.HalfLines = gb.HalfLines.Clone()
	ret.Segments = gb.Segments.Clone()
	return ret
}

// Add a point and do nothing
func (gb *Board) AddPoint(pt *Point) {
	if math.Abs(pt.x) > configs.MaxPointCoord || math.Abs(pt.y) > configs.MaxPointCoord {
		return
	}
	gb.Points.Add(pt)
}

// Add a half line and its end point, do nothing
func (gb *Board) AddHalfLine(h *HalfLine) {
	gb.AddPoint(h.point)
	gb.HalfLines.Add(h)
}

// Add a segment and its end points, do nothing
func (gb *Board) AddSegment(s *Segment) {
	gb.AddPoint(s.point1)
	gb.AddPoint(s.point2)
	gb.Segments.Add(s)
}

func (gb *Board) AddCircle(c *Circle) {
	// calculate new intersection points
	for _, elem := range gb.Circles.Dict() {
		circle := elem.(*Circle)
		inters := c.IntersectCircle(circle)
		for i := 0; i < inters.SolutionNumber; i++ {
			gb.AddPoint(inters.Solutions[i])
		}
	}
	for _, elem := range gb.Lines.Dict() {
		line := elem.(*Line)
		inters := c.IntersectLine(line)
		for i := 0; i < inters.SolutionNumber; i++ {
			gb.AddPoint(inters.Solutions[i])
		}
	}
	for _, elem := range gb.HalfLines.Dict() {
		hl := elem.(*HalfLine)
		inters := c.IntersectHalfLine(hl)
		for i := 0; i < inters.SolutionNumber; i++ {
			gb.AddPoint(inters.Solutions[i])
		}
	}
	for _, elem := range gb.Segments.Dict() {
		s := elem.(*Segment)
		inters := c.IntersectSegment(s)
		for i := 0; i < inters.SolutionNumber; i++ {
			gb.AddPoint(inters.Solutions[i])
		}
	}
	// Add circle in the end
	gb.Circles.Add(c)
}

func (gb *Board) AddLine(l *Line) {
	// calculate new intersection points
	for _, elem := range gb.Circles.Dict() {
		circle := elem.(*Circle)
		inters := l.IntersectCircle(circle)
		for i := 0; i < inters.SolutionNumber; i++ {
			gb.AddPoint(inters.Solutions[i])
		}
	}
	for _, elem := range gb.Lines.Dict() {
		line := elem.(*Line)
		inters := l.IntersectLine(line)
		for i := 0; i < inters.SolutionNumber; i++ {
			gb.AddPoint(inters.Solutions[i])
		}
	}
	for _, elem := range gb.HalfLines.Dict() {
		hl := elem.(*HalfLine)
		inters := l.IntersectHalfLine(hl)
		for i := 0; i < inters.SolutionNumber; i++ {
			gb.AddPoint(inters.Solutions[i])
		}
	}
	for _, elem := range gb.Segments.Dict() {
		s := elem.(*Segment)
		inters := l.IntersectSegment(s)
		for i := 0; i < inters.SolutionNumber; i++ {
			gb.AddPoint(inters.Solutions[i])
		}
	}
	gb.Lines.Add(l)
}

// get a random point on each geometry object
func (gb *Board) GenerateRandomPoints() []*Point {
	pts := make([]*Point, 0)
	for _, elem := range gb.Circles.Dict() {
		circle := elem.(*Circle)
		pts = append(pts, circle.GetRandomPoint())
	}
	for _, elem := range gb.Lines.Dict() {
		line := elem.(*Line)
		pts = append(pts, line.GetRandomPoint())
	}
	for _, elem := range gb.HalfLines.Dict() {
		hl := elem.(*HalfLine)
		pts = append(pts, hl.GetRandomPoint())
	}
	for _, elem := range gb.Segments.Dict() {
		s := elem.(*Segment)
		pts = append(pts, s.GetRandomPoint())
	}
	return pts
}

func (gb *Board) GeneratePlot(fileName string) error {
	dc := gg.NewContext(configs.ImageSize, configs.ImageSize)
	// p.Title.Text = "Graphics"

	// Generate Plot Range
	xmin, xmax, ymin, ymax := float64(0), float64(1), float64(0), float64(1)
	// Visible points includes points in point set, endpoints of segments and halflines
	visiblePoints := hashset.NewHashSet()
	for _, elem := range gb.Points.Dict() {
		pt := elem.(*Point)
		visiblePoints.Add(pt)
	}
	for _, elem := range gb.Segments.Dict() {
		s := elem.(*Segment)
		visiblePoints.Add(s.point1)
		visiblePoints.Add(s.point2)
	}
	for _, elem := range gb.HalfLines.Dict() {
		h := elem.(*HalfLine)
		visiblePoints.Add(h.point)
	}
	for _, elem := range visiblePoints.Dict() {
		pt := elem.(*Point)
		if pt.x > xmax {
			xmax = pt.x
		}
		if pt.x < xmin {
			xmin = pt.x
		}
		if pt.y > ymax {
			ymax = pt.y
		}
		if pt.y < ymin {
			ymin = pt.y
		}
	}
	xrange := xmax - xmin
	yrange := ymax - ymin
	xyrange := math.Max(xrange, yrange) * 1.1
	xmid := (xmin + xmax) / 2
	ymid := (ymin + ymax) / 2
	xmin = xmid - xyrange*0.5
	xmax = xmid + xyrange*0.5
	ymin = ymid - xyrange*0.5
	ymax = ymid + xyrange*0.5
	boundpt1 := NewPoint(xmin, ymin)
	boundpt2 := NewPoint(xmin, ymax)
	boundpt3 := NewPoint(xmax, ymin)
	boundpt4 := NewPoint(xmax, ymax)
	boundline1 := NewSegment(boundpt1, boundpt2)
	boundline2 := NewSegment(boundpt3, boundpt4)
	boundline3 := NewSegment(boundpt1, boundpt3)
	boundline4 := NewSegment(boundpt2, boundpt4)

	xCoordToImg := func(x float64) float64 {
		return (x - xmin) * configs.ImageSize / xyrange
	}
	yCoordToImg := func(y float64) float64 {
		return configs.ImageSize - (y-ymin)*configs.ImageSize/xyrange
	}
	rCoordToImg := func(r float64) float64 {
		return r * configs.ImageSize / xyrange
	}

	for _, elem := range gb.Lines.Dict() {
		l := elem.(*Line)
		intersectionPoints := hashset.NewHashSet()
		i1 := l.IntersectSegment(boundline1)
		if i1.SolutionNumber != 0 {
			intersectionPoints.Add(i1.Solutions[0])
		}
		i2 := l.IntersectSegment(boundline2)
		if i2.SolutionNumber != 0 {
			intersectionPoints.Add(i2.Solutions[0])
		}
		i3 := l.IntersectSegment(boundline3)
		if i3.SolutionNumber != 0 {
			intersectionPoints.Add(i3.Solutions[0])
		}
		i4 := l.IntersectSegment(boundline4)
		if i4.SolutionNumber != 0 {
			intersectionPoints.Add(i4.Solutions[0])
		}
		if len(intersectionPoints.Dict()) != 2 {
			fmt.Println("Warning: Error plotting line")
			for _, elem := range intersectionPoints.Dict() {
				pt := elem.(*Point)
				fmt.Println(l)
				fmt.Println(pt.x, pt.y)
			}
		}
		pts := make([]*Point, 0)
		for _, elem := range intersectionPoints.Dict() {
			pt := elem.(*Point)
			pts = append(pts, pt)
		}
		dc.DrawLine(xCoordToImg(pts[0].x), yCoordToImg(pts[0].y), xCoordToImg(pts[1].x), yCoordToImg(pts[1].y))

	}
	for _, elem := range gb.HalfLines.Dict() {
		h := elem.(*HalfLine)
		var pt *Point
		i1 := h.IntersectSegment(boundline1)
		if i1.SolutionNumber != 0 {
			pt = i1.Solutions[0]
		}
		i2 := h.IntersectSegment(boundline2)
		if i2.SolutionNumber != 0 {
			pt = i2.Solutions[0]
		}
		i3 := h.IntersectSegment(boundline3)
		if i3.SolutionNumber != 0 {
			pt = i3.Solutions[0]
		}
		i4 := h.IntersectSegment(boundline4)
		if i4.SolutionNumber != 0 {
			pt = i4.Solutions[0]
		}
		dc.DrawLine(xCoordToImg(h.point.x), yCoordToImg(h.point.y), xCoordToImg(pt.x), yCoordToImg(pt.y))

	}
	for _, elem := range gb.Segments.Dict() {
		s := elem.(*Segment)
		dc.DrawLine(xCoordToImg(s.point1.x), yCoordToImg(s.point1.y), xCoordToImg(s.point2.x), yCoordToImg(s.point2.y))
	}
	for _, elem := range gb.Circles.Dict() {
		c := elem.(*Circle)
		dc.DrawCircle(xCoordToImg(c.center.x), yCoordToImg(c.center.y), rCoordToImg(c.r))
	}

	dc.Stroke()
	err := dc.SavePNG(fileName)
	return err
}