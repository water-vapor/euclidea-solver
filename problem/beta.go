package problem

import (
	"github.com/water-vapor/euclidea-solver/pkg/geom"
	"math"
)

//Problem 1: Angle Bisector
func angleBisector() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(1, 0)
	pt3 := geom.NewPoint(1, math.Sqrt(3))

	problem.AddPoint(pt1)
	problem.HalfLines.Add(geom.NewHalfLineFromTwoPoints(pt1, pt2))
	problem.HalfLines.Add(geom.NewHalfLineFromTwoPoints(pt1, pt3))

	target := geom.NewTarget()
	target.Lines.Add(geom.NewLineAsAngleBisector(pt2, pt1, pt3))

	sequences := map[string]string{"E": "O+"}
	return NewStatement(problem, target, sequences, "2.1 Angle Bisector")
}

//Problem 2: Intersection of Angle Bisectors
func intersectionOfAngleBisectors() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(3, 0)
	pt3 := geom.NewPoint(1, 2)
	problem.AddPoint(pt1)
	problem.AddPoint(pt2)
	problem.AddPoint(pt3)
	problem.AddLine(geom.NewLineFromTwoPoints(pt1, pt2))
	problem.AddLine(geom.NewLineFromTwoPoints(pt2, pt3))
	problem.AddLine(geom.NewLineFromTwoPoints(pt1, pt3))

	l1 := geom.NewLineAsAngleBisector(pt1, pt2, pt3)
	l2 := geom.NewLineAsAngleBisector(pt2, pt1, pt3)
	pt4 := l1.IntersectLine(l2).Solutions[0]

	target := geom.NewTarget()
	target.Points.Add(pt4)

	sequences := map[string]string{"E": "OOOIOI", "L": "AA"}
	return NewStatement(problem, target, sequences, "2.2 Intersection of Angle Bisectors")
}

//Problem 3: Angle of 30 Degree
func angleOf30Degree() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(math.Sqrt(3), 1)

	problem.AddPoint(pt1)
	problem.HalfLines.Add(geom.NewHalfLineFromDirection(pt1, geom.NewVector2D(1, 0)))

	target := geom.NewTarget()
	target.Lines.Add(geom.NewLineFromTwoPoints(pt1, pt2))

	sequences := map[string]string{"E": "OOI"}
	return NewStatement(problem, target, sequences, "2.3 Angle of 30 Degree")
}

//Problem 4: Double Angle
func doubleAngle() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(math.Sqrt(3), 1)
	pt3 := geom.NewPoint(1, math.Sqrt(3))

	problem.AddPoint(pt1)
	problem.HalfLines.Add(geom.NewHalfLineFromTwoPoints(pt1, geom.NewPoint(1, 0)))
	problem.HalfLines.Add(geom.NewHalfLineFromTwoPoints(pt1, pt2))

	target := geom.NewTarget()
	target.Lines.Add(geom.NewLineFromTwoPoints(pt1, pt3))
	sequences := map[string]string{"E": "OOI"}
	return NewStatement(problem, target, sequences, "2.4 Double Angle")
}

//Problem 5: Cut Rectangle
func cutRectangle() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(3, 0)
	pt3 := geom.NewPoint(3, 2)
	pt4 := geom.NewPoint(0, 2)
	ptx := geom.NewPoint(2, 4)
	l1 := geom.NewLineFromTwoPoints(pt1, pt2)
	l2 := geom.NewLineFromTwoPoints(pt2, pt3)
	l3 := geom.NewLineFromTwoPoints(pt3, pt4)
	l4 := geom.NewLineFromTwoPoints(pt4, pt1)
	l5 := geom.NewLineFromTwoPoints(pt1, pt3)
	l6 := geom.NewLineFromTwoPoints(pt2, pt4)
	pt5 := l5.IntersectLine(l6).Solutions[0]
	problem.AddPoint(pt1)
	problem.AddPoint(pt2)
	problem.AddPoint(pt3)
	problem.AddPoint(pt4)
	problem.AddLine(l1)
	problem.AddLine(l2)
	problem.AddLine(l3)
	problem.AddLine(l4)
	problem.AddPoint(ptx)

	target := geom.NewTarget()
	lx := geom.NewLineFromTwoPoints(ptx, pt5)
	target.Lines.Add(lx)
	sequences := map[string]string{"E": "III"}
	return NewStatement(problem, target, sequences, "2.5 Cut Rectangle")
}

