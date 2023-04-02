package visitor

import "math"

type AreaCalc struct {
	area float64
}

func (a *AreaCalc) visitForSquare(s *Square) {
	a.area = s.side * s.side
}

func (a *AreaCalc) visitForCircle(c *Circle) {
	a.area = c.radius * c.radius * math.Pi
}

func (a *AreaCalc) visitForRectangle(r *Rectangle) {
	a.area = r.b * r.l
}
