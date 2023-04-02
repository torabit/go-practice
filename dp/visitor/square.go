package visitor

type Square struct {
	side float64
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() Shape {
	return SQUARE
}
