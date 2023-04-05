package flyweight

type CTDress struct {
	color string
}

func (ct *CTDress) getColor() string {
	return ct.color
}

func newCTDress() *CTDress {
	return &CTDress{color: "blue"}
}
