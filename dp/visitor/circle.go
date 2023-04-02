package visitor

type Circle struct {
	radius float64
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() Shape {
	return CIRCLE
}