//Problem 6: Drop a Perpendicular
func dropAPerpendicular() *Statement {
	problem := geom.NewBoard()
	l := geom.NewLineFromTwoPoints(geom.NewPoint(0, 0), geom.NewPoint(1, 0))
	pt := geom.NewPoint(0, 1)
	problem.AddLine(l)
	problem.AddPoint(pt)

	target := geom.NewTarget()
	target.Lines.Add(geom.NewLineFromTwoPoints(geom.NewPoint(0, 0), pt))
	sequences := map[string]string{"E": "OOI", "L": "O+"}
	return NewStatement(problem, target, sequences, "2.6 Drop a Perpendicular")
}

//Problem 7: Erect a Perpendicular
func erectAPerpendicular() *Statement {
	problem := geom.NewBoard()
	l := geom.NewLineFromTwoPoints(geom.NewPoint(0, 0), geom.NewPoint(1, 0))
	pt := geom.NewPoint(0, 0)
	problem.AddLine(l)
	problem.AddPoint(pt)
	// hint E: use an arbitrary point
	problem.AddPoint(geom.NewPoint(-1, -1))
	// hint L: use two arbitrary points on two sides
	// problem.AddPoint(geom.NewPoint(-1, 0))
	// problem.AddPoint(geom.NewPoint(1, 0))

	target := geom.NewTarget()
	target.Lines.Add(geom.NewLineFromTwoPoints(pt, geom.NewPoint(0, 1)))
	sequences := map[string]string{"E": "OII", "L": "A"}
	return NewStatement(problem, target, sequences, "2.7 Erect a Perpendicular")
}

//Problem 8: Tangent to Circle at Point
func tangentToCircleAtPoint() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(1, 0)
	c := geom.NewCircleByRadius(pt1, 1)
	problem.AddPoint(pt1)
	problem.AddPoint(pt2)
	problem.AddCircle(c)
	// hint E: use an arbitrary point on the circle
	problem.AddPoint(geom.NewPoint(0, -1))

	target := geom.NewTarget()
	target.Lines.Add(geom.NewLineFromTwoPoints(pt2, geom.NewPoint(1, 1)))
	sequences := map[string]string{"E": "OOI", "L": "IL"}
	return NewStatement(problem, target, sequences, "2.8 Tangent to Circle at Point")
}

//Problem 9: Circle Tangent to Line
func circleTangentToLine() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(1, 0)
	l := geom.NewLineFromTwoPoints(pt1, pt2)
	pt3 := geom.NewPoint(0, 1)
	problem.AddLine(l)
	problem.AddPoint(pt3)

	target := geom.NewTarget()
	target.Circles.Add(geom.NewCircleByRadius(pt3, 1))
	sequences := map[string]string{"E": "LO"}
	return NewStatement(problem, target, sequences, "2.9 Circle Tangent to Line")
}

//Problem 10: Circle in Rhombus
func circleInRhombus() *Statement {
	problem := geom.NewBoard()
	pt1 := geom.NewPoint(0, -1)
	pt2 := geom.NewPoint(3, 0)
	pt3 := geom.NewPoint(0, 1)
	pt4 := geom.NewPoint(-3, 0)
	l1 := geom.NewLineFromTwoPoints(pt1, pt2)
	l2 := geom.NewLineFromTwoPoints(pt2, pt3)
	l3 := geom.NewLineFromTwoPoints(pt3, pt4)
	l4 := geom.NewLineFromTwoPoints(pt4, pt1)
	pt5 := geom.NewPoint(0, 0)
	pt6 := l1.GetTangentLineWithPoint(pt5).IntersectLine(l1).Solutions[0]
	c := geom.NewCircleByPoint(pt5, pt6)
	problem.AddPoint(pt1)
	problem.AddPoint(pt2)
	problem.AddPoint(pt3)
	problem.AddPoint(pt4)
	problem.AddLine(l1)
	problem.AddLine(l2)
	problem.AddLine(l3)
	problem.AddLine(l4)

	target := geom.NewTarget()
	target.Circles.Add(c)
	sequences := map[string]string{"E": "IILO"}
	return NewStatement(problem, target, sequences, "2.10 Circle in Rhombus")
}
