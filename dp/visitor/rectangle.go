package visitor

type Rectangle struct {
	l float64
	b float64
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) getType() Shape {
	return RECTANGLE
}
