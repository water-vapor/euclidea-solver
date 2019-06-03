package problem

import "github.com/water-vapor/euclidea-solver/pkg/geom"

//Problem 8: Line-Circle Intersection 2
func lineCircleIntersection2() *Statement {
	problem := geom.NewGeomBoard()
	pt1 := geom.NewPoint(0, 0)
	c := geom.NewCircleByRadius(pt1, 2)
	pt2 := geom.NewPoint(-4, 0)
	pt3 := geom.NewPoint(-2, 0)
	pt4 := geom.NewPoint(2, 0)

	problem.AddCircle(c)
	problem.AddPoint(pt1)
	problem.AddPoint(pt2)

	target := geom.NewGoal()
	target.Points.Add(pt3)
	target.Points.Add(pt4)
	return NewStatement(problem, target, "OOOOOOO", "Line-Circle Intersection 2")
}

//Problem 10: Angle of 3 Degree
func angelOf3Degree() *Statement {
	problem := geom.NewGeomBoard()
	pt1 := geom.NewPoint(0, 0)
	pt2 := geom.NewPoint(1, 0.0524077792830412040388058244740)
	hl := geom.NewHalfLineFromDirection(pt1, geom.NewVector2D(1, 0))

	problem.AddPoint(pt1)
	problem.AddHalfLine(hl)

	target := geom.NewGoal()
	target.Lines.Add(geom.NewLineFromTwoPoints(pt1, pt2))

	return NewStatement(problem, target, "OOOIIOI", "Angle of 3 Degree")
}
